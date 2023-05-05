   ![Badge](https://img.shields.io/static/v1?label=Go&message=1.20&color=blue&style=flat&logo=go)
   ![Badge](https://img.shields.io/static/v1?label=Docker&message=23.0.5&color=blue&style=flat&logo=docker)
   ![Badge](https://img.shields.io/static/v1?label=MySQL&message=8&color=blue&style=flat&logo=mysql)

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

## 🎲 Running the Application

Make a copy of the `env.example` file named `.env` inside the `chatservice` folder. Enter your OpenAI API Key in `OPENAI_API_KEY` inside the `.env` file. You can get an OpenAI API Key [by clicking here](https://platform.openai.com/account/api-keys).

### Clone the project

```bash
# Clone this repository
git clone git@github.com:andre-ols/chatservice.git

# Access the project folder
cd chatservice

# Copy the env.example file to .env
cp env.example .env
```

### Using Docker

```bash
docker compose up -d
docker compose exec chatservice bash
go run cmd/chatservice/main.go
```

> _If you choose to run the chatservice microservice using Docker, make sure to change the value of DB_HOST to `DB_HOST=mysql` inside your `.env` file_

### Running locally

```bash
go run cmd/chatservice/main.go
```

> _If you choose to run the chatservice microservice locally, make sure to change the value of DB_HOST to `DB_HOST=localhost` inside your `.env` file_

### migrate

On the first run it will be necessary to apply the `migrate` to create the tables in the MySQL database, through the `Makefile`.

```bash
make migrate
```

> *When doing `make migrate` make sure that the MySQL connection string inside the `Makefile` points to *mysql:3306* when the chatservice is running in Docker, or *localhost:3306* when the chatservice is running locally.*

## Additional information

On Windows, use the Ubuntu terminal with WSL 2 to run the commands.\
For more details, see Full Cycle's [WSL2 + Docker Quick Start](https://github.com/codeedu/wsl2-docker-quickstart).

## 🛠 Technologies Used

- [Go](https://golang.org/)
- [Docker](https://www.docker.com/)
- [MySQL](https://www.mysql.com/)
- [gRPC](https://grpc.io/)
- [Go-OpenAI](https://github.com/sashabaranov/go-openai)
- [Go-Faker](https://github.com/go-faker/faker/v4)
- [Tiktoken](https://github.com/pkoukk/tiktoken-go)
- [Testify](https://github.com/stretchr/testify)
- [Go-Chi](github.com/go-chi/chi/v5)

### Author

<br />

[![Linkedin Badge](https://img.shields.io/badge/-André-blue?style=flat-square&logo=Linkedin&logoColor=white&link=https://www.linkedin.com/in/andre-ols/)](https://www.linkedin.com/in/andre-ols/)
[![Gmail Badge](https://img.shields.io/badge/-contato.andreols@gmail.com-c14438?style=flat-square&logo=Gmail&logoColor=white&link=mailto:contato.andreols@gmail.com)](mailto:contato.andreols@gmail.com)

## License

This project is under the license [MIT](./LICENSE).

<p>Made with ❤️ by <a href="https://www.linkedin.com/in/andre-ols/">André Oliveira</a>!</p>
