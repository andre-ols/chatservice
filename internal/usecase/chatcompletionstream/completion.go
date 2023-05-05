package chatcompletionstream

import (
	"context"
	"errors"
	"io"
	"strings"

	"github.com/andre-ols/chatservice/internal/domain/adapter"
	"github.com/andre-ols/chatservice/internal/domain/entity"
	"github.com/andre-ols/chatservice/internal/domain/gateway"
	"github.com/sashabaranov/go-openai"
)

type ChatCompletionConfigInputDto struct {
	Model                string
	ModelMaxTokens       int
	Temperature          float32
	TopP                 float32
	N                    int
	Stop                 []string
	MaxTokens            int
	PresencePenalty      float32
	FrequencyPenalty     float32
	InitialSystemMessage string
}

type ChatCompletionInputDto struct {
	ChatID      string
	UserID      string
	UserMessage string
	Config      ChatCompletionConfigInputDto
}

type ChatCompletionOutputDto struct {
	ChatID  string
	UserID  string
	Content string
}

type ChatCompletionUseCase struct {
	chatGateway  gateway.ChatGateway
	OpenAiClient *openai.Client
	Stream       chan ChatCompletionOutputDto
}

func NewChatCompletionUseCase(chatGateway gateway.ChatGateway, openAiClient *openai.Client, stream chan ChatCompletionOutputDto) *ChatCompletionUseCase {
	return &ChatCompletionUseCase{
		chatGateway:  chatGateway,
		OpenAiClient: openAiClient,
		Stream:       stream,
	}
}

func (uc *ChatCompletionUseCase) Execute(ctx context.Context, input ChatCompletionInputDto) (*ChatCompletionOutputDto, error) {
	chat, err := uc.chatGateway.FindChatByID(ctx, input.ChatID)
	if err != nil {

		if err.Error() == "chat not found" {
			// Create new chat (Entity)
			chat, err = uc.createNewChat(ctx, input)
			if err != nil {
				return nil, errors.New("error creating new chat: " + err.Error())
			}

			// Save on database
			err = uc.chatGateway.CreateChat(ctx, chat)
		} else {
			return nil, errors.New("error fetching existing chat: " + err.Error())
		}
	}

	tokenCounter := adapter.NewTokenCounter()

	userMessage, err := entity.NewMessage("user", input.UserMessage, chat.Config.Model, tokenCounter)

	if err != nil {
		return nil, errors.New("error creating user message: " + err.Error())
	}

	err = chat.AddMessage(userMessage)

	if err != nil {
		return nil, errors.New("error adding user message to chat: " + err.Error())
	}

	messages := []openai.ChatCompletionMessage{}

	for _, message := range chat.Messages {
		messages = append(messages, openai.ChatCompletionMessage{
			Role:    string(message.Role),
			Content: message.Content,
		})
	}

	res, err := uc.OpenAiClient.CreateChatCompletionStream(ctx, openai.ChatCompletionRequest{
		Model:            chat.Config.Model.Name,
		Messages:         messages,
		MaxTokens:        chat.Config.MaxTokens,
		Temperature:      chat.Config.Temperature,
		TopP:             chat.Config.TopP,
		N:                chat.Config.N,
		Stop:             chat.Config.Stop,
		PresencePenalty:  chat.Config.PresencePenalty,
		FrequencyPenalty: chat.Config.FrequencyPenalty,
		Stream:           true,
	})

	if err != nil {
		return nil, errors.New("error creating chat completion: " + err.Error())
	}

	var fullResponse strings.Builder

	for {
		response, err := res.Recv()

		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			return nil, errors.New("error receiving chat completion response: " + err.Error())
		}

		fullResponse.WriteString(response.Choices[0].Delta.Content)

		r := ChatCompletionOutputDto{
			ChatID:  chat.ID,
			UserID:  input.UserID,
			Content: fullResponse.String(),
		}

		uc.Stream <- r
	}

	assistant, err := entity.NewMessage("assistant", fullResponse.String(), chat.Config.Model, tokenCounter)

	if err != nil {
		return nil, errors.New("error creating assistant message: " + err.Error())
	}

	err = chat.AddMessage(assistant)

	if err != nil {
		return nil, errors.New("error adding assistant message to chat: " + err.Error())
	}

	err = uc.chatGateway.SaveChat(ctx, chat)

	if err != nil {
		return nil, errors.New("error saving chat: " + err.Error())
	}

	return &ChatCompletionOutputDto{
		ChatID:  chat.ID,
		UserID:  input.UserID,
		Content: fullResponse.String(),
	}, nil

}

func (uc *ChatCompletionUseCase) createNewChat(ctx context.Context, input ChatCompletionInputDto) (*entity.Chat, error) {
	// Create new chat (Entity)

	model, err := entity.NewModel(input.Config.Model, input.Config.ModelMaxTokens)

	if err != nil {
		return nil, errors.New("error creating new model: " + err.Error())
	}

	chatConfig := &entity.ChatConfig{
		Model:            model,
		Temperature:      input.Config.Temperature,
		TopP:             input.Config.TopP,
		N:                input.Config.N,
		Stop:             input.Config.Stop,
		MaxTokens:        input.Config.MaxTokens,
		PresencePenalty:  input.Config.PresencePenalty,
		FrequencyPenalty: input.Config.FrequencyPenalty,
	}

	tokenCounter := adapter.NewTokenCounter()

	initialSystemMessage, err := entity.NewMessage("system", input.Config.InitialSystemMessage, model, tokenCounter)

	if err != nil {
		return nil, errors.New("error creating initial system message: " + err.Error())
	}

	chat, err := entity.NewChat(input.UserID, initialSystemMessage, chatConfig)

	if err != nil {
		return nil, errors.New("error creating new chat: " + err.Error())
	}

	return chat, nil
}
