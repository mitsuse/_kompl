package compl

type Model struct {
}

func InflateModel(filePath string) (*Model, error) {
	// TODO: Deserialize a completion model from file.
	m := &Model{}

	return m, nil
}

func (m *Model) Predict(context []string, k int) []string {
	// TODO: Predict the next word.
	candSeq := []string{}

	return candSeq
}
