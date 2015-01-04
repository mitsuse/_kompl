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

func createGetTestSeq() []*KeyValueTest {
	testSeq := []*KeyValueTest{
		&KeyValueTest{Key: "aaaaa", Value: 0, Exist: true},
		&KeyValueTest{Key: "aabaa", Value: 8, Exist: true},
		&KeyValueTest{Key: "aacaa", Value: 2, Exist: true},
		&KeyValueTest{Key: "a", Value: 3, Exist: true},
		&KeyValueTest{Key: "aaaba", Value: 4, Exist: true},
		&KeyValueTest{Key: "ccccc", Value: 5, Exist: true},
		&KeyValueTest{Key: "ccccccc", Value: 6, Exist: true},
		&KeyValueTest{Key: "bbb", Value: 7, Exist: true},
		&KeyValueTest{Key: "aa", Value: 0, Exist: true},
		&KeyValueTest{Key: "ddd", Value: 0, Exist: false},
		&KeyValueTest{Key: "aaaad", Value: 0, Exist: false},
		&KeyValueTest{Key: "cccccc", Value: 0, Exist: true},
	}

	return testSeq
}

func TestAdd(t *testing.T) {
	rootNode := New()

	for _, test := range createAddTestSeq() {
		node, found := rootNode.Add([]int32(test.Key))

		if found != test.Exist {
			template := "The node corresposing to \"%s\" shouldn't exist."
			t.Errorf(template, test.Key)
			return
		}

		node.Value = test.Value
	}
}

func TestGet(t *testing.T) {
	rootNode := New()

	for _, test := range createAddTestSeq() {
		node, _ := rootNode.Add([]int32(test.Key))
		node.Value = test.Value
	}

	for _, test := range createGetTestSeq() {
		node, exist := rootNode.Get([]int32(test.Key))

		if exist != test.Exist {
			var negation string
			if exist {
				negation = "should"
			} else {
				negation = "shouldn't"
			}

			template := "The node corresposing to \"%s\" %s exist."
			t.Errorf(template, test.Key, negation)
			return
		}

		if exist && node.Value != test.Value {
			template := "The node corresposing to \"%s\" should have %d, but has %d."
			t.Errorf(template, test.Key, node.Value, test.Value)
			return
		}
	}
}

func TestFindMax(t *testing.T) {
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

	node := internalNode.FindMax(func(x, y int) bool {
		return x < y
	})

	if value := node.Value; value != 5 {
		template := "The child node should have %d as its \"Value\", but have \"%d\"."
		t.Errorf(template, 5, value)
		return
	}

	if char := node.Char(); char != 'd' {
		template := "The child node should have %s as its \"Char\", but have \"%s\"."
		t.Errorf(template, "d", string(char))
		return
	}
}

func TestGetChildByOffset(t *testing.T) {
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

	node, exist := internalNode.GetChildByOffset(3)
	if !exist {
		message := "Cannot get a child node with not existing index."
		t.Errorf(message)
		return
	}

	if value := node.Value; value != 5 {
		template := "The child node should have %d as its \"Value\", but have \"%d\"."
		t.Errorf(template, 5, value)
		return
	}

	if char := node.Char(); char != 'd' {
		template := "The child node should have %s as its \"Char\", but have \"%s\"."
		t.Errorf(template, "d", string(char))
		return
	}
}
