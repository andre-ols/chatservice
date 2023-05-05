package entity

import "errors"

type ChatConfig struct {
	Temperature      float32  `faker:"boundary_start=0.0, boundary_end=1.0"`  // 0.0 - 1.0
	TopP             float32  `faker:"boundary_start=0.0, boundary_end=1.0"`  // 0.0 - 1.0
	N                int      `faker:"boundary_start=1, boundary_end=100"`    // Number of tokens to generate
	Stop             []string `faker:"len=3"`                                 // Stop sequence
	MaxTokens        int      `faker:"boundary_start=1, boundary_end=100"`    // Max number of tokens to generate
	PresencePenalty  float32  `faker:"boundary_start=-2.0, boundary_end=2.0"` // -2.0 to 2.0 - Number beetwen -2.0 and 2.0
	FrequencyPenalty float32  `faker:"boundary_start=-2.0, boundary_end=2.0"` // -2.0 to 2.0 - Number beetwen -2.0 and 2.0
	Model            *Model   `faker:"-"`
}

type ChatConfigInput struct {
	Temperature      float32
	TopP             float32
	N                int
	Stop             []string
	MaxTokens        int
	PresencePenalty  float32
	FrequencyPenalty float32
	Model            *Model
}

func NewChatConfig(input ChatConfigInput) (*ChatConfig, error) {

	if input.Model == nil {
		return nil, errors.New("model is nil")
	}

	if err := input.Model.Validate(); err != nil {
		return nil, err
	}

	chat := &ChatConfig{
		Model:            input.Model,
		Temperature:      input.Temperature,
		TopP:             input.TopP,
		N:                input.N,
		Stop:             input.Stop,
		MaxTokens:        input.MaxTokens,
		PresencePenalty:  input.PresencePenalty,
		FrequencyPenalty: input.FrequencyPenalty,
	}

	if err := chat.Validate(); err != nil {
		return nil, err
	}

	return chat, nil
}

func (c *ChatConfig) Validate() error {

	if c.Temperature < 0.0 || c.Temperature > 1.0 {
		return errors.New("invalid temperature")
	}

	if c.TopP < 0.0 || c.TopP > 1.0 {
		return errors.New("invalid topP")
	}

	if c.N < 1 {
		return errors.New("invalid n")
	}

	if c.MaxTokens < 1 {
		return errors.New("invalid maxTokens")
	}

	if c.PresencePenalty < -2.0 || c.PresencePenalty > 2.0 {
		return errors.New("invalid presencePenalty")
	}

	if c.FrequencyPenalty < -2.0 || c.FrequencyPenalty > 2.0 {
		return errors.New("invalid frequencyPenalty")
	}

	if c.Stop == nil {
		return errors.New("invalid stop")
	}

	if c.Model == nil {
		return errors.New("model is nil")
	}

	if err := c.Model.Validate(); err != nil {
		return err
	}

	return nil
}
