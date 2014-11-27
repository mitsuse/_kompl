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

func InflateRawModel(reader io.Reader) (*Model, error) {
	// TODO: Convert a raw count file into a model for Compl server.
	m := &Model{}

	return m, nil
}

func (m *Model) Deflate(writer io.Writer) error {
	// TODO: Write this model into writer.
	return nil
}

func (m *Model) Predict(context []string, prefix string, k int) []string {
	// TODO: Predict the next word.
	candSeq := []string{}

	return candSeq
}
