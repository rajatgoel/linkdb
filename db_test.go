package linkdb

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/exp/maps"

	linkdbtestv1 "github.com/rajatgoel/linkdb/gen/linkdbtest/v1"
)

type memDisk struct {
	buf       []byte
	blockSize int64
}

func (m memDisk) BlockSize() int64 {
	return m.blockSize
}

func (m memDisk) Capacity() int64 {
	return int64(cap(m.buf))
}

func (m memDisk) ReadAt(buf []byte, offset int64) (int, error) {
	n := copy(buf, m.buf[offset:])
	return n, nil
}

func (m memDisk) WriteAt(p []byte, off int64) (int, error) {
	n := copy(m.buf[off:], p)
	if n != len(p) {
		panic("out of bound write")
	}

	return n, nil
}

func TestFormat(t *testing.T) {
	disk := memDisk{
		buf:       make([]byte, 2*4096),
		blockSize: 4096,
	}

	_, err := Open[*linkdbtestv1.LinkDBMeta](disk)
	require.Error(t, err)

	db, err := Format(disk, &linkdbtestv1.LinkDBMeta{
		Meta: &linkdbtestv1.LinkDBMeta_DiskMeta{
			DiskMeta: &linkdbtestv1.DiskMeta{DiskId: "test-disk"},
		},
	})
	require.NoError(t, err)
	require.Empty(t, db.List())

	got := db.Metadata()
	require.Equal(t, "test-disk", got.GetDiskMeta().DiskId)

	err = db.Put("key1", &linkdbtestv1.LinkDBMeta{
		Meta: &linkdbtestv1.LinkDBMeta_KeyMeta{
			KeyMeta: &linkdbtestv1.KeyMeta{Value: "key1-value1"},
		},
	}, nil)
	require.NoError(t, err)

	err = db.Put("key2", &linkdbtestv1.LinkDBMeta{
		Meta: &linkdbtestv1.LinkDBMeta_KeyMeta{
			KeyMeta: &linkdbtestv1.KeyMeta{Value: "key2-value2"},
		},
	}, nil)
	require.Error(t, err)

	require.Contains(t, db.List(), "key1")
}

func TestGetPut(t *testing.T) {
	disk := memDisk{
		buf:       make([]byte, 20000*4096),
		blockSize: 4096,
	}

	db, err := Format(disk, &linkdbtestv1.LinkDBMeta{
		Meta: &linkdbtestv1.LinkDBMeta_DiskMeta{
			DiskMeta: &linkdbtestv1.DiskMeta{DiskId: "test-disk"},
		},
	})
	require.NoError(t, err)
	require.Empty(t, db.List())

	keys := make(map[string]int)
	for i := 0; i < 5000; i++ {
		r := rand.Intn(100)
		err = db.Put(strconv.Itoa(r), &linkdbtestv1.LinkDBMeta{
			Meta: &linkdbtestv1.LinkDBMeta_KeyMeta{
				KeyMeta: &linkdbtestv1.KeyMeta{Value: strconv.Itoa(i)},
			},
		}, make([]byte, i))
		require.NoError(t, err)
		keys[strconv.Itoa(r)] = i
	}

	for i := 0; i < 5000; i++ {
		r := maps.Keys(keys)[rand.Intn(len(keys))]

		v, b, err := db.Get(r)
		require.NoError(t, err)
		require.Equal(t, strconv.Itoa(keys[r]), v.GetKeyMeta().Value)
		require.Len(t, b, keys[r])
	}
}

func TestAlignLeft(t *testing.T) {
	require.Equal(t, int64(0), alignLeft(0, 8))
	require.Equal(t, int64(0), alignLeft(1, 8))
	require.Equal(t, int64(8), alignLeft(8, 8))
	require.Equal(t, int64(8), alignLeft(9, 8))
}

func TestAlignRight(t *testing.T) {
	require.Equal(t, int64(0), alignRight(0, 8))
	require.Equal(t, int64(8), alignRight(1, 8))
	require.Equal(t, int64(8), alignRight(8, 8))
	require.Equal(t, int64(16), alignRight(9, 8))
}
