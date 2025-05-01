FROM golang:1.24-alpine AS base

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN rm .env 

RUN mv .env.prod .env

RUN go build -o go-portifolio

ENV GIN_MODE=release

EXPOSE 2053

CMD ["/build/go-portifolio"]

