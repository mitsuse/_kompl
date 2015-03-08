package predictor

import (
	"github.com/mitsuse/kompl/predictor/data"
	"github.com/mitsuse/kompl/trie"
)

const (
	_UNKNOWN_SYMBOL = -2
)

type Predictor struct {
	order     int
	wordSize  int
	wordTrie  *trie.Trie
	ngramTrie *trie.Trie
	valueSeq  []*data.Value
}

func (p *Predictor) Order() int {
	return p.order
}

func (p *Predictor) Predict(context []string, prefix string, k int) []string {
	candidateSet := make(map[string]struct{})
	candidateSeq := make([]string, 0, k)

	contextKey := p.encode(context, prefix)
	for start := 0; start <= len(context) && len(candidateSeq) < k; start++ {
		prefixNode, exist := p.ngramTrie.Get(contextKey[start:])
		if !exist {
			continue
		}

		queue := data.NewQueue()

		value := p.valueSeq[prefixNode.Value]
		candidate := data.NewCandidate(prefix, prefixNode, nil, value.Count)
		queue.Push(candidate)

		for queue.Len() > 0 {
			candidate, _ := queue.Pop()

			if first, exist := p.getFirst(candidate); exist {
				queue.Push(first)
			}

			if sibling, exist := p.getSibgling(candidate); exist {
				queue.Push(sibling)
			}

			candidateValue := p.valueSeq[candidate.Node().Value-1]
			if candidateValue.Count > 0 {
				word := candidate.Word()

				_, exist := candidateSet[word]
				if exist {
					continue
				}
				candidateSet[word] = struct{}{}

				candidateSeq = append(candidateSeq, word)
				if len(candidateSeq) == k {
					break
				}
			}
		}
	}

	return candidateSeq
}

func (p *Predictor) encode(context []string, prefix string) []int32 {
	if len(context) >= p.Order() {
		context = context[len(context)-p.Order()+1:]
	}

	key := make([]int32, 0, len(context)+len(prefix))

	for _, word := range context {
		node, exist := p.wordTrie.Get([]int32(word))
		if !exist {
			key = append(key, _UNKNOWN_SYMBOL)
		} else {
			key = append(key, int32(node.Value))
		}
	}

	key = append(key, _END_OF_CONTEXT)

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
