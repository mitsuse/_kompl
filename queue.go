package compl

import (
	"sort"
)

type Queue struct {
	seq CandidateSeq
}

func (q *Queue) Push(candidate *Candidate) {
	q.seq = append(q.seq, candidate)
	sort.Sort(q.seq)
}

func (q *Queue) Pop() (*Queue, bool) {
	if len(q.seq) == 0 {
		return nil, false
	}

	offset := len(q.seq) - 1

	candidate := q.seq[offset]
	q.seq = q.seq[offset]

	return candidate, true
}
