package compl

import (
	"bufio"
	"errors"
	"io"
	"strconv"
	"strings"
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
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		text := scanner.Text()

		textSplit := strings.Split(text, "\t")
		if len(textSplit) != 2 {
			// TODO: Write the error message.
			return nil, errors.New("")
		}

		ngram := textSplit[0]
		wordSeq := strings.Split(ngram, " ")

		count, err := strconv.Atoi(textSplit[1])
		if err != nil {
			return nil, err
		}

		_ = wordSeq
		_ = count
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

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
