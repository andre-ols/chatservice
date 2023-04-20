package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/andre-ols/chatservice/internal/domain/entity"
	"github.com/andre-ols/chatservice/internal/infra/db"
)

type ChatRepositoryMySQL struct {
	DB      *sql.DB
	Queries *db.Queries
}

func (c *ChatRepositoryMySQL) CreateChat(ctx context.Context, chat *entity.Chat) error {

	err := c.Queries.CreateChat(
		ctx,
		db.CreateChatParams{
			ID:               chat.ID,
			UserID:           chat.UserID,
			InitialMessageID: chat.InitialSystemMessage.Content,
			Status:           string(chat.Status),
			TokenUsage:       int32(chat.TokenUsage),
			Model:            chat.Config.Model.Name,
			ModelMaxTokens:   int32(chat.Config.Model.MaxTokens),
			Temperature:      float64(chat.Config.Temperature),
			TopP:             float64(chat.Config.TopP),
			N:                int32(chat.Config.N),
			Stop:             chat.Config.Stop[0],
			MaxTokens:        int32(chat.Config.MaxTokens),
			PresencePenalty:  float64(chat.Config.PresencePenalty),
			FrequencyPenalty: float64(chat.Config.FrequencyPenalty),
			CreatedAt:        time.Now(),
			UpdatedAt:        time.Now(),
		},
	)
	if err != nil {
		return err
	}

	err = c.Queries.AddMessage(
		ctx,
		db.AddMessageParams{
			ID:        chat.InitialSystemMessage.ID,
			ChatID:    chat.ID,
			Content:   chat.InitialSystemMessage.Content,
			Role:      string(chat.InitialSystemMessage.Role),
			Tokens:    int32(chat.InitialSystemMessage.Tokens),
			CreatedAt: chat.InitialSystemMessage.CreatedAt,
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func (c *ChatRepositoryMySQL) FindChatByID(ctx context.Context, id string) (*entity.Chat, error) {
	chat := &entity.Chat{}

	row, err := c.Queries.FindChatByID(ctx, id)

	if err != nil {
		return nil, errors.New("chat not found")
	}

	chat.ID = row.ID
	chat.UserID = row.UserID
	chat.Status = entity.Status(row.Status)
	chat.TokenUsage = int(row.TokenUsage)
	chat.Config = &entity.ChatConfig{
		Model: &entity.Model{
			Name:      row.Model,
			MaxTokens: int(row.ModelMaxTokens),
		},
		Temperature: float32(row.Temperature),
		TopP:        float32(row.TopP),

		N:                int(row.N),
		Stop:             []string{row.Stop},
		MaxTokens:        int(row.MaxTokens),
		PresencePenalty:  float32(row.PresencePenalty),
		FrequencyPenalty: float32(row.FrequencyPenalty),
	}

	messages, err := c.Queries.FindMessagesByChatID(ctx, id)

	if err != nil {
		return nil, err
	}

	for _, m := range messages {
		chat.Messages = append(chat.Messages, &entity.Message{
			ID:        m.ID,
			Content:   m.Content,
			Role:      entity.Role(m.Role),
			Tokens:    int(m.Tokens),
			Model:     &entity.Model{Name: m.Model},
			CreatedAt: m.CreatedAt,
		})
	}

	erased, err := c.Queries.FindErasedMessagesByChatID(ctx, id)

	if err != nil {
		return nil, err
	}

	for _, m := range erased {
		chat.ErasedMessages = append(chat.ErasedMessages, &entity.Message{
			ID:        m.ID,
			Content:   m.Content,
			Role:      entity.Role(m.Role),
			Tokens:    int(m.Tokens),
			Model:     &entity.Model{Name: m.Model},
			CreatedAt: m.CreatedAt,
		})
	}

	return chat, nil

}

func (c *ChatRepositoryMySQL) SaveChat(ctx context.Context, chat *entity.Chat) error {
	params := db.SaveChatParams{
		ID:               chat.ID,
		UserID:           chat.UserID,
		Status:           string(chat.Status),
		TokenUsage:       int32(chat.TokenUsage),
		Model:            chat.Config.Model.Name,
		ModelMaxTokens:   int32(chat.Config.Model.MaxTokens),
		Temperature:      float64(chat.Config.Temperature),
		TopP:             float64(chat.Config.TopP),
		N:                int32(chat.Config.N),
		Stop:             chat.Config.Stop[0],
		MaxTokens:        int32(chat.Config.MaxTokens),
		PresencePenalty:  float64(chat.Config.PresencePenalty),
		FrequencyPenalty: float64(chat.Config.FrequencyPenalty),
		UpdatedAt:        time.Now(),
	}

	err := c.Queries.SaveChat(
		ctx,
		params,
	)
	if err != nil {
		return err
	}
	// delete messages
	err = c.Queries.DeleteChatMessages(ctx, chat.ID)
	if err != nil {
		return err
	}
	// delete erased messages
	err = c.Queries.DeleteErasedChatMessages(ctx, chat.ID)
	if err != nil {
		return err
	}
	// save messages
	i := 0
	for _, message := range chat.Messages {
		err = c.Queries.AddMessage(
			ctx,
			db.AddMessageParams{
				ID:        message.ID,
				ChatID:    chat.ID,
				Content:   message.Content,
				Role:      string(message.Role),
				Tokens:    int32(message.Tokens),
				Model:     chat.Config.Model.Name,
				CreatedAt: message.CreatedAt,
				OrderMsg:  int32(i),
				Erased:    false,
			},
		)
		if err != nil {
			return err
		}
		i++
	}
	// save erased messages
	i = 0
	for _, message := range chat.ErasedMessages {
		err = c.Queries.AddMessage(
			ctx,
			db.AddMessageParams{
				ID:        message.ID,
				ChatID:    chat.ID,
				Content:   message.Content,
				Role:      string(message.Role),
				Tokens:    int32(message.Tokens),
				Model:     chat.Config.Model.Name,
				CreatedAt: message.CreatedAt,
				OrderMsg:  int32(i),
				Erased:    true,
			},
		)
		if err != nil {
			return err
		}
		i++
	}
	return nil
}

func NewChatRepositoryMySQL(dbt *sql.DB) *ChatRepositoryMySQL {
	return &ChatRepositoryMySQL{
		DB:      dbt,
		Queries: db.New(dbt),
	}
}
