
<a href="https://openai.com/blog/chatgpt/">
   <img src="https://freelogopng.com/images/all_img/1681038628chatgpt-icon-logo.png" alt="ChatGPT" title="ChatGPT" align="right" height="60" />
</a>

# ChatService

ChatService is a Golang project that provides an interactive chat service based on ChatGPT using Open AI's API. All user conversations are stored and managed in a MySQL database, allowing for easy data analysis and monitoring. The architecture follows clean architecture principles, separating the layers for easy maintenance and evolution. The service can be accessed through HTTP or gRPC (stream), and Docker is used for easy installation and execution.

## Documentation

Documentation for the ChatService microservice can be accessed through Postman at https://www.postman.com/blue-rocket-958887/workspace/chatservice

## Installing the prerequisites

Go 1.20\
https://go.dev/dl/

Docker\
https://docs.docker.com/get-docker/

sqlc
```bash
sudo snap install sqlc
```

migrate to Go
```bash
curl -s https://packagecloud.io/install/repositories/golang-migrate/migrate/script.deb.sh | sudo bash
sudo apt-get update
sudo apt-get install migrate
```

protoc\
https://grpc.io/docs/protoc-installation/

## How to run the project

Make a copy of the `env.example` file named `.env` inside the `chatservice` folder. Enter your OpenAI API Key in `OPENAI_API_KEY` inside the `.env` file. You can get an OpenAI API Key [by clicking here](https://platform.openai.com/account/api-keys).

### Clone the project

```bash
git clone git@github.com:andre-ols/chatservice.git
```

### Using Docker

```bash
cd chatservice
docker compose up -d
docker compose exec chatservice bash 
go run cmd/chatservice/main.go
```

> *If you choose to run the chatservice microservice using Docker, make sure to change the value of DB_HOST to `DB_HOST=mysql` inside your `.env` file*

### Running locally

```bash
cd chatservice
go run cmd/chatservice/main.go
```

> *If you choose to run the chatservice microservice locally, make sure to change the value of DB_HOST to `DB_HOST=localhost` inside your `.env` file*

### migrate

On the first run it will be necessary to apply the `migrate` to create the tables in the MySQL database, through the `Makefile`.

```bash
cd chatservice
make migrate
```

> *When doing `make migrate` make sure that the MySQL connection string inside the `Makefile` points to *mysql:3306* when the chatservice is running in Docker, or *localhost:3306* when the chatservice is running locally.*

## Additional information

On Windows, use the Ubuntu terminal with WSL 2 to run the commands.\
For more details, see Full Cycle's [WSL2 + Docker Quick Start](https://github.com/codeedu/wsl2-docker-quickstart).