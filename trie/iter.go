package trie

type TrieIter struct {
	nodeSeq   []*Trie
	offsetSeq []int
	node      *Trie
}

func (iter *TrieIter) HasNext() bool {
	if len(iter.nodeSeq) <= 0 {
		return false
	}

	for {
		node := iter.nodeSeq[len(iter.nodeSeq)-1]

		iter.offsetSeq[len(iter.offsetSeq)-1]++
		offset := iter.offsetSeq[len(iter.offsetSeq)-1]

		if offset < node.childSeq.Len() {
			iter.nodeSeq = append(iter.nodeSeq, node.childSeq[offset])
			iter.offsetSeq = append(iter.offsetSeq, -1)
		} else {
			iter.node = node
			iter.nodeSeq = iter.nodeSeq[:len(iter.nodeSeq)-1]
			iter.offsetSeq = iter.offsetSeq[:len(iter.offsetSeq)-1]
			break
		}
	}

	return true
}

func (iter *TrieIter) Get() *Trie {
	return iter.node
}

func (iter *TrieIter) Error() error {
	// TODO: Implement this.
	return nil
}

type NodeIter struct {
}

func (iter *NodeIter) HasNext() bool {
	// TODO: Implement this.
	return false
}

func (iter *NodeIter) Get() *Trie {
	// TODO: Implement this.
	return nil
}
