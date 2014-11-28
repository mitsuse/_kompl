package trie

import "sort"

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

func (t *Trie) Iter() (iter *TrieIter) {
	// TODO: Implement this.
	return nil
}
