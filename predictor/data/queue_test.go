package data

import (
	"testing"
)

func TestQueueSeqLen(t *testing.T) {
	queue := NewQueue()

	aCandidate := &Candidate{score: 10}
	bCandidate := &Candidate{score: 5}
	cCandidate := &Candidate{score: 20}
	dCandidate := &Candidate{score: 0}

	queue.Push(aCandidate)
	queue.Push(bCandidate)
	queue.Push(cCandidate)
	queue.Push(dCandidate)

	if length := queue.Len(); length != 4 {
		template := "\"(*Queue).Len\" should return %d, but returns %d."
		t.Errorf(template, 4, length)
		return
	}
}

func TestQueueSeqPop(t *testing.T) {
	queue := NewQueue()

	scoreSeq := []int{20, 10, 5, 0}

	aCandidate := &Candidate{score: 10}
	bCandidate := &Candidate{score: 5}
	cCandidate := &Candidate{score: 20}
	dCandidate := &Candidate{score: 0}

	queue.Push(aCandidate)
	queue.Push(bCandidate)
	queue.Push(cCandidate)
	queue.Push(dCandidate)

	for _, score := range scoreSeq {
		candidate, exist := queue.Pop()
		if !exist {
			message := "The number of candidates contained in queue is less than expected."
			t.Errorf(message)
		}

		if candidate.Score() != score {
			template := "The candidate's score should be %d, but is %d."
			t.Errorf(template, score, candidate.Score)
			return
		}
	}

	if queue.Len() > 0 {
		message := "The number of candidates contained in queue is more than expected."
		t.Errorf(message)
		return
	}
}

func TestQueueSeqPopEmpty(t *testing.T) {
	queue := NewQueue()

	_, exist := queue.Pop()
	if exist {
		message := "The empty queue should return \"false\", but returns \"true\"."
		t.Errorf(message)
		return
	}
}
