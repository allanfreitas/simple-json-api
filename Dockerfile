FROM golang:1.8.0-alpine

COPY . ~/go/api/
WORKDIR ~/go/api/

RUN go run api.go