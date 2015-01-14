package predictor

import (
	"testing"
)

func TestPredictorOrder(t *testing.T) {
	expectedOrder := 3

	p := &Predictor{
		order: expectedOrder,
	}

	if order := p.Order(); order != expectedOrder {
		template := "(*Predictor).Order should return %d, but returns %d."
		t.Errorf(template, expectedOrder, order)
		return
	}
}
