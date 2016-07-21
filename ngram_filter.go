package tokenizer

import (
	"unicode/utf8"
)

type NGramFilter struct {
	Min int
	Max int
}

func NewNGramFilter(min, max int) *NGramFilter {
	return &NGramFilter{min, max}
}

func (n *NGramFilter) Process(input []string) []string {
	output := []string{}
	for nlen := n.Min; nlen <= n.Max; nlen++ {
		for _, token := range input {
			output = append(output, ngrams(token, nlen)...)
		}
	}
	return output
}

func ngrams(text string, nlen int) []string {
	if nlen > len(text) {
		return []string{text}
	}
	input := text

	runes := make([]int, len(input)+1)
	ridx := 0 // rune index
	bidx := 0 // byte index
	for len(input) > 0 {
		_, width := utf8.DecodeRuneInString(input)
		input = input[width:]
		runes[ridx] = bidx
		bidx += width
		ridx++
	}
	runes[ridx] = len(text)

	tokens := make([]string, ridx+1-nlen)
	for i := 0; i <= ridx-nlen; i++ {
		end := i + nlen
		if end >= len(runes) {
			break
		}
		tokens[i] = text[runes[i]:runes[end]]
	}
	return tokens
}
