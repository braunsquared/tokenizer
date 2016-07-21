package tokenizer

import (
// "fmt"
)

type Tokenizer struct {
	pipeline         []Filter
	RemoveDuplicates bool
	RemoveEmpties    bool
}

func NewTokenizer() *Tokenizer {
	return &Tokenizer{[]Filter{}, true, true}
}

func DefaultTokenizer() *Tokenizer {
	return &Tokenizer{
		pipeline:         []Filter{NewSplitWordFilter(nil)},
		RemoveDuplicates: true,
		RemoveEmpties:    true,
	}
}

func (t *Tokenizer) Append(filters ...Filter) *Tokenizer {
	t.pipeline = append(t.pipeline, filters...)
	return t
}

func (t *Tokenizer) Tokenize(source string) ([]string, error) {
	input := []string{source}

	if len(t.pipeline) == 0 {
		return input, nil
	}

	for _, filter := range t.pipeline {
		// fmt.Printf("Input %#v Len: %d\n", input, len(input))
		input = filter.Process(input)
		// fmt.Printf("Output %#v Len: %d\n", input, len(input))
		if t.RemoveDuplicates {
			input = removeDuplicates(input)
		}
		if t.RemoveEmpties {
			input = removeEmpties(input)
		}
	}

	return input, nil
}

func removeDuplicates(set []string) []string {
	// fmt.Printf("Dup Input %#v Len: %d\n", set, len(set))
	length := len(set) - 1
	for i := 0; i < length; i++ {
		for j := i + 1; j <= length; j++ {
			if set[i] == set[j] {
				set[j] = set[length]
				set = set[0:length]
				length--
				j--
			}
		}
	}
	// fmt.Printf("Dup Output %#v Len: %d\n", set, len(set))
	return set
}

func removeEmpties(set []string) []string {
	// fmt.Printf("Empty Input %#v Len: %d\n", set, len(set))
	output := []string{}
	for i, _ := range set {
		if len(set[i]) != 0 {
			output = append(output, set[i])
		}
	}
	// fmt.Printf("Empty Output %#v Len: %d\n", output, len(output))

	return output
}
