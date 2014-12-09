package compl

import (
	"bufio"
	"io"
)

type NgramIterator struct {
	scanner *bufio.Scanner
	wordSeq []string
	order   int
}

func NewNgramIterator(order int, reader io.Reader) *NgramIterator {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)

	wordSeq := make([]string, order)
	for i := 0; i < len(wordSeq); i++ {
		wordSeq[i] = ""
	}

	iter := &NgramIterator{
		scanner: scanner,
		wordSeq: wordSeq,
		order:   order,
	}

	return iter
}

func (iter *NgramIterator) Order() int {
	return iter.order
}

func (iter *NgramIterator) Iterate() bool {
	if iter.scanner.Scan() {
		for i := 1; i < len(iter.wordSeq); i++ {
			iter.wordSeq[i-1] = iter.wordSeq[i]
		}
		iter.wordSeq[iter.order-1] = iter.scanner.Text()

		return true
	}

	return false
}

func (iter *NgramIterator) Get() []string {
	wordSeq := make([]string, iter.order)
	copy(wordSeq, iter.wordSeq)

	return wordSeq
}

func (iter *NgramIterator) Error() error {
	return iter.scanner.Err()
}
