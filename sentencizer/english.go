package sentencizer

import (
	"regexp"
)

const (
	_EN_DOT_EOS_PATTERN    = `^[\.]$`
	_EN_SYMBOL_EOS_PATTERN = `^[!\?]+$`
)

type EnglishSentencizer struct {
	dotEosPattern   *regexp.Regexp
	symboEosPattern *regexp.Regexp
}

func NewEnglishSentencizer() *EnglishSentencizer {
	s := &EnglishSentencizer{
		dotEosPattern:   regexp.MustCompile(_EN_DOT_EOS_PATTERN),
		symboEosPattern: regexp.MustCompile(_EN_SYMBOL_EOS_PATTERN),
	}

	return s
}

func (s *EnglishSentencizer) Sentencize(tokenSeq []string) [][]string {
	sentenceSeq := make([][]string, 0)
	sentence := make([]string, 0)

	for _, token := range tokenSeq {
		sentence = append(sentence, token)

		if s.dotEosPattern.MatchString(token) || s.symboEosPattern.MatchString(token) {
			sentenceSeq = append(sentenceSeq, sentence)
			sentence = make([]string, 0)
			continue
		}
	}

	sentenceSeq = append(sentenceSeq, sentence)

	return sentenceSeq
}
