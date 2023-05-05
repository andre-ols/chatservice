package fake

import (
	"github.com/andre-ols/chatservice/internal/domain/entity"
	"github.com/go-faker/faker/v4"
)

func Chat() *entity.Chat {
	chat := new(entity.Chat)

	faker.FakeData(chat)

	chat.Config = ChatConfig()
	chat.InitialSystemMessage = Message()
	chat.Messages = []*entity.Message{Message()}
	chat.ErasedMessages = []*entity.Message{Message()}

	if err := chat.Validate(); err != nil {
		panic(err)
	}

	return chat
}
