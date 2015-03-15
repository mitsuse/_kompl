package ngram

import (
	"bufio"
	"io"
	"strings"
)

type Iterator struct {
	wordScanner *bufio.Scanner
	lineScanner *bufio.Scanner
	wordSeq     []string
	order       int
	err         error
}

func NewIterator(order int, reader io.Reader) *Iterator {
	lineScanner := bufio.NewScanner(reader)

	wordSeq := make([]string, order)
	for i := 0; i < len(wordSeq); i++ {
		wordSeq[i] = ""
	}

	iter := &Iterator{
		lineScanner: lineScanner,
		wordSeq:     wordSeq,
		order:       order,
	}

	return iter
}

func (iter *Iterator) Order() int {
	return iter.order
}

func (iter *Iterator) Iterate() bool {
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

		iter.wordSeq = make([]string, iter.order)
		for i := 0; i < len(iter.wordSeq); i++ {
			iter.wordSeq[i] = ""
		}
	}

	return hasNext
}

func (iter *Iterator) Get() []string {
	wordSeq := make([]string, iter.order)
	copy(wordSeq, iter.wordSeq)

	return wordSeq
}

func (iter *Iterator) Error() error {
	return iter.err
}
