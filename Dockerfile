# Use imagem oficial do Go 1.23.0 para build
FROM golang:1.23.0-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o api ./cmd/api

FROM alpine:latest

RUN apk --no-cache add ca-certificates

COPY --from=build /app/api /api

EXPOSE 8080

ENTRYPOINT ["/api"]
