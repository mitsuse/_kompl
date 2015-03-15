package tokenizer

import (
	"regexp"
	"strings"
)

const (
	_EN_SYMBOL_PATTERN = `([~!@#\$%\^&\*\(\)\-_\+=\[\]\{\}\|\\;:"',<>\/\?])`
	_EN_DOT_PATTERN    = `([^A-Z\.])\.`
	_EN_START_PATTERN  = `^\s+`
	_EN_END_PATTERN    = `\s+$`
)

type EnglishTokenizer struct {
	symbolRegexp *regexp.Regexp
	dotRegexp    *regexp.Regexp
	startRegexp  *regexp.Regexp
	endRegexp    *regexp.Regexp
}

func NewEnglishTokenizer() *EnglishTokenizer {
	t := &EnglishTokenizer{
		symbolRegexp: regexp.MustCompile(_EN_SYMBOL_PATTERN),
		dotRegexp:    regexp.MustCompile(_EN_DOT_PATTERN),
		startRegexp:  regexp.MustCompile(_EN_START_PATTERN),
		endRegexp:    regexp.MustCompile(_EN_END_PATTERN),
	}

	return t
}

func (t *EnglishTokenizer) Tokenize(s string) []string {
	cs := t.symbolRegexp.ReplaceAllString(s, " $1 ")
	cs = t.dotRegexp.ReplaceAllString(cs, "$1 .")
	cs = t.startRegexp.ReplaceAllString(cs, "")
	cs = t.endRegexp.ReplaceAllString(cs, "")

	return strings.Split(cs, " ")
}
