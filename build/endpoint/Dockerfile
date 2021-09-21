FROM golang:latest
ENV GO111MODULE=on

RUN mkdir /app
WORKDIR /app


COPY ./internal/ internal/
COPY ./configs/ configs/

WORKDIR /app/internal/endpoint/

CMD go run main.go

