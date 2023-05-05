package adapter

import (
	"fmt"

	"github.com/andre-ols/chatservice/internal/domain/port"
	"github.com/pkoukk/tiktoken-go"
)

type Tokenizer struct{}

// CountTokens implements entity.TokenCounter
func (*Tokenizer) CountTokens(model string, text string) (int, error) {
	tkm, err := tiktoken.EncodingForModel(model)
	if err != nil {
		err = fmt.Errorf("getEncoding: %v", err)
		return 0, err
	}

	// encode
	token := tkm.Encode(text, nil, nil)

	// num_tokens
	num_tokens := len(token)

	return num_tokens, nil
}

func NewTokenCounter() port.Tokenizer {
	return &Tokenizer{}
}
