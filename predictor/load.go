package predictor

import (
	"io"

	"github.com/mitsuse/kompl/binary"
	"github.com/mitsuse/kompl/predictor/data"
	"github.com/mitsuse/kompl/trie"
)

func Load(reader io.Reader) (*Predictor, error) {
	var order int64
	var wordSize int64

	errReader := binary.NewReader(reader)
	errReader.Read(&order)
	errReader.Read(&wordSize)
	valueSeq := loadValueSeq(errReader)

	if err := errReader.Error(); err != nil {
		return nil, err
	}

	wordTrie, err := trie.Load(reader)
	if err != nil {
		return nil, err
	}

	ngramTrie, err := trie.Load(reader)
	if err != nil {
		return nil, err
	}

	p := &Predictor{
		order:     int(order),
		wordSize:  int(wordSize),
		wordTrie:  wordTrie,
		ngramTrie: ngramTrie,
		valueSeq:  valueSeq,
	}

	return p, nil
}

func loadValueSeq(reader *binary.Reader) []*data.Value {
	var valueSeqSize int64
	reader.Read(&valueSeqSize)

	valueSeq := make([]*data.Value, valueSeqSize)
	for i := 0; i < len(valueSeq); i++ {
		valueSeq[i] = loadValue(reader)
	}

	return valueSeq
}

func loadValue(reader *binary.Reader) *data.Value {
	var count int64
	var maxCount int64
	var first int64
	var sibling int64

	reader.Read(&count)
	reader.Read(&maxCount)
	reader.Read(&first)
	reader.Read(&sibling)

	value := &data.Value{
		Count:    int(count),
		MaxCount: int(maxCount),
		First:    int(first),
		Sibling:  int(sibling),
	}

	return value
}
