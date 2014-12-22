package kompl

import (
	"encoding/binary"
	"io"

	"github.com/mitsuse/kompl/trie"
)

type Predictor struct {
	wordSize  int
	wordTrie  *trie.Trie
	ngramTrie *trie.Trie
	valueSeq  []*Value
}

func (p *Predictor) Dump(writer io.Writer) error {
	if err := binary.Write(writer, binary.LittleEndian, int64(p.wordSize)); err != nil {
		return err
	}

	if err := p.wordTrie.Deflate(writer); err != nil {
		return err
	}

	if err := p.ngramTrie.Deflate(writer); err != nil {
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

func (p *Predictor) Predict(context []string, prefix string, k int) []string {
	candSeq := []string{}

	contextKey := p.encode(context, prefix)

	prefixNode, exist := p.ngramTrie.Get(contextKey)
	if !exist {
		return candSeq
	}

	return p.generateCandidates(prefix, prefixNode, k)
}

func (p *Predictor) encode(context []string, prefix string) []int32 {
	key := make([]int32, 0, len(context)+len(prefix))

	for _, word := range context {
		node, exist := p.wordTrie.Get([]int32(word))
		if !exist {
			key = append(key, 0)
		} else {
			key = append(key, int32(node.Value))
		}
	}

	for _, char := range []int32(prefix) {
		key = append(key, char)
	}

	return key
}

func (p *Predictor) generateCandidates(prefix string, node *trie.Trie, k int) []string {
	candidateSeq := make([]string, 0, k)

	queue := NewQueue()

	value := p.valueSeq[node.Value]
	candidate := NewCandidate(prefix, node, nil, value.Count)
	queue.Push(candidate)

	for queue.Len() > 0 {
		candidate, _ := queue.Pop()

		candidateValue := p.valueSeq[candidate.Node().Value-1]
		if candidateValue.Count > 0 {
			candidateSeq = append(candidateSeq, candidate.Word())
			if len(candidateSeq) == k {
				break
			}
		}

		if first, exist := p.getFirst(candidate); exist {
			queue.Push(first)
		}

		if sibling, exist := p.getSibgling(candidate); exist {
			queue.Push(sibling)
		}
	}

	return candidateSeq
}

func (p *Predictor) getFirst(candidate *Candidate) (*Candidate, bool) {
	parentValue := p.valueSeq[candidate.node.Value-1]

	childNode, exist := candidate.node.GetChildByOffset(parentValue.First)
	if !exist {
		return nil, false
	}

	childValue := p.valueSeq[childNode.Value-1]
	childWord := string(append([]int32(candidate.word), childNode.Char()))

	childCandidate := NewCandidate(
		childWord,
		childNode,
		candidate.node,
		childValue.Count,
	)

	return childCandidate, true
}

func (p *Predictor) getSibgling(candidate *Candidate) (*Candidate, bool) {
	nodeValue := p.valueSeq[candidate.node.Value-1]

	if candidate.parent == nil {
		return nil, false
	}

	siblingNode, exist := candidate.parent.GetChildByOffset(nodeValue.Sibling)
	if !exist {
		return nil, false
	}

	nodeWord := []int32(candidate.word)

	siblingValue := p.valueSeq[siblingNode.Value-1]
	siblingWord := string(append(nodeWord[:len(nodeWord)-1], siblingNode.Char()))

	siblingCandidate := NewCandidate(
		siblingWord,
		siblingNode,
		candidate.parent,
		siblingValue.Count,
	)

	return siblingCandidate, true
}

type Value struct {
	Count    int
	MaxCount int
	First    int
	Sibling  int
}

type IndexedNode struct {
	Node  *trie.Trie
	Index int
}

type IndexedNodeSeq struct {
	seq      []*IndexedNode
	valueSeq []*Value
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
