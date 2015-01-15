package predictor

import (
	"bytes"
	"io"
	"testing"
)

func createReader() io.Reader {
	byteSeq, _ := Asset("predictor/test/lorem.txt")

	return bytes.NewReader(byteSeq)
}

func TestPredictorBuildSucceed(t *testing.T) {
	order := 3
	reader := createReader()

	_, err := Build(order, reader)
	if err != nil {
		template := "Failed to build a predictor: %s"
		t.Errorf(template, err.Error())
		return
	}
}

func TestPredictorBuildFail(t *testing.T) {
	// TODO: Implement this.
}

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
