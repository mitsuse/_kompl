package binary

import (
	"encoding/binary"
	"io"
)

type Writer struct {
	writer io.Writer
	err    error
}

func NewWriter(writer io.Writer) *Writer {
	w := &Writer{
		writer: writer,
	}

	return w
}

func (w *Writer) Write(value interface{}) {
	if w.err != nil {
		return
	}

	w.err = binary.Write(w.writer, binary.LittleEndian, value)
}

func (w *Writer) Error() error {
	return w.err
}
