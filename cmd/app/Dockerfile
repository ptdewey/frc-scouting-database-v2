FROM golang:1.22-alpine

WORKDIR /app

COPY ./go.mod .
RUN go mod download

RUN apk --no-cache add tzdata
ENV TZ=America/New_York

COPY . .

# ARG GID
# ARG UID

# RUN addgroup -g $GID appgroup && \
#     adduser -u $UID -G appgroup -D appuser
# RUN chown -R appuser:appgroup /app
# USER appuser

ENTRYPOINT [ "go", "run", "/app/cmd/app/main.go" ]
