package ngram

import (
	"regexp"
)

const (
	_SYMBOL_PATTERN = `^[~!@#\$%\^&\*\(\)\-_\+=\[\]\{\}\|\\;:"',\.<>\/\?]$`
)

var symbolRegexp *regexp.Regexp

func init() {
	symbolRegexp = regexp.MustCompile(_SYMBOL_PATTERN)
}
