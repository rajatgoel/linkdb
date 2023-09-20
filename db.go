// linkdb is a key-value store for persistent storage. It aims to store both small and
// large values efficiently with minimal disk IOPS per operation. Since it keeps all the
// keys in memory, it is not suitable for storing large number of keys.
package linkdb

import (
	"errors"
	"fmt"
	"hash/crc64"
	"os"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"

	linkdbv1 "github.com/rajatgoel/linkdb/gen/linkdb/v1"
)

type Database[T proto.Message] interface {
	Metadata() T
	Get(key string) (T, []byte, error)
	GetRange(key string, off int64, len int64) (T, []byte, error)
	Put(key string, t T, value []byte) error
	List() map[string]T
}

type db[T proto.Message] struct {
	disk Disk

	nextVersion int
	header      *linkdbv1.DiskHeader
	keys        map[string]*linkdbv1.KeyValue
	gaps        []*linkdbv1.Gap
}

func Format[T proto.Message](disk Disk, metadata T, initValues ...T) (Database[T], error) {
	bs, c := disk.BlockSize(), disk.Capacity()
	if bs < MinBlockSize {
		return nil, errors.New("disk block size too small")
	}
	if c < 2*bs {
		return nil, errors.New("disk too small")
	}
	if c%bs != 0 {
		return nil, errors.New("disk capacity not a multiple of block size")
	}

	metadataBytes, err := anypb.New(metadata)
	if err != nil {
		return nil, errors.New("invalid disk header metadata")
	}

	buf := make([]byte, 0, (len(initValues)+1)*int(bs))
	for _, val := range initValues {
		m, err := marshal(val)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal init value: %w", err)
		}

		buf = append(buf, m...)
	}

	gapBytes, err := marshal(&linkdbv1.Gap{
		Offset:    bs,
		GapLength: c - bs,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to marshal gapBytes header: %w", err)
	}
	buf = append(buf, gapBytes...)

	if n, err := disk.WriteAt(buf, bs); err != nil || n != len(buf) {
		return nil, fmt.Errorf("failed to write gapBytes header: %w", err)
	}

	header := &linkdbv1.DiskHeader{
		BlockSize: bs,
		Capacity:  c,
		Metadata:  metadataBytes,
	}
	headerBytes, err := marshal(header)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal disk header: %w", err)
	}
	if n, err := disk.WriteAt(headerBytes, 0); err != nil || n != int(bs) {
		return nil, fmt.Errorf("failed to write disk header: %w", err)
	}

	return &db[T]{
		disk: disk,

		nextVersion: 1,
		header:      header,
		keys:        make(map[string]*linkdbv1.KeyValue),
		gaps:        []*linkdbv1.Gap{{Offset: bs, GapLength: c - bs}},
	}, nil
}

func Open[T proto.Message](disk Disk) (Database[T], error) {
	bs := disk.BlockSize()
	if bs < MinBlockSize {
		return nil, errors.New("disk block size too small")
	}

	headerBytes := make([]byte, bs)
	n, err := disk.ReadAt(headerBytes[:], 0)
	if n != int(bs) || err != nil {
		return nil, fmt.Errorf("failed to read header: %w", err)
	}

	header, err := unmarshal[*linkdbv1.DiskHeader](headerBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal disk header: %w", err)
	}

	if header.Capacity != disk.Capacity() || header.BlockSize != disk.BlockSize() {
		return nil, errors.New("disk parameters do not match original parameters from format")
	}

	ty, err := header.Metadata.UnmarshalNew()
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal metadata from disk header: %w", err)
	}
	if _, ok := ty.(T); !ok {
		return nil, errors.New("disk header metadata type does not match")
	}

	nextVersion := int64(1)
	offset := header.BlockSize
	keys, gaps := make(map[string]*linkdbv1.KeyValue), make([]*linkdbv1.Gap, 10)
	for offset < header.Capacity {
		nextHeaderBytes := make([]byte, header.BlockSize)
		n, err := disk.ReadAt(nextHeaderBytes[:], offset)
		if n != int(header.BlockSize) || err != nil {
			return nil, fmt.Errorf("failed to read header at offset %d: %w", offset, err)
		}

		nextHeader, err := unmarshal[*linkdbv1.Wrapper](nextHeaderBytes)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal header at offset %d: %w", offset, err)
		}

		switch m := nextHeader.Wrapped.(type) {
		case *linkdbv1.Wrapper_Gap:
			gaps = append(gaps, m.Gap)
			offset += m.Gap.GapLength
		case *linkdbv1.Wrapper_KeyValue:
			ty, err := m.KeyValue.Metadata.UnmarshalNew()
			if err != nil {
				return nil, fmt.Errorf("failed to unmarshal metadata at offset %d: %w", offset, err)
			}
			if _, ok := ty.(T); !ok {
				return nil, fmt.Errorf("metadata type at offset %d does not match", offset)
			}

			if m.KeyValue.Version > nextVersion {
				nextVersion = m.KeyValue.Version + 1
			}

			if _, ok := keys[m.KeyValue.Key]; !ok || m.KeyValue.Version > keys[m.KeyValue.Key].Version {
				keys[m.KeyValue.Key] = m.KeyValue
			} else {
				gaps = append(gaps, &linkdbv1.Gap{
					Offset:    offset,
					GapLength: bs + alignRight(m.KeyValue.ValueLength, bs),
				})
			}

			offset += bs + alignRight(m.KeyValue.ValueLength, bs)
		}
	}

	return &db[T]{
		disk: disk,

		nextVersion: int(nextVersion),
		header:      header,
		keys:        keys,
		gaps:        gaps,
	}, nil
}

