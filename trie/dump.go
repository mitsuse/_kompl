package trie

import (
	"encoding/binary"
	"io"
)

func (t *Trie) Dump(writer io.Writer) error {
	nodeStack := []*Trie{t}
	offsetStack := []int{-1}

	if err := t.dumpNode(writer); err != nil {
		return err
	}

	for len(nodeStack) > 0 {
		offsetStack[len(offsetStack)-1]++
		node := nodeStack[len(nodeStack)-1]
		offset := offsetStack[len(offsetStack)-1]

		if offset < node.childSeq.Len() {
			child := node.childSeq[offset]
			if err := child.dumpNode(writer); err != nil {
				return err
			}

			nodeStack = append(nodeStack, child)
			offsetStack = append(offsetStack, -1)
		} else {
			nodeStack = nodeStack[:len(nodeStack)-1]
			offsetStack = offsetStack[:len(offsetStack)-1]
		}
	}

	return nil
}

func (t *Trie) dumpNode(writer io.Writer) error {
	endian := binary.LittleEndian

	if err := binary.Write(writer, endian, t.char); err != nil {
		return err
	}

	if err := binary.Write(writer, endian, int64(t.Value)); err != nil {
		return err
	}

	if err := binary.Write(writer, endian, int64(len(t.childSeq))); err != nil {
		return err
	}

	return nil
}
