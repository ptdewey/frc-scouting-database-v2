FROM golang:1.22-alpine

WORKDIR /bot

COPY ./go.mod .
RUN go mod download

COPY . .
# COPY ./cmd/bot /bot/cmd/bot
# COPY ./config/bot.env /bot/config/bot.env
# COPY ./internal/bot /bot/internal/bot
# COPY ./internal/utils /bot/internal/utils

ENTRYPOINT [ "go", "run", "cmd/bot/main.go"]
