package compl

import (
	"github.com/mitsuse/compl/trie"
)

type Candidate struct {
	word  string
	node  *trie.Trie
	score int
}

func NewCandidate(word string, node *trie.Trie, score int) *Candidate {
	c := &Candidate{
		word:  word,
		node:  node,
		score: score,
	}

	return c
}

func (c *Candidate) Word() string {
	return c.word
}

func (c *Candidate) Node() *trie.Trie {
	return c.node
}

func (c *Candidate) Score() int {
	return c.score
}

type CandidateSeq []*Candidate

func (s CandidateSeq) Len() int {
	return len(s)
}

func (s CandidateSeq) Less(i, j int) bool {
	return s[i].score < s[j].score
}

func (s CandidateSeq) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
