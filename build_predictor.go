package kompl

import (
	"io"
	"sort"

	"github.com/mitsuse/kompl/trie"
)

// TODO: Get the order of N-gram as a argument..
func BuildPredictor(reader io.Reader) (*Predictor, error) {
	p := &Predictor{
		wordSize:  0,
		wordTrie:  trie.New(),
		ngramTrie: trie.New(),
	}

	iterator := NewNgramIterator(3, reader)
	for iterator.Iterate() {
		storeWordSeq(p, iterator.Get())
	}

	if err := iterator.Error(); err != nil {
		return nil, err
	}

	fillMaxScore(p)
	fillFirstAndSibling(p)

	return p, nil
}

func storeWordSeq(p *Predictor, wordSeq []string) {
	// TODO: Support for the N-grams which have start symbols as context.
	if len(wordSeq) > 0 && wordSeq[0] == "" {
		return
	}

	key := p.encodeNew(wordSeq)

	node, _ := p.ngramTrie.Add(key)
	if node.Value > 0 {
		value := p.valueSeq[node.Value-1]
		value.Count++
	} else {
		value := &Value{
			Count:    1,
			MaxCount: 0,
			First:    -1,
			Sibling:  -1,
		}

		p.valueSeq = append(p.valueSeq, value)
		node.Value = len(p.valueSeq)
	}
}

func (p *Predictor) encodeNew(wordSeq []string) (encodedSeq []int32) {
	encodedSeq = []int32{}

	// Encode only context words with "wordTrie".
	for i := 0; i < len(wordSeq)-1; i++ {
		charSeq := []int32(wordSeq[i])

		node, exist := p.wordTrie.Add(charSeq)
		if !exist {
			p.wordSize++
			node.Value = p.wordSize
		}

		encodedSeq = append(encodedSeq, int32(node.Value))
	}

	charSeq := []int32(wordSeq[len(wordSeq)-1])
	encodedSeq = append(encodedSeq, charSeq...)

	return
}

func fillMaxScore(p *Predictor) {
	iter := p.ngramTrie.Iter()
	for iter.HasNext() {
		node := iter.Get()
		if node.Value == 0 {
			value := &Value{
				Count:    0,
				MaxCount: 0,
				First:    -1,
				Sibling:  -1,
			}

			p.valueSeq = append(p.valueSeq, value)
			node.Value = len(p.valueSeq)
		}

		maxChild := node.FindMax(func(x, y int) bool {
			return p.valueSeq[x-1].Count-p.valueSeq[y-1].Count < 0
		})

		if maxChild == nil {
			p.valueSeq[node.Value-1].MaxCount = p.valueSeq[node.Value-1].Count
		} else {
			p.valueSeq[node.Value-1].MaxCount = p.valueSeq[maxChild.Value-1].Count
		}
	}
}

func fillFirstAndSibling(p *Predictor) {
	nodeStack := []*trie.Trie{p.ngramTrie}

	for len(nodeStack) > 0 {
		node := nodeStack[len(nodeStack)-1]
		nodeStack = nodeStack[:len(nodeStack)-1]

		indexedChildSeq := &IndexedNodeSeq{
			seq:      []*IndexedNode{},
			valueSeq: p.valueSeq,
		}

		iter := node.ChildIter()
		offset := 0

		for iter.HasNext() {
			child := iter.Get()
			nodeStack = append(nodeStack, child)

			indexedChild := &IndexedNode{
				Node:  child,
				Index: offset,
			}
			indexedChildSeq.seq = append(indexedChildSeq.seq, indexedChild)

			offset++
		}

		sort.Sort(indexedChildSeq)

		if indexedChildSeq.Len() > 0 {
			previousChild := indexedChildSeq.seq[0]
			p.valueSeq[node.Value-1].First = previousChild.Index

			for offset := 1; offset < indexedChildSeq.Len(); offset++ {
				indexedChild := indexedChildSeq.seq[offset]
				p.valueSeq[previousChild.Node.Value-1].Sibling = indexedChild.Index

				previousChild = indexedChild
			}
		}
	}
}
