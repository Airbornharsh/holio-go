# Build stage
FROM golang:1.21 AS builder
WORKDIR /app
COPY . .
RUN go mod download && go mod verify

RUN go get github.com/pilu/fresh

EXPOSE 8080

CMD ["fresh"] 