package main

import (
	"database/sql"
	"fmt"

	"github.com/andre-ols/chatservice/configs"
	"github.com/andre-ols/chatservice/internal/infra/repository"
	"github.com/andre-ols/chatservice/internal/infra/web"
	"github.com/andre-ols/chatservice/internal/infra/web/webserver"
	"github.com/andre-ols/chatservice/internal/usecase/chatcompletion"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sashabaranov/go-openai"
)

func main() {
	configs := configs.LoadConfig(".")

	conn, err := sql.Open(configs.DBDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&multiStatements=true",
		configs.DBUser, configs.DBPassword, configs.DBHost, configs.DBPort, configs.DBName))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	repo := repository.NewChatRepositoryMySQL(conn)
	client := openai.NewClient(configs.OpenAIApiKey)

	chatConfig := chatcompletion.ChatCompletionConfigInputDTO{
		Model:                configs.Model,
		ModelMaxTokens:       configs.ModelMaxTokens,
		Temperature:          float32(configs.Temperature),
		TopP:                 float32(configs.TopP),
		N:                    configs.N,
		Stop:                 configs.Stop,
		MaxTokens:            configs.MaxTokens,
		InitialSystemMessage: configs.InitialChatMessage,
	}

	// chatConfigStream := chatcompletionstream.ChatCompletionConfigInputDto{
	// 	Model:                configs.Model,
	// 	ModelMaxTokens:       configs.ModelMaxTokens,
	// 	Temperature:          float32(configs.Temperature),
	// 	TopP:                 float32(configs.TopP),
	// 	N:                    configs.N,
	// 	Stop:                 configs.Stop,
	// 	MaxTokens:            configs.MaxTokens,
	// 	InitialSystemMessage: configs.InitialChatMessage,
	// }

	chatCompletionUseCase := chatcompletion.NewChatCompletionUseCase(repo, client)

	// streamChannel := make(chan chatcompletionstream.ChatCompletionOutputDto)
	// usecaseStream := chatcompletionstream.NewChatCompletionUseCase(repo, client, streamChannel)

	// fmt.Println("Starting gRPC server on port " + configs.GRPCServerPort)
	// grpcServer := server.NewGRPCServer(
	// 	*usecaseStream,
	// 	chatConfigStream,
	// 	configs.GRPCServerPort,
	// 	configs.AuthToken,
	// 	streamChannel,
	// )
	// go grpcServer.Start()

	webserver := webserver.NewWebServer(":" + configs.WebServerPort)
	webserverChatHandler := web.NewWebChatGPTHandler(*chatCompletionUseCase, chatConfig, configs.AuthToken)
	webserver.AddHandler("/chat", webserverChatHandler.Handle)

	fmt.Println("Server running on port " + configs.WebServerPort)
	webserver.Start()
}
