package trie

import (
	"io"

	"github.com/mitsuse/kompl/binary"
)

func Dump(t *Trie, writer io.Writer) error {
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
	errWriter := binary.NewWriter(writer)

	errWriter.Write(t.char)
	errWriter.Write(t.Value)
	errWriter.Write(int64(t.Value))
	errWriter.Write(int64(len(t.childSeq)))

	if err := errWriter.Error(); err != nil {
		return err
	}

	return nil
}
