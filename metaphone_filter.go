package tokenizer

import (
	"github.com/dotcypress/phonetics"
)

type MetaphoneFilter struct {
}

func NewMetaphoneFilter() *MetaphoneFilter {
	return &MetaphoneFilter{}
}

func (m *MetaphoneFilter) Process(input []string) []string {
	output := make([]string, len(input))
	for idx, token := range input {
		output[idx] = phonetics.EncodeMetaphone(token)
	}
	return output
}
