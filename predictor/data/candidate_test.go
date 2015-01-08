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

func TestCandidateSeqLen(t *testing.T) {
	seq := CandidateSeq{
		NewCandidate("", nil, nil, 4),
		NewCandidate("", nil, nil, 3),
		NewCandidate("", nil, nil, 1),
		NewCandidate("", nil, nil, 5),
		NewCandidate("", nil, nil, 0),
		NewCandidate("", nil, nil, 2),
	}

	if length := seq.Len(); length != len(seq) {
		template := "\"(CandidateSeq).Len\" should return %d, but returns %d."
		t.Errorf(template, len(seq), length)
		return
	}
}

func TestCandidateSeqLess(t *testing.T) {
	seq := CandidateSeq{
		NewCandidate("", nil, nil, 4),
		NewCandidate("", nil, nil, 3),
		NewCandidate("", nil, nil, 1),
		NewCandidate("", nil, nil, 5),
		NewCandidate("", nil, nil, 0),
		NewCandidate("", nil, nil, 2),
	}

	if seq.Less(0, 1) {
		message := "\"(CandidateSeq).Less\" should return false, but returns true."
		t.Errorf(message)
		return
	}
}

func TestCandidateSeqSwqp(t *testing.T) {
	seq := CandidateSeq{
		NewCandidate("", nil, nil, 4),
		NewCandidate("", nil, nil, 3),
		NewCandidate("", nil, nil, 1),
		NewCandidate("", nil, nil, 5),
		NewCandidate("", nil, nil, 0),
		NewCandidate("", nil, nil, 2),
	}

	seq.Swap(0, 1)

	if seq[0].score != 3 {
		template := "\"seq[%d].score\" should be %d but is %d."
		t.Errorf(template, 0, 3, seq[0].score)
		return
	}

	if seq[1].score != 4 {
		template := "\"seq[%d].score\" should be %d but is %d."
		t.Errorf(template, 1, 4, seq[1].score)
		return
	}
}
