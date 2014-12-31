package predictor

import (
	"encoding/binary"
	"io"

	"github.com/mitsuse/kompl/predictor/data"
	"github.com/mitsuse/kompl/trie"
)

func Dump(p *Predictor, writer io.Writer) error {
	if err := binary.Write(writer, binary.LittleEndian, int64(p.order)); err != nil {
		return err
	}

	if err := binary.Write(writer, binary.LittleEndian, int64(p.wordSize)); err != nil {
		return err
	}

	if err := dumpValueSeq(p.valueSeq, writer); err != nil {
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

func dumpValueSeq(valueSeq []*data.Value, writer io.Writer) error {
	valueSeqSize := int64(len(valueSeq))
	if err := binary.Write(writer, binary.LittleEndian, valueSeqSize); err != nil {
		return err
	}

	for _, value := range valueSeq {
		if err := dumpValue(value, writer); err != nil {
			return err
		}
	}

	return nil
}

func dumpValue(value *data.Value, writer io.Writer) error {
	err := binary.Write(writer, binary.LittleEndian, int64(value.Count))
	if err != nil {
		return err
	}

	err = binary.Write(writer, binary.LittleEndian, int64(value.MaxCount))
	if err != nil {
		return err
	}

	err = binary.Write(writer, binary.LittleEndian, int64(value.First))
	if err != nil {
		return err
	}

	err = binary.Write(writer, binary.LittleEndian, int64(value.Sibling))
	if err != nil {
		return err
	}

	return nil
}
