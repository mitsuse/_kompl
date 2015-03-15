/*
Package "tokenizer" provides the interface for tokenizers and several implementations.
*/
package tokenizer

type Tokenizer interface {
	Tokenize(s string) []string
}
