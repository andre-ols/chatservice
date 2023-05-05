FROM golang:1.20.3

# Set environment variables
ENV PATH="/root/.cargo/bin:${PATH}"
ENV USER=root

WORKDIR /go/src
RUN ln -sf /bin/bash /bin/sh
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .

RUN apt update && apt install -y protobuf-compiler && apt install -y protoc-gen-go
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

CMD [ "tail", "-f", "/dev/null" ]