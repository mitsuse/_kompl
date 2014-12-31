package ngram

import (
	"testing"
)

func TestIsSymbol(t *testing.T) {
	charSeq := []rune{
		'~',
		'!',
		'@',
		'#',
		'$',
		'%',
		'^',
		'&',
		'*',
		'(',
		')',
		'-',
		'_',
		'+',
		'=',
		'[',
		']',
		'{',
		'}',
		'|',
		'\\',
		';',
		':',
		'"',
		'\'',
		',',
		'.',
		'<',
		'>',
		'/',
		'?',
	}

	for _, char := range charSeq {
		if !IsSymbol(char) {
			t.Errorf("\"%s\" should be a symbol.", char)
		}
	}
}

func TestIsNotSymbol(t *testing.T) {
	charSeq := []rune{
		'a',
		'b',
		'c',
		'x',
		'y',
		'z',
		'0',
		'1',
		'2',
		'7',
		'8',
		'9',
		'あ',
		'い',
		'う',
		'ば',
		'び',
		'ぶ',
		'わ',
		'を',
		'ん',
		' ',
		'`',
	}

	for _, char := range charSeq {
		if IsSymbol(char) {
			t.Errorf("\"%s\" should not be a symbol.", char)
		}
	}
}
