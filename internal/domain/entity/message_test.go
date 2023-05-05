package entity_test

import (
	"testing"

	"github.com/andre-ols/chatservice/internal/domain/entity"
	fake "github.com/andre-ols/chatservice/testutils/fake/domain/entity"
	"github.com/andre-ols/chatservice/testutils/suite"
)

type MessageTestSuite struct {
	suite.Suite
}

func TestMessageSuite(t *testing.T) {
	suite.RunUnitTest(t, new(MessageTestSuite))
}

func (s *MessageTestSuite) TestNewMessage() {
	type Sut struct {
		Sut  func() (*entity.Message, error)
		Data *entity.Message
	}

	makeSut := func() *Sut {
		data := fake.Message()

		sut := func() (*entity.Message, error) {
			return entity.NewMessage(data.Role, data.Content, data.Model)
		}

		return &Sut{Sut: sut, Data: data}
	}

	s.Run("should return an error if role is invalid with a message", func() {
		sut := makeSut()
		sut.Data.Role = "invalid"
		msg, err := sut.Sut()
		s.Error(err)
		s.Nil(msg)
		s.Equal("invalid role", err.Error())
	})

	s.Run("should return an error if content is empty with a message", func() {
		sut := makeSut()
		sut.Data.Content = ""
		msg, err := sut.Sut()
		s.Error(err)
		s.Nil(msg)
		s.Equal("content is empty", err.Error())
	})

	s.Run("should return an error if model is nil with a message", func() {
		sut := makeSut()
		sut.Data.Model = nil
		msg, err := sut.Sut()
		s.Error(err)
		s.Nil(msg)
		s.Equal("model is nil", err.Error())
	})

	s.Run("should return a new message", func() {
		sut := makeSut()
		msg, err := sut.Sut()
		s.NoError(err)
		s.NotNil(msg)
		s.Equal(sut.Data.Role, msg.Role)
		s.Equal(sut.Data.Content, msg.Content)
		s.Equal(sut.Data.Model, msg.Model)
		s.Equal(sut.Data.CreatedAt, msg.CreatedAt)
	})
}
