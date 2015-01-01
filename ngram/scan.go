package ngram

import (
	"unicode"
	"unicode/utf8"
)

func ScanTokens(data []byte, atEOF bool) (advance int, token []byte, err error) {
	start := 0

	// Skip leading spaces.
	for width := 0; start < len(data); start += width {
		var r rune
		r, width = utf8.DecodeRune(data[start:])
		if !unicode.IsSpace(r) {
			break
		}
	}

	// Check wheter the first charactor is a symbol or not.
	if r, width := utf8.DecodeRune(data[start:]); IsSymbol(r) {
		return start + width, data[start : start+width], nil
	}

	// Scan until space, marking end of word.
	for width, i := 0, start; i < len(data); i += width {
		var r rune
		r, width = utf8.DecodeRune(data[i:])
		if unicode.IsSpace(r) || IsSymbol(r) {
			return i, data[start:i], nil
		}
	}

	// If we're at EOF, we have a final, non-empty, non-terminated word. Return it.
	if atEOF && len(data) > start {
		return len(data), data[start:], nil
	}

	// Request more data.
	return start, nil, nil
}

func IsSymbol(char rune) bool {
	return symbolRegexp.MatchString(string(char))
}
