package data

import (
	"testing"

	"github.com/mitsuse/kompl/trie"
)

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
