package kompl

import (
	"encoding/binary"
	"io"

	"github.com/mitsuse/kompl/trie"
)

func LoadPredictor(reader io.Reader) (*Predictor, error) {
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

	valueSeq, err := loadValueSeq(reader)
	if err != nil {
		return nil, err
	}

	p := &Predictor{
		wordSize:  int(wordSize),
		wordTrie:  wordTrie,
		ngramTrie: ngramTrie,
		valueSeq:  valueSeq,
	}

	return p, nil
}

func loadValueSeq(reader io.Reader) ([]*Value, error) {
	var valueSeqSize int64
	if err := binary.Read(reader, binary.LittleEndian, &valueSeqSize); err != nil {
		return nil, err
	}

	valueSeq := make([]*Value, valueSeqSize)
	for i := 0; i < len(valueSeq); i++ {
		value, err := loadValue(reader)
		if err != nil {
			return nil, err
		}

		valueSeq[i] = value
	}

	return valueSeq, nil
}

func loadValue(reader io.Reader) (*Value, error) {
	var count int64
	var maxCount int64
	var first int64
	var sibling int64

	if err := binary.Read(reader, binary.LittleEndian, &count); err != nil {
		return nil, err
	}

	if err := binary.Read(reader, binary.LittleEndian, &maxCount); err != nil {
		return nil, err
	}

	if err := binary.Read(reader, binary.LittleEndian, &first); err != nil {
		return nil, err
	}

	if err := binary.Read(reader, binary.LittleEndian, &sibling); err != nil {
		return nil, err
	}

	value := &Value{
		Count:    int(count),
		MaxCount: int(maxCount),
		First:    int(first),
		Sibling:  int(sibling),
	}

	return value, nil
}
