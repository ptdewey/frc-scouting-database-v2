FROM golang:1.22-alpine

WORKDIR /app

COPY ./go.mod .
RUN go mod download

COPY . .
# COPY ./cmd/app /app/cmd/app
# COPY ./config/app.env /app/config/app.env
# COPY ./pkg/app /app/pkg/app
# COPY ./internal /app/internal

ARG UID
ARG GID

RUN addgroup -g $GID appgroup && \
    adduser -u $UID -G appgroup -D appuser
RUN chown -R appuser:appgroup /app
USER appuser

RUN ls -la /app/cmd/app
RUN pwd

ENTRYPOINT [ "go", "run", "/app/cmd/app/main.go" ]
