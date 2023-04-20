package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
	tiktoken_go "github.com/j178/tiktoken-go"
)

type Role string

const (
	User      Role = "user"
	System    Role = "system"
	Assistant Role = "assistant"
)

type Message struct {
	ID        string
	Role      Role
	Content   string
	Tokens    int
	Model     *Model
	CreatedAt time.Time
}

func NewMessage(role Role, content string, model *Model) (*Message, error) {

	totalTokens := tiktoken_go.CountTokens(model.GetName(), content)

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

	if m.Role != User && m.Role != System && m.Role != Assistant {
		return errors.New("invalid role")
	}

	if m.Content == "" {
		return errors.New("content is empty")
	}

	if m.Model == nil {
		return errors.New("model is nil")
	}

	if m.CreatedAt.IsZero() {
		return errors.New("created_at is empty")
	}

	return nil
}

func (m *Message) GetQuantityTokens() int {
	return m.Tokens
}
