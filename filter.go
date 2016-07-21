package tokenizer

type Filter interface {
	Process(input []string) []string
}
