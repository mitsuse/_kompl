package sentencizer

import (
	"testing"
)

func TestEnglishSentencizerIsSentencizer(t *testing.T) {
	var _ Sentencizer = NewEnglishSentencizer()
}

func TestEnglishSentencizerSentencize(t *testing.T) {
	testSeq := []*SentencizeTest{
		&SentencizeTest{
			TokenSeq: []string{"AA", "bb", ".", "cC", "dd", "."},
			SentenceSeq: [][]string{
				[]string{"AA", "bb", "."},
				[]string{"cC", "dd", "."},
			},
		},

		&SentencizeTest{
			TokenSeq: []string{"AA", "bb", "...", "cC", "dd", "."},
			SentenceSeq: [][]string{
				[]string{"AA", "bb", "...", "cC", "dd", "."},
			},
		},

		&SentencizeTest{
			TokenSeq: []string{"AA", "bb", ".", "cC", "dd", "..."},
			SentenceSeq: [][]string{
				[]string{"AA", "bb", "."},
				[]string{"cC", "dd", "..."},
			},
		},

		&SentencizeTest{
			TokenSeq: []string{".", "cC", "dd", "."},
			SentenceSeq: [][]string{
				[]string{"."},
				[]string{"cC", "dd", "."},
			},
		},

		&SentencizeTest{
			TokenSeq: []string{"AA", "bb", "?", "cC", "dd", "!"},
			SentenceSeq: [][]string{
				[]string{"AA", "bb", "?"},
				[]string{"cC", "dd", "!"},
			},
		},

		&SentencizeTest{
			TokenSeq: []string{"AA", "bb", "???", "cC", "dd", "!!!!"},
			SentenceSeq: [][]string{
				[]string{"AA", "bb", "???"},
				[]string{"cC", "dd", "!!!!"},
			},
		},

		&SentencizeTest{
			TokenSeq: []string{"AA", "bb", "?!!??", "cC", "dd", "!!!??!?"},
			SentenceSeq: [][]string{
				[]string{"AA", "bb", "?!!??"},
				[]string{"cC", "dd", "!!!??!?"},
			},
		},
	}

	sentencizer := NewEnglishSentencizer()

	for _, test := range testSeq {
		sentenceSeq := sentencizer.Sentencize(test.TokenSeq)

		if len(sentenceSeq) != len(test.SentenceSeq) {
			template := "The tokens sould be split into %d sentence(s), but into %d."
			t.Errorf(template, len(test.SentenceSeq), len(sentenceSeq))
			return
		}

		for i := 0; i < len(sentenceSeq); i++ {
			expectedTokenSeq := test.SentenceSeq[i]
			resultTokenSeq := sentenceSeq[i]

			if len(resultTokenSeq) != len(expectedTokenSeq) {
				template := "Expected: %v\nResult: %v"
				t.Errorf(template, expectedTokenSeq, resultTokenSeq)
				return
			}
		}
	}
}
