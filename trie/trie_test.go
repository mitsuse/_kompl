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
