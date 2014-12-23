package data

import (
	"github.com/mitsuse/kompl/trie"
)

type IndexedNode struct {
	Node  *trie.Trie
	Index int
}

type IndexedNodeSeq struct {
	seq      []*IndexedNode
	valueSeq []*Value
}

func NewIndexedNodeSeq(valueSeq []*Value) *IndexedNodeSeq {
	s := &IndexedNodeSeq{
		seq:      []*IndexedNode{},
		valueSeq: valueSeq,
	}

	return s
}

func (s *IndexedNodeSeq) Len() int {
	return len(s.seq)
}

func (s *IndexedNodeSeq) Less(i, j int) bool {
	iCount := s.valueSeq[s.seq[i].Node.Value-1].Count
	jCount := s.valueSeq[s.seq[j].Node.Value-1].Count

	return iCount < jCount
}

func (s *IndexedNodeSeq) Swap(i, j int) {
	s.seq[i], s.seq[j] = s.seq[j], s.seq[i]
}

func (s *IndexedNodeSeq) Append(node *IndexedNode) {
	s.seq = append(s.seq, node)
}

func (s *IndexedNodeSeq) Get(i int) *IndexedNode {
	return s.seq[i]
}
