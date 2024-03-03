FROM golang:1.22-alpine

WORKDIR /bot

COPY ./go.mod .
RUN go mod download

RUN apk --no-cache add tzdata
ENV TZ=America/New_York

COPY . .

ENTRYPOINT [ "go", "run", "cmd/bot/main.go"]
