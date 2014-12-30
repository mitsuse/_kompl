package trie

import (
	"encoding/binary"
	"io"
)

func Load(reader io.Reader) (*Trie, error) {
	t, err := loadNode(reader)
	if err != nil {
		return nil, err
	}

	nodeStack := []*Trie{t}
	offsetStack := []int{0}

	for len(nodeStack) > 0 {
		node := nodeStack[len(nodeStack)-1]
		offset := offsetStack[len(offsetStack)-1]

		if offset < node.childSeq.Len() {
			child, err := loadNode(reader)
			if err != nil {
				return nil, err
			}

			node.childSeq[offset] = child
			offset++
			offsetStack[len(offsetStack)-1] = offset

			nodeStack = append(nodeStack, child)
			offsetStack = append(offsetStack, 0)
		} else {
			nodeStack = nodeStack[:len(nodeStack)-1]
			offsetStack = offsetStack[:len(offsetStack)-1]
		}
	}

	return t, nil
}

func loadNode(reader io.Reader) (*Trie, error) {
	endian := binary.LittleEndian

	var char int32
	if err := binary.Read(reader, endian, &char); err != nil {
		return nil, err
	}

	var value int64
	if err := binary.Read(reader, endian, &value); err != nil {
		return nil, err
	}

	var childSeqSize int64
	if err := binary.Read(reader, endian, &childSeqSize); err != nil {
		return nil, err
	}

	t := &Trie{
		char:     char,
		childSeq: make([]*Trie, childSeqSize),
		Value:    int(value),
	}

	return t, nil
}
