package ngram

import (
	"bytes"
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

type ScanTokensTest struct {
	Data    []byte
	AtEOF   bool
	Advance int
	Token   []byte
}

func TestScanTokensSucceed(t *testing.T) {
	testSeq := []*ScanTokensTest{
		&ScanTokensTest{
			Data: []byte("aaa bbb ccc"), AtEOF: false,
			Advance: 3, Token: []byte("aaa"),
		},

		&ScanTokensTest{
			Data: []byte("   aaa bbb ccc"), AtEOF: false,
			Advance: 6, Token: []byte("aaa"),
		},

		&ScanTokensTest{
			Data: []byte("aaa    bbb ccc"), AtEOF: false,
			Advance: 3, Token: []byte("aaa"),
		},

		&ScanTokensTest{
			Data: []byte("aa(a bbb ccc"), AtEOF: false,
			Advance: 2, Token: []byte("aa"),
		},

		&ScanTokensTest{
			Data: []byte("[aaa bbb ccc"), AtEOF: false,
			Advance: 1, Token: []byte("["),
		},

		&ScanTokensTest{
			Data: []byte("  :aaa bbb ccc"), AtEOF: false,
			Advance: 3, Token: []byte(":"),
		},

		&ScanTokensTest{
			Data: []byte("aaa"), AtEOF: false,
			Advance: 0, Token: nil,
		},

		&ScanTokensTest{
			Data: []byte("  aaa"), AtEOF: false,
			Advance: 0, Token: nil,
		},

		&ScanTokensTest{
			Data: []byte("aaa bbb ccc"), AtEOF: true,
			Advance: 3, Token: []byte("aaa"),
		},

		&ScanTokensTest{
			Data: []byte("   aaa bbb ccc"), AtEOF: true,
			Advance: 6, Token: []byte("aaa"),
		},

		&ScanTokensTest{
			Data: []byte("aaa    bbb ccc"), AtEOF: true,
			Advance: 3, Token: []byte("aaa"),
		},

		&ScanTokensTest{
			Data: []byte("aa(a bbb ccc"), AtEOF: true,
			Advance: 2, Token: []byte("aa"),
		},

		&ScanTokensTest{
			Data: []byte("[aaa bbb ccc"), AtEOF: true,
			Advance: 1, Token: []byte("["),
		},

		&ScanTokensTest{
			Data: []byte("  :aaa bbb ccc"), AtEOF: true,
			Advance: 3, Token: []byte(":"),
		},

		&ScanTokensTest{
			Data: []byte("aaa"), AtEOF: true,
			Advance: 3, Token: []byte("aaa"),
		},

		&ScanTokensTest{
			Data: []byte("  aaa"), AtEOF: true,
			Advance: 5, Token: []byte("aaa"),
		},
	}

	for _, test := range testSeq {
		advance, token, err := ScanTokens(test.Data, test.AtEOF)

		if err != nil {
			template := "Any error shouldn't be caused: %s"
			t.Errorf(template, err.Error())
			return
		}

		if advance != test.Advance {
			template := "%d byte(s) should be read, but %d byte(s) have read."
			t.Errorf(template, test.Advance, advance)
			return
		}

		if bytes.Compare(token, test.Token) != 0 {
			template := "\"%s\" should be returned, but \"%s\" is returned."
			t.Errorf(template, test.Token, token)
			return
		}
	}
}
