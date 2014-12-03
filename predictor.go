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

type Predictor struct {
	wordSize  int
	wordTrie  *trie.Trie
	ngramTrie *trie.Trie
	valueSeq  []*Value
}

func InflatePredictor(reader io.Reader) (*Predictor, error) {
	// TODO: Deserialize a predictor from file.
	var wordSize int64

	if err := binary.Read(reader, binary.LittleEndian, &wordSize); err != nil {
		return nil, err
	}

	wordTrie, err := trie.Inflate(reader)
	if err != nil {
		return nil, err
	}

	ngramTrie, err := trie.Inflate(reader)
	if err != nil {
		return nil, err
	}

	var valueSeqSize int64
	if err := binary.Read(reader, binary.LittleEndian, &valueSeqSize); err != nil {
		return nil, err
	}

	valueSeq := make([]*Value, valueSeqSize)
	for i := 0; i < len(valueSeq); i++ {
		var count int64
		var maxCount int64

		if err := binary.Read(reader, binary.LittleEndian, &count); err != nil {
			return nil, err
		}

		if err := binary.Read(reader, binary.LittleEndian, &maxCount); err != nil {
			return nil, err
		}

		value := &Value{
			Count:    int(count),
			MaxCount: int(maxCount),
		}
		valueSeq[i] = value
	}

	p := &Predictor{
		wordSize:  int(wordSize),
		wordTrie:  wordTrie,
		ngramTrie: ngramTrie,
		valueSeq:  valueSeq,
	}

	return p, nil
}

func InflateRawPredictor(reader io.Reader) (*Predictor, error) {
	// TODO: Convert a raw count file into a predictor for Compl server.
	p := &Predictor{
		wordSize:  0,
		wordTrie:  trie.New(),
		ngramTrie: trie.New(),
	}

	if err := p.inflateRaw(reader); err != nil {
		return nil, err
	}

	return p, nil
}

func (p *Predictor) inflateRaw(reader io.Reader) error {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		text := scanner.Text()

		wordSeq, count, err := p.processRawLine(text)
		if err != nil {
			return err
		}

		key := p.encodeNew(wordSeq)

		node, exist := p.ngramTrie.Add(key)
		if !exist {
			value := &Value{
				Count:    count,
				MaxCount: 0,
			}

			p.valueSeq = append(p.valueSeq, value)
			node.Value = len(p.valueSeq)
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	p.fillMaxScore()

	return nil
}

func (p *Predictor) processRawLine(text string) (wordSeq []string, count int, err error) {
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

func (p *Predictor) encodeNew(wordSeq []string) (encodedSeq []int32) {
	encodedSeq = []int32{}

	// Encode only context words with "wordTrie".
	for i := 0; i < len(wordSeq)-1; i++ {
		charSeq := []int32(wordSeq[i])

		node, exist := p.wordTrie.Add(charSeq)
		if !exist {
			p.wordSize++
			node.Value = p.wordSize
		}

		encodedSeq = append(encodedSeq, int32(node.Value))
	}

	charSeq := []int32(wordSeq[len(wordSeq)-1])
	encodedSeq = append(encodedSeq, charSeq...)

	return
}

func (p *Predictor) fillMaxScore() {
	iter := p.ngramTrie.Iter()
	for iter.HasNext() {
		node := iter.Get()
		if node.Value == 0 {
			value := &Value{
				Count:    0,
				MaxCount: 0,
			}

			p.valueSeq = append(p.valueSeq, value)
			node.Value = len(p.valueSeq)
		}

		maxChild := node.FindMax(func(x, y int) bool {
			return p.valueSeq[x].Count-p.valueSeq[y].Count < 0
		})

		if maxChild == nil {
			p.valueSeq[node.Value-1].MaxCount = p.valueSeq[node.Value-1].Count
		} else {
			p.valueSeq[node.Value-1].MaxCount = p.valueSeq[maxChild.Value-1].Count
		}
	}
}

func (p *Predictor) Deflate(writer io.Writer) error {
	if err := binary.Write(writer, binary.LittleEndian, int64(p.wordSize)); err != nil {
		return err
	}

	if err := p.wordTrie.Deflate(writer); err != nil {
		return err
	}

	if err := p.ngramTrie.Deflate(writer); err != nil {
		return err
	}

	valueSeqSize := int64(len(p.valueSeq))
	if err := binary.Write(writer, binary.LittleEndian, valueSeqSize); err != nil {
		return err
	}

	for _, value := range p.valueSeq {
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

func (p *Predictor) Predict(context []string, prefix string, k int) []string {
	// TODO: Predict the next word.
	candSeq := []string{}

	return candSeq
}

type Value struct {
	Count    int
	MaxCount int
}
