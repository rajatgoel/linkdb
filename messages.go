package linkdb

import (
	"encoding/binary"
	"errors"
	"hash/crc64"

	"google.golang.org/protobuf/proto"
)

var (
	tbl = crc64.MakeTable(crc64.ISO)
)

func marshal[T proto.Message](t T) ([]byte, error) {
	s := proto.Size(t)
	ts := s + 8 /* length */ + 8 /* checksum */
	if ts > MinBlockSize {
		return nil, errors.New("proto message too large")
	}

	buf := make([]byte, MinBlockSize)
	_, err := proto.MarshalOptions{}.MarshalAppend(buf[16:], t)
	if err != nil {
		return nil, err
	}

	checksum := crc64.Checksum(buf[16:16+s], tbl)
	binary.LittleEndian.PutUint64(buf[:8], uint64(s))
	binary.LittleEndian.PutUint64(buf[8:16], checksum)
	return buf, nil
}

func unmarshal[T proto.Message](data []byte) (T, error) {
	var zero T
	if len(data) < 16 {
		return zero, errors.New("too small to unmarshal")
	}
	l := int(binary.LittleEndian.Uint64(data[:8]))
	if l == 0 || l > len(data)-16 {
		return zero, errors.New("invalid length")
	}
	checksum := binary.LittleEndian.Uint64(data[8:16])
	if crc64.Checksum(data[16:16+l], tbl) != checksum {
		return zero, errors.New("checksum mismatch")
	}

	var t T
	if err := proto.Unmarshal(data[16:16+l], t); err != nil {
		return zero, err
	}
	return t, nil
}
