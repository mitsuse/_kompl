package compl

import (
	"bufio"
	"encoding/binary"
	"errors"
	"io"
	"strconv"
	"strings"

	"github.com/mitsuse/compl/trie"
)

type Model struct {
	wordSize  int
	wordTrie  *trie.Trie
	ngramTrie *trie.Trie
	valueSeq  []*Value
}

func InflateModel(reader io.Reader) (*Model, error) {
	// TODO: Deserialize a completion model from file.
	m := &Model{}

	return m, nil
}

func InflateRawModel(reader io.Reader) (*Model, error) {
	// TODO: Convert a raw count file into a model for Compl server.
	m := &Model{
		wordSize:  0,
		wordTrie:  trie.New(),
		ngramTrie: trie.New(),
	}

	if err := m.inflateRaw(reader); err != nil {
		return nil, err
	}

	return m, nil
}

func (m *Model) inflateRaw(reader io.Reader) error {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		text := scanner.Text()

		wordSeq, count, err := m.processRawLine(text)
		if err != nil {
			return err
		}

		key := m.encodeNew(wordSeq)

		node, exist := m.ngramTrie.Add(key)
		if !exist {
			value := &Value{
				Count:    count,
				MaxCount: 0,
			}

			m.valueSeq = append(m.valueSeq, value)
			node.Value = len(m.valueSeq)
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	m.fillMaxScore()

	return nil
}

func (m *Model) processRawLine(text string) (wordSeq []string, count int, err error) {
	textSplit := strings.Split(text, "\t")
	if len(textSplit) != 2 {
		// TODO: Write the error message.
		err = errors.New("")
		return
	}

	ngram := textSplit[0]
	wordSeq = strings.Split(ngram, " ")

	count, err = strconv.Atoi(textSplit[1])
	if err != nil {
		return
	}

	return
}

func (m *Model) encodeNew(wordSeq []string) (encodedSeq []int32) {
	encodedSeq = []int32{}

	// Encode only context words with "wordTrie".
	for i := 0; i < len(wordSeq)-1; i++ {
		charSeq := []int32(wordSeq[i])
		if node, exist := m.wordTrie.Add(charSeq); !exist {
			m.wordSize++
			node.Value = m.wordSize
		}

		encodedSeq = append(encodedSeq, charSeq...)
	}

	charSeq := []int32(wordSeq[len(wordSeq)-1])
	encodedSeq = append(encodedSeq, charSeq...)

	return
}

func (m *Model) fillMaxScore() {
	iter := m.ngramTrie.Iter()
	for iter.HasNext() {
		node := iter.Get()
		if node.Value == 0 {
			value := &Value{
				Count:    0,
				MaxCount: 0,
			}

			m.valueSeq = append(m.valueSeq, value)
			node.Value = len(m.valueSeq)
		}

		maxChild := node.FindMax(func(x, y int) bool {
			return m.valueSeq[x].Count-m.valueSeq[y].Count < 0
		})

		if maxChild == nil {
			m.valueSeq[node.Value-1].MaxCount = m.valueSeq[node.Value-1].Count
		} else {
			m.valueSeq[node.Value-1].MaxCount = m.valueSeq[maxChild.Value-1].Count
		}
	}
}

func (m *Model) Deflate(writer io.Writer) error {
	if err := binary.Write(writer, binary.LittleEndian, int64(m.wordSize)); err != nil {
		return err
	}

	if err := m.wordTrie.Deflate(writer); err != nil {
		return err
	}

	if err := m.ngramTrie.Deflate(writer); err != nil {
		return err
	}

	valueSeqSize := int64(m.valueSeq)
	if err := binary.Write(writer, binary.LittleEndian, valueSeqSize); err != nil {
		return err
	}

	for _, value := range m.valueSeq {
		err := binary.Write(writer, binary.LittleEndian, int64(value.Count))
		if err != nil {
			return err
		}

		err = binary.Write(writer, binary.LittleEndian, int64(value.MaxCount))
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *Model) Predict(context []string, prefix string, k int) []string {
	// TODO: Predict the next word.
	candSeq := []string{}

	return candSeq
}

type Value struct {
	Count    int
	MaxCount int
}
