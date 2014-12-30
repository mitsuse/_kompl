package predictor

import (
	"encoding/binary"
	"io"
)

func Dump(p *Predictor, writer io.Writer) error {
	if err := binary.Write(writer, binary.LittleEndian, int64(p.order)); err != nil {
		return err
	}

	if err := binary.Write(writer, binary.LittleEndian, int64(p.wordSize)); err != nil {
		return err
	}

	if err := p.wordTrie.Dump(writer); err != nil {
		return err
	}

	if err := p.ngramTrie.Dump(writer); err != nil {
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

		err = binary.Write(writer, binary.LittleEndian, int64(value.First))
		if err != nil {
			return err
		}

		err = binary.Write(writer, binary.LittleEndian, int64(value.Sibling))
		if err != nil {
			return err
		}
	}

	return nil
}
