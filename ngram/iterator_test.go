package ngram

import (
	"strings"
	"testing"
)

func TestIteratorOrder(t *testing.T) {
	order := 3

	iter := NewIterator(order, nil)
	if iter.Order() != order {
		template := "\"(*Iterator).Order\" should return \"%d\", but returns \"%d\"."
		t.Errorf(template, order, iter.Order())
	}
}

type IterateTest struct {
	Order    int
	Data     string
	NgramSeq [][]string
}

func TestIteratorIterateSucceed(t *testing.T) {
	test := &IterateTest{
		Order: 3,
		Data:  "' a ( bbb ccc ddd\n; eee",
		NgramSeq: [][]string{
			[]string{"", "", "'"},
			[]string{"", "'", "a"},
			[]string{"'", "a", "("},
			[]string{"a", "(", "bbb"},
			[]string{"(", "bbb", "ccc"},
			[]string{"bbb", "ccc", "ddd"},
			[]string{"", "", ";"},
			[]string{"", ";", "eee"},
		},
	}

	iter := NewIterator(test.Order, strings.NewReader(test.Data))
	index := 0

	for iter.Iterate() {
		if ngram := iter.Get(); !ngramEq(ngram, test.NgramSeq[index]) {
			template := "\"(*Iterator).Get\" should return %v, but returns %v."
			t.Errorf(template, test.NgramSeq[index], ngram)
		}

		index++
	}

	if err := iter.Error(); err != nil {
		template := "\"(*Iterator).Iterate\" should finished with no error: %s"
		t.Errorf(template, err.Error())
		return
	}
}

func ngramEq(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}

	for i := 0; i < len(x); i++ {
		if x[i] != y[i] {
			return false
		}
	}

	return true
}
