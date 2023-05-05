package fake

import (
	"fmt"

	"github.com/andre-ols/chatservice/internal/domain/entity"
	"github.com/go-faker/faker/v4"
)

func Model() *entity.Model {
	model := new(entity.Model)

	faker.FakeData(model)

	if err := model.Validate(); err != nil {
		panic(fmt.Errorf("error while generating fake model: %w", err))
	}

	return model
}
