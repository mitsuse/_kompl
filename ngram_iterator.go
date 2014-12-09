package compl

import (
	"io"
)

type NgramIterator struct {
	reader io.Reader
}

func NewNgramIterator(reader io.Reader) *NgramIterator {
	iter := &NgramIterator{
		reader: reader,
	}

	return iter
}

func (iter *NgramIterator) Iterate() bool {
	// TODO: Implement this.
	return false
}

func (iter *NgramIterator) Get() []string {
	// TODO: Implement this.
	return nil
}

func (iter *NgramIterator) Error() error {
	// TODO: Implement this.
	return nil
}
