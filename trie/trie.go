package trie

import (
	"encoding/binary"
	"io"
	"sort"
)

type Trie struct {
	char     int32
	childSeq TrieSeq
	Value    int
}

func New() (t *Trie) {
	t = &Trie{
		char:     0,
		childSeq: make(TrieSeq, 0),
		Value:    0,
	}

	return t
}

func Inflate(reader io.Reader) (*Trie, error) {
	t, err := inflateNode(reader)
	if err != nil {
		return nil, err
	}

	nodeStack := []*Trie{t}
	offsetStack := []int{0}

	for len(nodeStack) > 0 {
		node := nodeStack[len(nodeStack)-1]
		offset := offsetStack[len(offsetStack)-1]

		if offset < node.childSeq.Len() {
			child, err := inflateNode(reader)
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

func inflateNode(reader io.Reader) (*Trie, error) {
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

func (t *Trie) Deflate(writer io.Writer) error {
	nodeStack := []*Trie{t}
	offsetStack := []int{-1}

	if err := t.deflateNode(writer); err != nil {
		return err
	}

	for len(nodeStack) > 0 {
		offsetStack[len(offsetStack)-1]++
		node := nodeStack[len(nodeStack)-1]
		offset := offsetStack[len(offsetStack)-1]

		if offset < node.childSeq.Len() {
			child := node.childSeq[offset]
			if err := child.deflateNode(writer); err != nil {
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

func (t *Trie) deflateNode(writer io.Writer) error {
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

func (t *Trie) Get(key []int32) (node *Trie, exist bool) {
	node, read := t.find(key)

	return node, len(key) == read
}

func (t *Trie) Add(key []int32) (node *Trie, found bool) {
	node, read := t.find(key)
	found = len(key) == read

	for ; read < len(key); read++ {
		child := &Trie{
			char:     key[read],
			childSeq: make(TrieSeq, 0),
			Value:    0,
		}

		node.childSeq = append(node.childSeq, child)
		sort.Sort(node.childSeq)

		node = child
	}

	return node, found
}

func (t *Trie) FindMax(f func(x, y int) bool) *Trie {
	var maxChild *Trie = nil

	for _, child := range t.childSeq {
		if maxChild == nil {
			maxChild = child
			continue
		}

		if f(maxChild.Value, child.Value) {
			maxChild = child
		}
	}

	return maxChild
}

func (t *Trie) find(key []int32) (node *Trie, read int) {
	node = t

	for keyOffset, char := range key {
		if len(node.childSeq) == 0 {
			return node, keyOffset
		}

		childOffset := sort.Search(len(node.childSeq)-1, func(offset int) bool {
			return node.childSeq[offset].char >= char
		})
		child := node.childSeq[childOffset]

		if child.char != char {
			return node, keyOffset
		}

		node = child
	}

	return node, len(key)
}

func (t *Trie) Iter() *TrieIter {
	iter := &TrieIter{
		nodeSeq:   []*Trie{t},
		offsetSeq: []int{-1},
	}

	return iter
}
