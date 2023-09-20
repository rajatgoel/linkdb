package linkdb

import (
	"io"
)

const (
	MinBlockSize = 4 * 1024
)

type Disk interface {
	io.WriterAt
	io.ReaderAt

	BlockSize() int64
	Capacity() int64
}
