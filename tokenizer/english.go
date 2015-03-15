package tokenizer

import (
	"regexp"
	"strings"
)

const (
	_EN_SYMBOL_PATTERN = `([~!@#\$%\^&\*\(\)\-_\+=\[\]\{\}\|\\;:"',<>\/\?])`
	_EN_DOTSEQ_PATTERN = `(\.{2,})`
	_EN_DOT_PATTERN    = `([^A-Z\.])\.`
	_EN_START_PATTERN  = `^\s+`
	_EN_END_PATTERN    = `\s+$`
	_EN_SPACE_PATTERN  = `\s+`
)

type EnglishTokenizer struct {
	symbolRegexp *regexp.Regexp
	dotSeqRegexp *regexp.Regexp
	dotRegexp    *regexp.Regexp
	startRegexp  *regexp.Regexp
	endRegexp    *regexp.Regexp
	spaceRegexp  *regexp.Regexp
}

func NewEnglishTokenizer() *EnglishTokenizer {
	t := &EnglishTokenizer{
		symbolRegexp: regexp.MustCompile(_EN_SYMBOL_PATTERN),
		dotSeqRegexp: regexp.MustCompile(_EN_DOTSEQ_PATTERN),
		dotRegexp:    regexp.MustCompile(_EN_DOT_PATTERN),
		startRegexp:  regexp.MustCompile(_EN_START_PATTERN),
		endRegexp:    regexp.MustCompile(_EN_END_PATTERN),
		spaceRegexp:  regexp.MustCompile(_EN_SPACE_PATTERN),
	}

	return t
}

/*
"(*EnglishTokenizer).Tokenize" tokenize the given string "s" stupidly,
and returns the sequece of tokens typed as "[]string".
*/
func (t *EnglishTokenizer) Tokenize(s string) []string {
	cs := t.symbolRegexp.ReplaceAllString(s, " $1 ")
	cs = t.dotSeqRegexp.ReplaceAllString(cs, " $1 ")
	cs = t.dotRegexp.ReplaceAllString(cs, "$1 .")
	cs = t.startRegexp.ReplaceAllString(cs, "")
	cs = t.endRegexp.ReplaceAllString(cs, "")
	cs = t.spaceRegexp.ReplaceAllString(cs, " ")

	return strings.Split(cs, " ")
}
