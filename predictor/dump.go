package predictor

import (
	"io"

	"github.com/mitsuse/kompl/binary"
	"github.com/mitsuse/kompl/trie"
)

func Dump(p *Predictor, writer io.Writer) error {
	errWriter := binary.NewWriter(writer)

	errWriter.Write(int64(p.order))
	errWriter.Write(int64(p.wordSize))
	errWriter.Write(int64(len(p.valueSeq)))

	for _, value := range p.valueSeq {
		errWriter.Write(int64(value.Count))
		errWriter.Write(int64(value.MaxCount))
		errWriter.Write(int64(value.First))
		errWriter.Write(int64(value.Sibling))
	}

	if err := errWriter.Error(); err != nil {
		return err
	}

	if err := trie.Dump(p.wordTrie, writer); err != nil {
		return err
	}

	if err := trie.Dump(p.ngramTrie, writer); err != nil {
		return err
	}

	return nil
}
