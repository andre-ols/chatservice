package service

import (
	"github.com/andre-ols/chatservice/internal/infra/grpc/pb"
	"github.com/andre-ols/chatservice/internal/usecase/chatcompletionstream"
)

type ChatService struct {
	pb.UnimplementedChatServiceServer
	ChatCompletionStreamUseCase chatcompletionstream.ChatCompletionUseCase
	ChatConfigStream            chatcompletionstream.ChatCompletionConfigInputDto
	StreamChanel                chan chatcompletionstream.ChatCompletionOutputDto
}

func NewChatService(chatCompletionStream chatcompletionstream.ChatCompletionUseCase, chatConfigStream chatcompletionstream.ChatCompletionConfigInputDto, streamChanel chan chatcompletionstream.ChatCompletionOutputDto) *ChatService {
	return &ChatService{
		ChatCompletionStreamUseCase: chatCompletionStream,
		ChatConfigStream:            chatConfigStream,
		StreamChanel:                streamChanel,
	}
}

func (c *ChatService) ChatStream(req *pb.ChatRequest, stream pb.ChatService_ChatStreamServer) error {
	chatConfig := chatcompletionstream.ChatCompletionConfigInputDto{
		Model:                c.ChatConfigStream.Model,
		ModelMaxTokens:       c.ChatConfigStream.ModelMaxTokens,
		Temperature:          c.ChatConfigStream.Temperature,
		TopP:                 c.ChatConfigStream.TopP,
		N:                    c.ChatConfigStream.N,
		Stop:                 c.ChatConfigStream.Stop,
		MaxTokens:            c.ChatConfigStream.MaxTokens,
		PresencePenalty:      c.ChatConfigStream.PresencePenalty,
		FrequencyPenalty:     c.ChatConfigStream.FrequencyPenalty,
		InitialSystemMessage: c.ChatConfigStream.InitialSystemMessage,
	}

	input := chatcompletionstream.ChatCompletionInputDto{
		UserMessage: req.GetUserMessage(),
		UserID:      req.GetUserId(),
		ChatID:      req.GetChatId(),
		Config:      chatConfig,
	}

	ctx := stream.Context()

	go func() {
		for msg := range c.StreamChanel {
			stream.Send(&pb.ChatResponse{
				ChatId:  msg.ChatID,
				UserId:  msg.UserID,
				Content: msg.Content,
			})
		}
	}()

	_, err := c.ChatCompletionStreamUseCase.Execute(ctx, input)

	if err != nil {
		return err
	}

	return nil

}
