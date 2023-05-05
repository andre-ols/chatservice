package entity_test

import (
	"testing"

	"github.com/andre-ols/chatservice/internal/domain/entity"
	fake "github.com/andre-ols/chatservice/testutils/fake/domain/entity"
	"github.com/andre-ols/chatservice/testutils/suite"
)

type ChatConfigTestSuite struct {
	suite.Suite
}

func TestChatConfigSuit(t *testing.T) {
	suite.RunUnitTest(t, new(ChatConfigTestSuite))
}

func (s *ChatConfigTestSuite) TestNewChatConfig() {
	type Sut struct {
		Sut  func() (*entity.ChatConfig, error)
		Data *entity.ChatConfig
	}

	makeSut := func() *Sut {
		data := fake.ChatConfig()

		sut := func() (*entity.ChatConfig, error) {
			input := entity.ChatConfigInput{
				Temperature:      data.Temperature,
				TopP:             data.TopP,
				N:                data.N,
				Stop:             data.Stop,
				MaxTokens:        data.MaxTokens,
				PresencePenalty:  data.PresencePenalty,
				FrequencyPenalty: data.FrequencyPenalty,
				Model:            data.Model,
			}
			return entity.NewChatConfig(input)
		}

		return &Sut{Sut: sut, Data: data}
	}

	s.Run("TestNewChatConfig Error", func() {
		s.Run("should return an error if temperature is invalid with a message", func() {
			sut := makeSut()
			sut.Data.Temperature = -1
			msg, err := sut.Sut()
			s.Error(err)
			s.Nil(msg)
			s.Equal("invalid temperature", err.Error())
		})

		s.Run("should return an error if topP is invalid with a message", func() {
			sut := makeSut()
			sut.Data.TopP = -1
			msg, err := sut.Sut()
			s.Error(err)
			s.Nil(msg)
			s.Equal("invalid topP", err.Error())
		})

		s.Run("should return an error if n is invalid with a message", func() {
			sut := makeSut()
			sut.Data.N = -1
			msg, err := sut.Sut()
			s.Error(err)
			s.Nil(msg)
			s.Equal("invalid n", err.Error())
		})

		s.Run("should return an error if stop is invalid with a message", func() {
			sut := makeSut()
			sut.Data.Stop = nil
			msg, err := sut.Sut()
			s.Error(err)
			s.Nil(msg)
			s.Equal("invalid stop", err.Error())
		})

		s.Run("should return an error if maxTokens is invalid with a message", func() {
			sut := makeSut()
			sut.Data.MaxTokens = -1
			msg, err := sut.Sut()
			s.Error(err)
			s.Nil(msg)
			s.Equal("invalid maxTokens", err.Error())
		})

		s.Run("should return an error if presencePenalty is invalid with a message", func() {
			sut := makeSut()
			sut.Data.PresencePenalty = -2.1
			msg, err := sut.Sut()
			s.Error(err)
			s.Nil(msg)
			s.Equal("invalid presencePenalty", err.Error())
		})

		s.Run("should return an error if frequencyPenalty is invalid with a message", func() {
			sut := makeSut()
			sut.Data.FrequencyPenalty = -2.1
			msg, err := sut.Sut()
			s.Error(err)
			s.Nil(msg)
			s.Equal("invalid frequencyPenalty", err.Error())
		})

		s.Run("should return an error if model is invalid with a message", func() {
			sut := makeSut()
			sut.Data.Model = nil
			msg, err := sut.Sut()
			s.Error(err)
			s.Nil(msg)
			s.Equal("model is nil", err.Error())
		})

	})
}
