FROM golang:latest

RUN mkdir /app
WORKDIR /app


COPY ./internal/  internal/
COPY ./configs/  configs/

WORKDIR /app/internal/worker

CMD go run main.go

