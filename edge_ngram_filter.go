package tokenizer

import ()

type EdgeNGramFilter struct {
	Min         int
	Max         int
	Leading     bool
	Trailing    bool
	IncludeFull bool
}

func NewEdgeNGramFilter(min, max int) *EdgeNGramFilter {
	return &EdgeNGramFilter{min, max, true, false, true}
}

func (e *EdgeNGramFilter) Process(input []string) []string {
	output := []string{}

	if e.IncludeFull {
		output = append(output, input...)
	}

	for _, token := range input {
		runes := []rune(token)
		for nlen := e.Min; nlen <= e.Max && nlen < len(runes); nlen++ {
			if nlen > len(token) {
				continue
			}

			if e.Leading {
				output = append(output, string(runes[0:nlen]))
			}
			if e.Trailing {
				output = append(output, string(runes[nlen:]))
			}
		}
	}
	return output
}
