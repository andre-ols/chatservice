package port

type Tokenizer interface {
	CountTokens(model string, text string) (int, error)
}
