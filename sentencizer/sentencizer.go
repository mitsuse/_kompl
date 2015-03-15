/*
Package "sentencizer" provides the interface for sentencizers and several implementations.
*/
package sentencizer

type Sentencizer interface {
	Sentencize(tokenSeq []string) [][]string
}
