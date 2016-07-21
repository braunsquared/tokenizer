package tokenizer

import (
	"strings"
)

type SplitWordFilter struct {
	Seperator string
}

func NewSplitWordFilter(sep *string) *SplitWordFilter {
	if sep == nil {
		return &SplitWordFilter{" "}
	}

	return &SplitWordFilter{*sep}
}

func (s *SplitWordFilter) Process(input []string) []string {
	output := []string{}
	for _, token := range input {
		tokens := strings.Split(token, s.Seperator)
		output = append(output, tokens...)
	}
	return output
}
