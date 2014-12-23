package predictor

import (
	"io"
	"sort"

	"github.com/mitsuse/kompl/ngram"
	"github.com/mitsuse/kompl/predictor/data"
	"github.com/mitsuse/kompl/trie"
)

// TODO: Get the order of N-gram as a argument..
func Build(reader io.Reader) (*Predictor, error) {
	p := &Predictor{
		wordSize:  0,
		wordTrie:  trie.New(),
		ngramTrie: trie.New(),
	}

	iterator := ngram.NewIterator(3, reader)
	for iterator.Iterate() {
		// TODO: Support for the N-grams which have start symbols as context.
		wordSeq := iterator.Get()
		if len(wordSeq) > 0 && wordSeq[0] == "" {
			continue
		}

		key := encodeNew(p, wordSeq)
		storeKey(p, key)
	}

	if err := iterator.Error(); err != nil {
		return nil, err
	}

	fillMaxScore(p)
	fillFirstAndSibling(p)

	return p, nil
}

func storeKey(p *Predictor, key []int32) {
	node, _ := p.ngramTrie.Add(key)
	if node.Value > 0 {
		value := p.valueSeq[node.Value-1]
		value.Count++
	} else {
		value := &data.Value{
			Count:    1,
			MaxCount: 0,
			First:    -1,
			Sibling:  -1,
		}

		p.valueSeq = append(p.valueSeq, value)
		node.Value = len(p.valueSeq)
	}
}

func encodeNew(p *Predictor, wordSeq []string) (encodedSeq []int32) {
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
			value := &data.Value{
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

		indexedChildSeq := data.NewIndexedNodeSeq(p.valueSeq)

		iter := node.ChildIter()
		offset := 0

		for iter.HasNext() {
			child := iter.Get()
			nodeStack = append(nodeStack, child)

			indexedChild := &data.IndexedNode{
				Node:  child,
				Index: offset,
			}
			indexedChildSeq.Append(indexedChild)

			offset++
		}

		sort.Sort(indexedChildSeq)

		if indexedChildSeq.Len() > 0 {
			previousChild := indexedChildSeq.Get(0)
			p.valueSeq[node.Value-1].First = previousChild.Index

			for offset := 1; offset < indexedChildSeq.Len(); offset++ {
				indexedChild := indexedChildSeq.Get(offset)
				p.valueSeq[previousChild.Node.Value-1].Sibling = indexedChild.Index

				previousChild = indexedChild
			}
		}
	}
}