func (d *db[T]) Metadata() T {
	ty, _ := d.header.Metadata.UnmarshalNew()
	return ty.(T)
}

func (d *db[T]) Get(key string) (T, []byte, error) {
	var zero T

	kv, ok := d.keys[key]
	if !ok {
		return zero, nil, os.ErrNotExist
	}

	l := alignRight(kv.ValueLength, d.header.BlockSize)

	buf := make([]byte, l)
	n, err := d.disk.ReadAt(buf[:], kv.Offset+d.header.BlockSize)
	if n != int(l) || err != nil {
		return zero, nil, fmt.Errorf("failed to read value at offset %d: %w", kv.Offset, err)
	}

	ty, _ := kv.Metadata.UnmarshalNew()
	return ty.(T), buf[:kv.ValueLength], nil
}

func (d *db[T]) GetRange(key string, off int64, len int64) (T, []byte, error) {
	var zero T

	kv, ok := d.keys[key]
	if !ok {
		return zero, nil, os.ErrNotExist
	}

	if off < 0 || off >= kv.ValueLength {
		return zero, nil, errors.New("invalid offset")
	}
	if len < 0 || off+len > kv.ValueLength {
		return zero, nil, errors.New("invalid length")
	}

	alignedStart := alignLeft(off, d.header.BlockSize)
	alignedEnd := alignRight(off+len, d.header.BlockSize)
	buf := make([]byte, alignedEnd-alignedStart)
	n, err := d.disk.ReadAt(buf[:], kv.Offset+d.header.BlockSize+alignedStart)
	if n != int(alignedEnd-alignedStart) || err != nil {
		return zero, nil, fmt.Errorf("failed to read value at offset %d: %w", kv.Offset, err)
	}

	ty, _ := kv.Metadata.UnmarshalNew()
	return ty.(T), buf[off-alignedStart : off-alignedStart+len], nil
}

func (d *db[T]) Put(key string, t T, value []byte) error {
	l := alignRight(int64(len(value)), d.header.BlockSize)

	tpb, err := anypb.New(t)
	if err != nil {
		return errors.New("failed to marshal metadata")
	}

	gapIdx := -1
	minGapLength := d.header.BlockSize + l
	for idx, gap := range d.gaps {
		if gap.GapLength >= minGapLength || (gapIdx != -1 && gap.GapLength < d.gaps[gapIdx].GapLength) {
			gapIdx = idx
		}
	}
	if gapIdx == -1 {
		return errors.New("no gaps large enough to fit value")
	}

	kv := &linkdbv1.KeyValue{
		Key:           key,
		Version:       int64(d.nextVersion),
		Offset:        d.gaps[gapIdx].Offset,
		Metadata:      tpb,
		ValueLength:   int64(len(value)),
		ValueChecksum: crc64.Checksum(value, tbl),
	}

	kvBytes, err := marshal(kv)
	if err != nil {
		return fmt.Errorf("failed to marshal key-value header: %w", err)
	}

	var gapsToAdd []*linkdbv1.Gap
	if oldKV, ok := d.keys[key]; ok {
		gapsToAdd = append(gapsToAdd, &linkdbv1.Gap{
			Offset:    oldKV.Offset,
			GapLength: d.header.BlockSize + alignRight(oldKV.ValueLength, d.header.BlockSize),
		})
	}

	var bufToWrite []byte
	if d.gaps[gapIdx].GapLength == d.header.BlockSize+l {
		if int(l) == len(value) {
			bufToWrite = value
		} else {
			bufToWrite = make([]byte, l)
			copy(bufToWrite[:len(value)], value)
		}
	} else {
		bufToWrite = make([]byte, d.header.BlockSize+l)
		copy(bufToWrite[:len(value)], value)

		newGap := &linkdbv1.Gap{
			Offset:    kv.Offset + d.header.BlockSize + l,
			GapLength: d.gaps[gapIdx].GapLength - d.header.BlockSize - l,
		}

		gapBytes, err := marshal(newGap)
		if err != nil {
			return fmt.Errorf("failed to marshal new gap header: %w", err)
		}

		copy(bufToWrite[d.header.BlockSize+l:], gapBytes)
		gapsToAdd = append(gapsToAdd, newGap)
	}

	if n, err := d.disk.WriteAt(bufToWrite, kv.Offset+d.header.BlockSize); err != nil || n != len(bufToWrite) {
		return fmt.Errorf("failed to write value: %w", err)
	}

	if n, err := d.disk.WriteAt(kvBytes, kv.Offset); err != nil || n != int(d.header.BlockSize) {
		// At this point we can't be sure about the state of the disk.
		// We need to start over and load the full state from disk again.
		panic("failed to write key-value header")
	}

	d.gaps[gapIdx] = d.gaps[len(d.gaps)-1]
	d.gaps[len(d.gaps)-1] = nil
	d.gaps = d.gaps[:len(d.gaps)-1]
	d.gaps = append(d.gaps, gapsToAdd...)

	d.keys[key] = kv
	d.nextVersion++
	return nil
}

func (d *db[T]) List() map[string]T {
	r := make(map[string]T, len(d.keys))
	for k, v := range d.keys {
		ty, _ := v.Metadata.UnmarshalNew()
		r[k] = ty.(T)
	}
	return r
}

func alignLeft(n, a int64) int64 {
	return n / a * a
}

func alignRight(n, a int64) int64 {
	return (n + a - 1) / a * a
}
