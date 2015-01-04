package trie

import (
	"testing"
)

func createIterTestSeq() []*KeyValueTest {
	testSeq := []*KeyValueTest{
		&KeyValueTest{Key: "aaaaa", Value: 0, Exist: true},
		&KeyValueTest{Key: "aaaa", Value: 0, Exist: false},
		&KeyValueTest{Key: "aaaba", Value: 4, Exist: true},
		&KeyValueTest{Key: "aaab", Value: 0, Exist: false},
		&KeyValueTest{Key: "aaa", Value: 0, Exist: false},
		&KeyValueTest{Key: "aabaa", Value: 8, Exist: true},
		&KeyValueTest{Key: "aaba", Value: 0, Exist: false},
		&KeyValueTest{Key: "aab", Value: 0, Exist: false},
		&KeyValueTest{Key: "aacaa", Value: 2, Exist: true},
		&KeyValueTest{Key: "aaca", Value: 0, Exist: false},
		&KeyValueTest{Key: "aac", Value: 0, Exist: false},
		&KeyValueTest{Key: "aa", Value: 0, Exist: false},
		&KeyValueTest{Key: "a", Value: 3, Exist: true},
		&KeyValueTest{Key: "bbb", Value: 7, Exist: true},
		&KeyValueTest{Key: "bb", Value: 0, Exist: false},
		&KeyValueTest{Key: "b", Value: 0, Exist: false},
		&KeyValueTest{Key: "ccccccc", Value: 6, Exist: true},
		&KeyValueTest{Key: "cccccc", Value: 0, Exist: false},
		&KeyValueTest{Key: "ccccc", Value: 5, Exist: true},
		&KeyValueTest{Key: "cccc", Value: 0, Exist: false},
		&KeyValueTest{Key: "ccc", Value: 0, Exist: false},
		&KeyValueTest{Key: "cc", Value: 0, Exist: false},
		&KeyValueTest{Key: "c", Value: 0, Exist: false},
		&KeyValueTest{Key: "", Value: 0, Exist: false},
	}

	return testSeq
}

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

func TestIterGet(t *testing.T) {
	rootNode := New()

	for _, test := range createAddTestSeq() {
		node, _ := rootNode.Add([]int32(test.Key))
		node.Value = test.Value
	}

	iter := rootNode.Iter()

	testSeq := createIterTestSeq()
	index := 0

	for iter.HasNext() {
		if index >= len(testSeq) {
			template := "Failed to traverse in depth-first post-order: %s"
			desc := "Visited too many nodes."
			t.Errorf(template, desc)
			return
		}

		node := iter.Get()
		test := testSeq[index]

		if node.Value != test.Value {
			template := "The node corresponding to \"%s\" should have %d, but has %d."
			t.Errorf(template, test.Key, test.Value, node.Value)
			return
		}

		index++
	}

	if err := iter.Error(); err != nil {
		template := "An error is occured on iterating trie's nodes: %s"
		t.Errorf(template, err.Error())
		return
	}

	if index < len(testSeq) {
		template := "Failed to traverse in depth-first post-order: %s"
		desc := "Visited too less nodes."
		t.Errorf(template, desc)
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

func TestChildIterGet(t *testing.T) {
	testSeq := []*KeyValueTest{
		&KeyValueTest{Key: "aaa", Value: 2, Exist: true},
		&KeyValueTest{Key: "aab", Value: 1, Exist: true},
		&KeyValueTest{Key: "aac", Value: 4, Exist: true},
		&KeyValueTest{Key: "aad", Value: 5, Exist: true},
		&KeyValueTest{Key: "aae", Value: 3, Exist: true},
	}

	rootNode := New()

	for _, test := range testSeq {
		node, _ := rootNode.Add([]int32(test.Key))
		node.Value = test.Value
	}

	internalNode, _ := rootNode.Get([]int32("aa"))

	iter := internalNode.ChildIter()
	index := 0

	for iter.HasNext() {
		if index >= len(testSeq) {
			message := "Failed to iterate children:Visited too many nodes."
			t.Errorf(message)
			return
		}

		node := iter.Get()
		test := testSeq[index]

		if node.Value != test.Value {
			template := "The node corresponding to \"%s\" should have %d, but has %d."
			t.Errorf(template, test.Key, test.Value, node.Value)
			return
		}

		index++
	}

	if index < len(testSeq) {
		message := "Failed to iterate children:Visited too less nodes."
		t.Errorf(message)
		return
	}
}
