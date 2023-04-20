package entity

import (
	"errors"

	"github.com/google/uuid"
)

type Status string

const (
	Active Status = "active"
	Ended  Status = "ended"
)

type ChatConfig struct {
	Model            *Model
	Temperature      float32  // 0.0 - 1.0
	TopP             float32  // 0.0 - 1.0
	N                int      // Number of tokens to generate
	Stop             []string // Stop tokens
	MaxTokens        int      // Max number of tokens to generate
	PresencePenalty  float32  // -2.0 to 2.0 - Number beetwen -2.0 and 2.0
	FrequencyPenalty float32  // -2.0 to 2.0 - Number beetwen -2.0 and 2.0
}

type Chat struct {
	ID                   string
	UserID               string
	InitialSystemMessage *Message
	Messages             []*Message
	ErasedMessages       []*Message
	Status               Status
	TokenUsage           int
	Config               *ChatConfig
}

func NewChat(userID string, InitialSystemMessage *Message, config *ChatConfig) (*Chat, error) {
	chat := &Chat{
		ID:         uuid.New().String(),
		UserID:     userID,
		Status:     Active,
		Config:     config,
		TokenUsage: 0,
	}

	if err := chat.AddMessage(InitialSystemMessage); err != nil {
		return nil, err
	}

	if err := chat.Validate(); err != nil {
		return nil, err
	}

	return chat, nil
}

func (c *Chat) Validate() error {
	if c.UserID == "" {
		return errors.New("user_id is empty")
	}

	if c.Config == nil {
		return errors.New("config is nil")
	}

	if c.Status != Active && c.Status != Ended {
		return errors.New("invalid status")
	}

	if c.Config.Model == nil {
		return errors.New("model is nil")
	}

	if c.Config.Temperature < 0.0 || c.Config.Temperature > 1.0 {
		return errors.New("invalid temperature")
	}

	if c.Config.TopP < 0.0 || c.Config.TopP > 1.0 {
		return errors.New("invalid top_p")
	}

	if c.Config.N < 1 {
		return errors.New("invalid n")
	}

	if c.Config.MaxTokens < 1 {
		return errors.New("invalid max_tokens")
	}

	if c.Config.PresencePenalty < -2.0 || c.Config.PresencePenalty > 2.0 {
		return errors.New("invalid presence_penalty")
	}

	if c.Config.FrequencyPenalty < -2.0 || c.Config.FrequencyPenalty > 2.0 {
		return errors.New("invalid frequency_penalty")
	}

	return nil
}

func (c *Chat) AddMessage(msg *Message) error {
	if c.Status == Ended {
		return errors.New("chat is ended. No more messages allowed")
	}

	for {
		if c.Config.Model.GetMaxTokens() >= msg.GetQuantityTokens()+c.TokenUsage {
			c.Messages = append(c.Messages, msg)
			c.RefreshTokenUsage()
			break
		}
		c.ErasedMessages = append(c.ErasedMessages, c.Messages[0])
		c.Messages = c.Messages[1:]
		c.RefreshTokenUsage()
	}

	return nil
}

func (c *Chat) GetMessages() []*Message {
	return c.Messages
}

func (c *Chat) CountMessages() int {
	return len(c.Messages)
}

func (c *Chat) End() {
	c.Status = Ended
}

func (c *Chat) RefreshTokenUsage() {
	c.TokenUsage = 0
	for i := range c.Messages {
		c.TokenUsage += c.Messages[i].GetQuantityTokens()
	}
}
