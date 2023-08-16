# Use a multi-stage build
FROM golang:alpine AS builder

# Install build dependencies
RUN apk --no-cache add gcc libc-dev

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -tags musl -o app ./main.go

# Create the final lightweight image
FROM alpine:latest

RUN apk update && apk upgrade

RUN rm -rf /var/cache/apk/* && \
    rm -rf /tmp/*

RUN adduser -D appuser
USER appuser

WORKDIR /app

COPY --from=builder /app/app .

CMD ["./app"]
