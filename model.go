package compl

import (
	"io"
)

type Model struct {
}

func InflateModel(reader io.Reader) (*Model, error) {
	// TODO: Deserialize a completion model from file.
	m := &Model{}

	return m, nil
}

func InflateArpaModel(reader io.Reader) (*Model, error) {
	// TODO: Convert an ARPA-formatted model into a model.for Compl server.
	m := &Model{}

	return m, nil
}

func (m *Model) Predict(context []string, prefix string, k int) []string {
	// TODO: Predict the next word.
	candSeq := []string{}

	return candSeq
}
