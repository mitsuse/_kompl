package predictor

import (
	"github.com/mitsuse/kompl/predictor/data"
	"github.com/mitsuse/kompl/trie"
)

type Predictor struct {
	wordSize  int
	wordTrie  *trie.Trie
	ngramTrie *trie.Trie
	valueSeq  []*data.Value
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

	queue := data.NewQueue()

	value := p.valueSeq[node.Value]
	candidate := data.NewCandidate(prefix, node, nil, value.Count)
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

func (p *Predictor) getFirst(candidate *data.Candidate) (*data.Candidate, bool) {
	parentValue := p.valueSeq[candidate.Node().Value-1]

	childNode, exist := candidate.Node().GetChildByOffset(parentValue.First)
	if !exist {
		return nil, false
	}

	childValue := p.valueSeq[childNode.Value-1]
	childWord := string(append([]int32(candidate.Word()), childNode.Char()))

	childCandidate := data.NewCandidate(
		childWord,
		childNode,
		candidate.Node(),
		childValue.Count,
	)

	return childCandidate, true
}

func (p *Predictor) getSibgling(candidate *data.Candidate) (*data.Candidate, bool) {
	nodeValue := p.valueSeq[candidate.Node().Value-1]

	if candidate.Parent() == nil {
		return nil, false
	}

	siblingNode, exist := candidate.Parent().GetChildByOffset(nodeValue.Sibling)
	if !exist {
		return nil, false
	}

	nodeWord := []int32(candidate.Word())

	siblingValue := p.valueSeq[siblingNode.Value-1]
	siblingWord := string(append(nodeWord[:len(nodeWord)-1], siblingNode.Char()))

	siblingCandidate := data.NewCandidate(
		siblingWord,
		siblingNode,
		candidate.Parent(),
		siblingValue.Count,
	)

	return siblingCandidate, true
}
