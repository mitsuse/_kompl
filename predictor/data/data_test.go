package data

import (
	"testing"

	"github.com/mitsuse/kompl/trie"
)

func TestNewCandidate(t *testing.T) {
	rootNode := trie.New()

	char := "l"
	prefix := "komp"
	word := prefix + char

	node, _ := rootNode.Add([]int32(word))
	parentNode, _ := rootNode.Get([]int32(prefix))
	score := 10

	candidate := NewCandidate(word, node, parentNode, score)

	if w := candidate.Word(); word != w {
		template := "\"(*Candidate).Word\" should return \"%s\", but returns \"%s\"."
		t.Errorf(template, word, w)
		return
	}

	if n := candidate.Node(); node != n {
		template := "\"(*Candidate).Node\" should return %v@%p, but returns %v@%p."
		t.Errorf(template, node, node, n, n)
		return
	}

	if p := candidate.Parent(); parentNode != p {
		template := "\"(*Candidate).Parent\" should return %v@%p, but returns %v@%p."
		t.Errorf(template, parentNode, p)
		return
	}

	if s := candidate.Score(); score != s {
		template := "\"(*Candidate).Score\" should return %d, but returns %d."
		t.Errorf(template, score, s)
		return
	}
}
