package entity

import (
	"errors"
	"time"

	"github.com/andre-ols/chatservice/internal/domain/port"
	"github.com/google/uuid"
)

type Role string

const (
	User      Role = "user"
	System    Role = "system"
	Assistant Role = "assistant"
)

type Message struct {
	ID        string         `faker:"uuid_digit"`
	Role      Role           `faker:"oneof: user,system,assistant"`
	Content   string         `faker:"sentence"`
	Tokens    int            `faker:"boundary_start=1, boundary_end=1024"`
	Model     *Model         `faker:"-"`
	Tokenizer port.Tokenizer `faker:"-"`
	CreatedAt time.Time      `faker:"time_now"`
}

func NewMessage(role Role, content string, model *Model, tokenizer port.Tokenizer) (*Message, error) {

	if model == nil {
		return nil, errors.New("model is nil")
	}

	if err := model.Validate(); err != nil {
		return nil, err
	}

	msg := &Message{
		ID:        uuid.New().String(),
		Role:      role,
		Content:   content,
		Model:     model,
		Tokenizer: tokenizer,
		CreatedAt: time.Now(),
	}

	err := msg.CountTokens()

	if err != nil {
		return nil, err
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

	if m.Tokens < 0 {
		return errors.New("tokens is less than 0")
	}

	if m.CreatedAt.IsZero() {
		return errors.New("created_at is empty")
	}

	if m.Model == nil {
		return errors.New("model is nil")
	}

	if err := m.Model.Validate(); err != nil {
		return err
	}

	if m.Tokenizer == nil {
		return errors.New("tokenizer is nil")
	}

	return nil
}

func (m *Message) CountTokens() error {

	totalTokens, err := m.Tokenizer.CountTokens(m.Model.Name, m.Content)

	if err != nil {
		return err
	}

	m.Tokens = totalTokens

	return nil
}
