package fake

import (
	"fmt"

	"github.com/andre-ols/chatservice/internal/domain/entity"
	stub "github.com/andre-ols/chatservice/testutils/stub/domain/entity"

	"github.com/go-faker/faker/v4"
)

func Message() *entity.Message {
	message := new(entity.Message)

	faker.FakeData(message)

	message.Model = Model()
	message.Tokenizer = &stub.TokenCounterStub{}

	if err := message.Validate(); err != nil {
		panic(fmt.Errorf("error while generating fake message: %w", err))
	}

	return message
}
