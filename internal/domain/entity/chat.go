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

type Chat struct {
	ID                   string      `faker:"uuid_digit"`
	UserID               string      `faker:"uuid_digit"`
	InitialSystemMessage *Message    `faker:"-"`
	Messages             []*Message  `faker:"-"`
	ErasedMessages       []*Message  `faker:"-"`
	Status               Status      `faker:"oneof: active,ended"`
	TokenUsage           int         `faker:"boundary_start=0, boundary_end=100"`
	Config               *ChatConfig `faker:"-"`
}

func NewChat(userID string, InitialSystemMessage *Message, config *ChatConfig) (*Chat, error) {

	if InitialSystemMessage == nil {
		return nil, errors.New("initial system message is nil")
	}

	if err := InitialSystemMessage.Validate(); err != nil {
		return nil, err
	}

	if config == nil {
		return nil, errors.New("config is nil")
	}

	if err := config.Validate(); err != nil {
		return nil, err
	}

	chat := &Chat{
		ID:                   uuid.New().String(),
		UserID:               userID,
		Status:               Active,
		InitialSystemMessage: InitialSystemMessage,
		TokenUsage:           0,
		Config:               config,
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

	if c.ID == "" {
		return errors.New("id is empty")
	}

	if c.UserID == "" {
		return errors.New("user id is empty")
	}

	if c.Status != Active && c.Status != Ended {
		return errors.New("invalid status")
	}

	if c.TokenUsage < 0 {
		return errors.New("invalid token usage")
	}

	if c.Messages == nil {
		return errors.New("messages is nil")
	}

	// Validate InitialSystemMessage
	if c.InitialSystemMessage == nil {
		return errors.New("initial system message is nil")
	}

	if err := c.InitialSystemMessage.Validate(); err != nil {
		return err
	}

	// Validate Config
	if c.Config == nil {
		return errors.New("config is nil")
	}

	if err := c.Config.Validate(); err != nil {
		return err
	}

	return nil
}

func (c *Chat) AddMessage(msg *Message) error {
	if c.Status == Ended {
		return errors.New("chat is ended. No more messages allowed")
	}

	for {
		if c.Config.Model.GetMaxTokens() >= msg.Tokens+c.TokenUsage {
			c.Messages = append(c.Messages, msg)
			c.RefreshTokenUsage()
			break
		}

		if len(c.Messages) == 0 {
			return errors.New("message is too long")
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
		c.TokenUsage += c.Messages[i].Tokens
	}
}
