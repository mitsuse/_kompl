package predictor

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

type predictTest struct {
	Context []string
	Prefix  string
	K       int
	CandSeq []string
}

func createReader() io.Reader {
	byteSeq, _ := Asset("predictor/test/wiki.txt")

	return bytes.NewReader(byteSeq)
}

func createPredictTestSeq() []*predictTest {
	testSeq := []*predictTest{
		&predictTest{
			Context: []string{"is"},
			Prefix:  "",
			K:       10,
			CandSeq: []string{"a", "now"},
		},
	}

	return testSeq
}

func TestPredictorBuildSucceed(t *testing.T) {
	order := 3
	reader := createReader()

	p, err := Build(order, reader)
	if err != nil {
		template := "Failed to build a predictor: %s"
		t.Errorf(template, err.Error())
		return
	}

	for _, test := range createPredictTestSeq() {
		candSeq := p.Predict(test.Context, test.Prefix, test.K)
		if !compareCandSeq(candSeq, test.CandSeq) {
			descTemplate := "context = %v, prefix = %v, k = %d"
			desc := fmt.Sprintf(descTemplate, test.Context, test.Prefix, test.K)

			template := "The predictor should return %v, but %v: %s"
			t.Errorf(template, test.CandSeq, candSeq, desc)
			return
		}
	}
}

func compareCandSeq(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}

	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}

	return true
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
