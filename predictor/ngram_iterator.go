package predictor

import (
	"bufio"
	"io"
	"strings"
)

type NgramIterator struct {
	wordScanner *bufio.Scanner
	lineScanner *bufio.Scanner
	wordSeq     []string
	order       int
	err         error
}

func NewNgramIterator(order int, reader io.Reader) *NgramIterator {
	lineScanner := bufio.NewScanner(reader)

	wordSeq := make([]string, order)
	for i := 0; i < len(wordSeq); i++ {
		wordSeq[i] = ""
	}

	iter := &NgramIterator{
		lineScanner: lineScanner,
		wordSeq:     wordSeq,
		order:       order,
	}

	return iter
}

func (iter *NgramIterator) Order() int {
	return iter.order
}

func (iter *NgramIterator) Iterate() bool {
	var hasNext bool

	for {
		if iter.wordScanner != nil && iter.wordScanner.Scan() {
			for i := 1; i < len(iter.wordSeq); i++ {
				iter.wordSeq[i-1] = iter.wordSeq[i]
			}
			iter.wordSeq[iter.order-1] = iter.wordScanner.Text()

			hasNext = true
			iter.err = iter.wordScanner.Err()

			break
		}

		if !iter.lineScanner.Scan() {
			hasNext = false
			iter.err = iter.lineScanner.Err()

			break
		}

		reader := strings.NewReader(iter.lineScanner.Text())
		iter.wordScanner = bufio.NewScanner(reader)
		iter.wordScanner.Split(bufio.ScanWords)
	}

	return hasNext
}

func (iter *NgramIterator) Get() []string {
	wordSeq := make([]string, iter.order)
	copy(wordSeq, iter.wordSeq)

	return wordSeq
}

func (iter *NgramIterator) Error() error {
	return iter.err
}
