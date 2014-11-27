package compl

import (
	"bufio"
	"errors"
	"io"
	"strconv"
	"strings"

	"github.com/mitsuse/compl/trie"
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
	lastId := 0
	wordTrie := trie.New()
	ngramTrie := trie.New()

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

		key := []int32{}

		// Encode only context words with "wordTrie".
		for i := 0; i < len(wordSeq)-1; i++ {
			charSeq := []int32(wordSeq[i])
			if node, exist := wordTrie.Add(charSeq); !exist {
				lastId++
				node.Value = lastId
			}

			key = append(key, charSeq...)
		}

		charSeq := []int32(wordSeq[len(wordSeq)-1])
		key = append(key, charSeq...)

		node, exist := ngramTrie.Add(key)
		if !exist {
			node.Value = count
		}
	}

	_ = wordTrie
	_ = ngramTrie

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
