package entity_test

import (
	"testing"

	"github.com/andre-ols/chatservice/internal/domain/entity"
	fake "github.com/andre-ols/chatservice/testutils/fake/domain/entity"
	"github.com/andre-ols/chatservice/testutils/suite"
)

type ModelTestSuite struct {
	suite.Suite
}

func TestModelSuite(t *testing.T) {
	suite.RunUnitTest(t, new(ModelTestSuite))
}

func (s *ModelTestSuite) TestNewModel() {

	type Sut struct {
		Sut  func() (*entity.Model, error)
		Data *entity.Model
	}

	makeSut := func() *Sut {
		data := fake.Model()

		sut := func() (*entity.Model, error) {
			return entity.NewModel(data.Name, data.MaxTokens)
		}

		return &Sut{Sut: sut, Data: data}
	}

	s.Run("should return a error if name is empty with a message", func() {
		sut := makeSut()
		sut.Data.Name = ""
		m, err := sut.Sut()
		s.Error(err)
		s.Nil(m)
		s.Equal("name is empty", err.Error())
	})

	s.Run("should return a error if maxTokens is less than or equal to 0 with a message", func() {
		sut := makeSut()
		sut.Data.MaxTokens = -1
		m, err := sut.Sut()
		s.Error(err)
		s.Nil(m)
		s.Equal("maxTokens is less than or equal to 0", err.Error())
	})

	s.Run("should return a new model", func() {
		sut := makeSut()
		m, err := sut.Sut()
		s.NoError(err)
		s.NotNil(m)
		s.Equal(sut.Data.Name, m.Name)
		s.Equal(sut.Data.MaxTokens, m.MaxTokens)
	})
}

func (s *ModelTestSuite) TestModel_GetName() {
	s.Run("should return the name", func() {
		m, _ := entity.NewModel("davinci", 1024)
		s.Equal("davinci", m.GetName())
	})
}

func (s *ModelTestSuite) TestModel_GetMaxTokens() {
	s.Run("should return the maxTokens", func() {
		m, _ := entity.NewModel("davinci", 1024)
		s.Equal(1024, m.GetMaxTokens())
	})
}
