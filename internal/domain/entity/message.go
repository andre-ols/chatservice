package entity

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/pkoukk/tiktoken-go"
)

type Role string

const (
	User      Role = "user"
	System    Role = "system"
	Assistant Role = "assistant"
)

type TokenCounter interface {
	CountTokens(model string, text string) (int, error)
}

type Message struct {
	ID        string       `faker:"uuid_digit"`
	Role      Role         `faker:"oneof: user,system,assistant"`
	Content   string       `faker:"sentence"`
	Tokens    int          `faker:"boundary_start=1, boundary_end=1024"`
	Model     *Model       `faker:"-"`
	Tokenizer TokenCounter `faker:"-"`
	CreatedAt time.Time    `faker:"time_now"`
}

func NewMessage(role Role, content string, model *Model, tokenizer TokenCounter) (*Message, error) {

	if model == nil {
		return nil, errors.New("model is nil")
	}

	totalTokens, err := tokenizer.CountTokens(model.Name, content)

	if err != nil {
		return nil, err
	}

	msg := &Message{
		ID:        uuid.New().String(),
		Role:      role,
		Content:   content,
		Tokens:    totalTokens,
		Model:     model,
		CreatedAt: time.Now(),
	}

	if err := msg.Validate(); err != nil {
		return nil, err
	}

	return msg, nil

}

func (m *Message) Validate() error {

	if m.ID == "" {
		return errors.New("id is empty")
	}

	if m.Role != User && m.Role != System && m.Role != Assistant {
		return errors.New("invalid role")
	}

	if m.Content == "" {
		return errors.New("content is empty")
	}

	if m.Model == nil {
		return errors.New("model is nil")
	}

	if m.Tokens <= 0 {
		return errors.New("tokens is less than or equal to 0")
	}

	if m.CreatedAt.IsZero() {
		return errors.New("created_at is empty")
	}

	return nil
}

func CountTokens(model string, text string) (int, error) {

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
