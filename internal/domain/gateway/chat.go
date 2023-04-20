package gateway

import (
	"context"

	"github.com/andre-ols/chatservice/internal/domain/entity"
)

type ChatGateway interface {
	CreateChat(ctx context.Context, chat *entity.Chat) error
	FindChatByID(ctx context.Context, id string) (*entity.Chat, error)
	SaveChat(ctx context.Context, chat *entity.Chat) error
}
