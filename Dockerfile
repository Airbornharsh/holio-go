# Build stage
FROM golang:1.21.3-alpine3.18 AS builder
WORKDIR /app
COPY . .
RUN go mod download && go mod verify

RUN go build -o holio-go

EXPOSE 8080

CMD ["./holio-go"] 