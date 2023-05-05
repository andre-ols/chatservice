package fake

import (
	"github.com/andre-ols/chatservice/internal/domain/entity"
	"github.com/go-faker/faker/v4"
)

func ChatConfig() *entity.ChatConfig {
	chatConfig := new(entity.ChatConfig)
	faker.FakeData(chatConfig)

	chatConfig.Model = Model()

	return chatConfig
}
