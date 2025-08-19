FROM golang:1.24.2-alpine3.21 AS builder

WORKDIR /app

COPY go.* .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o app .

FROM alpine:3.21 AS alpine-app

WORKDIR /app

COPY --from=builder /app/app .

RUN adduser -D programuser

USER programuser

ENTRYPOINT ["./app"]

FROM golang:1.24.2-alpine3.21 AS dev

WORKDIR /app

RUN apk --no-cache update && apk upgrade && \
    apk add --no-cache curl make git

COPY go.* .

RUN go mod download

RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest && \
    go install github.com/matryer/moq@latest

ENV PATH="$PATH:/go/bin"

COPY . .

# Build to /usr/local/bin instead of /app
# to avoid the volume mount (.:/app) hiding built binary
RUN CGO_ENABLED=0 go build -o /usr/local/bin/app .

CMD ["app"]
