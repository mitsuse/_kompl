package trie

type TrieSeq []*Trie

func (s TrieSeq) Len() int {
	return len(s)
}

func (s TrieSeq) Less(i, j int) bool {
	return s[i].char < s[j].char
}

func (s TrieSeq) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
