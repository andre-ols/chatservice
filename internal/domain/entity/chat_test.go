package entity_test

import (
	"testing"

	"github.com/andre-ols/chatservice/internal/domain/entity"
	fake "github.com/andre-ols/chatservice/testutils/fake/domain/entity"
	"github.com/andre-ols/chatservice/testutils/suite"
)

type ChatTestSuite struct {
	suite.Suite
}

func TestChatSuite(t *testing.T) {
	suite.RunUnitTest(t, new(ChatTestSuite))
}

func (s *ChatTestSuite) TestNewChat() {
	type Sut struct {
		Sut  func() (*entity.Chat, error)
		Data *entity.Chat
	}

	makeSut := func() *Sut {
		data := fake.Chat()

		sut := func() (*entity.Chat, error) {
			return entity.NewChat(data.UserID, data.InitialSystemMessage, data.Config)
		}

		return &Sut{Sut: sut, Data: data}
	}

	s.Run("TestNewChat Error", func() {

		s.Run("should return an error if userID is empty", func() {
			sut := makeSut()
			sut.Data.UserID = ""
			chat, err := sut.Sut()

			s.Error(err)
			s.Nil(chat)
			s.Equal("user id is empty", err.Error())

		})

		s.Run("should return an error if config is nil with a message", func() {
			sut := makeSut()
			sut.Data.Config = nil
			chat, err := sut.Sut()
			s.Error(err)
			s.Nil(chat)
			s.Equal("config is nil", err.Error())
		})

		s.Run("should return an error if initial system message is nil with a message", func() {
			sut := makeSut()
			sut.Data.InitialSystemMessage = nil
			chat, err := sut.Sut()
			s.Error(err)
			s.Nil(chat)
			s.Equal("initial system message is nil", err.Error())
		})

	})

	s.Run("TestNewChat Success", func() {
		sut := makeSut()
		chat, err := sut.Sut()
		s.NoError(err)
		s.NotNil(chat)
		s.Equal(sut.Data.UserID, chat.UserID)
		s.Equal(sut.Data.InitialSystemMessage, chat.InitialSystemMessage)
		s.Equal(sut.Data.Config, chat.Config)
		s.Equal(entity.Active, chat.Status)
		s.Equal(sut.Data.InitialSystemMessage.Tokens, chat.TokenUsage)
		s.Equal(1, len(chat.Messages))
		s.Equal(0, len(chat.ErasedMessages))
	})
}
