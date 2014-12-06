package compl

import (
	"bufio"
	"encoding/binary"
	"errors"
	"io"
	"sort"
	"strconv"
	"strings"

	"github.com/mitsuse/compl/trie"
)

type Predictor struct {
	wordSize  int
	wordTrie  *trie.Trie
	ngramTrie *trie.Trie
	valueSeq  []*Value
}

func InflatePredictor(reader io.Reader) (*Predictor, error) {
	// TODO: Deserialize a predictor from file.
	var wordSize int64

	if err := binary.Read(reader, binary.LittleEndian, &wordSize); err != nil {
		return nil, err
	}

	wordTrie, err := trie.Inflate(reader)
	if err != nil {
		return nil, err
	}

	ngramTrie, err := trie.Inflate(reader)
	if err != nil {
		return nil, err
	}

	var valueSeqSize int64
	if err := binary.Read(reader, binary.LittleEndian, &valueSeqSize); err != nil {
		return nil, err
	}

	valueSeq := make([]*Value, valueSeqSize)
	for i := 0; i < len(valueSeq); i++ {
		var count int64
		var maxCount int64
		var first int64
		var sibling int64

		if err := binary.Read(reader, binary.LittleEndian, &count); err != nil {
			return nil, err
		}

		if err := binary.Read(reader, binary.LittleEndian, &maxCount); err != nil {
			return nil, err
		}

		if err := binary.Read(reader, binary.LittleEndian, &first); err != nil {
			return nil, err
		}

		if err := binary.Read(reader, binary.LittleEndian, &sibling); err != nil {
			return nil, err
		}

		value := &Value{
			Count:    int(count),
			MaxCount: int(maxCount),
			First:    int(first),
			Sibling:  int(maxCount),
		}
		valueSeq[i] = value
	}

	p := &Predictor{
		wordSize:  int(wordSize),
		wordTrie:  wordTrie,
		ngramTrie: ngramTrie,
		valueSeq:  valueSeq,
	}

	return p, nil
}

func InflateRawPredictor(reader io.Reader) (*Predictor, error) {
	// TODO: Convert a raw count file into a predictor for Compl server.
	p := &Predictor{
		wordSize:  0,
		wordTrie:  trie.New(),
		ngramTrie: trie.New(),
	}

	if err := p.inflateRaw(reader); err != nil {
		return nil, err
	}

	return p, nil
}

func (p *Predictor) inflateRaw(reader io.Reader) error {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		text := scanner.Text()

		wordSeq, count, err := p.processRawLine(text)
		if err != nil {
			return err
		}

		key := p.encodeNew(wordSeq)

		node, exist := p.ngramTrie.Add(key)
		if !exist {
			value := &Value{
				Count:    count,
				MaxCount: 0,
				First:    -1,
				Sibling:  -1,
			}

			p.valueSeq = append(p.valueSeq, value)
			node.Value = len(p.valueSeq)
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	p.fillMaxScore()
	p.fillFirstAndSibling()

	return nil
}

func (p *Predictor) processRawLine(text string) (wordSeq []string, count int, err error) {
	textSplit := strings.Split(text, "\t")
	if len(textSplit) != 2 {
		// TODO: Write the error message.
		err = errors.New("")
		return
	}

	ngram := textSplit[0]
	wordSeq = strings.Split(ngram, " ")

	count, err = strconv.Atoi(textSplit[1])
	if err != nil {
		return
	}

	return
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

func (p *Predictor) fillMaxScore() {
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
			return p.valueSeq[x].Count-p.valueSeq[y].Count < 0
		})

		if maxChild == nil {
			p.valueSeq[node.Value-1].MaxCount = p.valueSeq[node.Value-1].Count
		} else {
			p.valueSeq[node.Value-1].MaxCount = p.valueSeq[maxChild.Value-1].Count
		}
	}
}

func (p *Predictor) fillFirstAndSibling() {
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

func (p *Predictor) Deflate(writer io.Writer) error {
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
		if candidate.Node().Value != 0 {
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

	siblingValue := p.valueSeq[siblingNode.Value-1]
	siblingWord := string(append([]int32(candidate.word), siblingNode.Char()))

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
	iCount := s.valueSeq[s.seq[i].Node.Value].Count
	jCount := s.valueSeq[s.seq[j].Node.Value].Count

	return iCount < jCount
}

func (s *IndexedNodeSeq) Swap(i, j int) {
	s.seq[i], s.seq[j] = s.seq[j], s.seq[i]
}
