package ngram

import (
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
