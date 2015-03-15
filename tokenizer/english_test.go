package tokenizer

import (
	"testing"
)

func TestEnglishTokenizerIsTokenize(t *testing.T) {
	var _ Tokenizer = NewEnglishTokenizer()
}

func TestEnglishTokenizerTokenize(t *testing.T) {
	testSeq := []*TokenizeTest{
		&TokenizeTest{
			CharSeq: "If you see the \"hello, world\" message, \"Go\" is working.",
			TokenSeq: []string{
				"If",
				"you",
				"see",
				"the",
				"\"",
				"hello",
				",",
				"world",
				"\"",
				"message",
				",",
				"\"",
				"Go",
				"\"",
				"is",
				"working",
				".",
			},
		},

		&TokenizeTest{
			CharSeq: "   If you see the \"hello,  world\"   message, \"Go\" is working.  ",
			TokenSeq: []string{
				"If",
				"you",
				"see",
				"the",
				"\"",
				"hello",
				",",
				"world",
				"\"",
				"message",
				",",
				"\"",
				"Go",
				"\"",
				"is",
				"working",
				".",
			},
		},

		&TokenizeTest{
			CharSeq: "If you see the \"hello, world\" message, \"Go\" is working...!",
			TokenSeq: []string{
				"If",
				"you",
				"see",
				"the",
				"\"",
				"hello",
				",",
				"world",
				"\"",
				"message",
				",",
				"\"",
				"Go",
				"\"",
				"is",
				"working",
				"...",
				"!",
			},
		},
	}

	tokenizer := NewEnglishTokenizer()

	for _, test := range testSeq {
		tokenSeq := tokenizer.Tokenize(test.CharSeq)

		if len(tokenSeq) != len(test.TokenSeq) {
			template := "The input should be split into %d token(s), but into %d token(s)."
			t.Errorf(template, len(test.TokenSeq), len(tokenSeq))
			return
		}

		for i := 0; i < len(tokenSeq); i++ {
			if tokenSeq[i] != test.TokenSeq[i] {
				template := "Expected: %v\nResult: %v"
				t.Errorf(template, test.TokenSeq, tokenSeq)
				return
			}
		}
	}
}
