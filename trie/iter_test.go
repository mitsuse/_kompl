package trie

import (
	"testing"
)

func TestIter(t *testing.T) {
	rootNode := New()

	for _, test := range createAddTestSeq() {
		rootNode.Add([]int32(test.Key))
	}

	iter := rootNode.Iter()

	if len(iter.nodeSeq) != 1 {
		message := "The lenght of \"(*TrieIter).nodeSeq\" should be 1 initially."
		t.Errorf(message)
		return
	}

	if len(iter.offsetSeq) != 1 {
		message := "The lenght of \"(*TrieIter).offsetSeq\" should be 1 initially."
		t.Errorf(message)
		return
	}

	if iter.node != nil {
		message := "\"(*TrieIter).node\" should be \"nil\" initially."
		t.Errorf(message)
		return
	}
}

func TestChildIter(t *testing.T) {
	rootNode := New()

	for _, test := range createAddTestSeq() {
		rootNode.Add([]int32(test.Key))
	}

	iter := rootNode.ChildIter()

	if len(iter.nodeSeq) != len(rootNode.childSeq) {
		template := "The lenght of \"(*NodeIter).nodeSeq\" should be %d initially."
		t.Errorf(template, len(rootNode.childSeq))
		return
	}

	if iter.offset != -1 {
		message := "\"(*NodeIter).offset\" should be -1 initially."
		t.Errorf(message)
		return
	}
}
