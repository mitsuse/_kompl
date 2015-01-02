package trie

import (
	"testing"
)

func TestNew(t *testing.T) {
	node := New()

	if char := node.Char(); char != 0 {
		template := "The default value of \"(*Trie).char\" should be %d, but is %d."
		t.Errorf(template, 0, char)
		return
	}

	if value := node.Value; value != 0 {
		template := "The default value of \"(*Trie).value\" should be %d, but is %d."
		t.Errorf(template, 0, value)
		return
	}

	if size := len(node.childSeq); size != 0 {
		template := "\"(*Trie).childSeq\" should be empty, but has %d element(s)."
		t.Errorf(template, size)
		return
	}
}

type KeyValueTest struct {
	Key   string
	Value int
	Exist bool
}

func createAddTestSeq() []*KeyValueTest {
	testSeq := []*KeyValueTest{
		&KeyValueTest{Key: "aaaaa", Value: 0, Exist: false},
		&KeyValueTest{Key: "aabaa", Value: 1, Exist: false},
		&KeyValueTest{Key: "aacaa", Value: 2, Exist: false},
		&KeyValueTest{Key: "a", Value: 3, Exist: true},
		&KeyValueTest{Key: "aaaba", Value: 4, Exist: false},
		&KeyValueTest{Key: "ccccc", Value: 5, Exist: false},
		&KeyValueTest{Key: "ccccccc", Value: 6, Exist: false},
		&KeyValueTest{Key: "bbb", Value: 7, Exist: false},
		&KeyValueTest{Key: "aabaa", Value: 8, Exist: true},
	}

	return testSeq
}

func TestAdd(t *testing.T) {
	rootNode := New()

	for _, test := range createAddTestSeq() {
		node, found := rootNode.Add([]int32(test.Key))

		if found != test.Exist {
			template := "The node corresposing to \"%s\" shouldn't have existed."
			t.Errorf(template, test.Key)
			return
		}

		node.Value = test.Value
	}
}
