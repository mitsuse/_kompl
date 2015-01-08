package data

import (
	"fmt"
	"testing"

	"github.com/mitsuse/kompl/trie"
)

func TestIndexedNodeSeqLen(t *testing.T) {
	seq := NewIndexedNodeSeq([]*Value{})

	aNode := &IndexedNode{Node: trie.New()}
	seq.Append(aNode)

	bNode := &IndexedNode{Node: trie.New()}
	seq.Append(bNode)

	cNode := &IndexedNode{Node: trie.New()}
	seq.Append(cNode)

	if length := seq.Len(); length != 3 {
		template := "\"(*IndexedNodeSeq).Len\" should return %p, but returns %d."
		t.Errorf(template, 3, length)
		return
	}
}

func TestIndexedNodeSeqLess(t *testing.T) {
	valueSeq := []*Value{
		&Value{Count: 0},
		&Value{Count: 10},
		&Value{Count: 0},
		&Value{Count: 0},
		&Value{Count: 40},
		&Value{Count: 0},
		&Value{Count: 0},
		&Value{Count: 0},
		&Value{Count: 0},
		&Value{Count: 0},
		&Value{Count: 100},
		&Value{Count: 0},
	}

	seq := NewIndexedNodeSeq(valueSeq)

	aNode := &IndexedNode{Node: trie.New()}
	aNode.Node.Value = 10
	seq.Append(aNode)

	bNode := &IndexedNode{Node: trie.New()}
	bNode.Node.Value = 1
	seq.Append(bNode)

	cNode := &IndexedNode{Node: trie.New()}
	cNode.Node.Value = 4
	seq.Append(cNode)

	if seq.Less(0, 1) {
		template := "%s shouldn't be smaller than %s."
		first := "The value of the first node, which is %d"
		second := "The value of the second node, which is %d"
		t.Errorf(template, fmt.Sprintf(first, 10), fmt.Sprintf(second, 1))
		return
	}
}

func TestIndexedNodeSeqSwap(t *testing.T) {
	seq := NewIndexedNodeSeq([]*Value{})

	aNode := &IndexedNode{Node: trie.New()}
	aNode.Node.Value = 10
	seq.Append(aNode)

	bNode := &IndexedNode{Node: trie.New()}
	bNode.Node.Value = 1
	seq.Append(bNode)

	cNode := &IndexedNode{Node: trie.New()}
	cNode.Node.Value = 4
	seq.Append(cNode)

	seq.Swap(0, 1)

	if node := seq.Get(0); node != bNode {
		template := "The first node should be %p, but is %p."
		t.Errorf(template, bNode, node)
		return
	}

	if node := seq.Get(1); node != aNode {
		template := "The second node should be %p, but is %p."
		t.Errorf(template, aNode, node)
		return
	}
}

func TestIndexedNodeSeqAppend(t *testing.T) {
	seq := NewIndexedNodeSeq([]*Value{})

	if length := len(seq.seq); length > 0 {
		template := "The length of \"%s\" should be %d intially, but is %d."
		structName := "(*IndexedNodeSeq)"
		t.Errorf(template, structName, 0, length)
		return
	}

	node := &IndexedNode{Node: trie.New()}

	iteration := 10
	for i := 0; i < iteration; i++ {
		seq.Append(node)
	}

	if length := len(seq.seq); length != iteration {
		template := "The length of \"%s\" should be %d after appending, but is %d."
		structName := "(*IndexedNodeSeq)"
		t.Errorf(template, structName, 0, length)
		return
	}

	for _, n := range seq.seq {
		if n != node {
			template := "The appended node should be %p, but is %p."
			t.Errorf(template, node, n)
			return
		}
	}
}

func TestIndexedNodeSeqGet(t *testing.T) {
	seq := NewIndexedNodeSeq([]*Value{})

	aNode := &IndexedNode{Node: trie.New()}
	seq.Append(aNode)

	bNode := &IndexedNode{Node: trie.New()}
	seq.Append(bNode)

	cNode := &IndexedNode{Node: trie.New()}
	seq.Append(cNode)

	if node := seq.Get(1); bNode != node {
		template := "\"(*IndexedNodeSeq).Get\" should return a node %p, but returns %p."
		t.Errorf(template, bNode, node)
		return
	}
}
