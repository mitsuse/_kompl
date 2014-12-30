package binary

import (
	"encoding/binary"
	"io"
)

type Reader struct {
	reader io.Reader
	err    error
}

func NewReader(reader io.Reader) *Reader {
	r := &Reader{
		reader: reader,
	}

	return r
}

func (r *Reader) Read(value interface{}) {
	if r.err != nil {
		return
	}

	r.err = binary.Read(r.reader, binary.LittleEndian, value)
}

func (r *Reader) Error() error {
	return r.err
}
