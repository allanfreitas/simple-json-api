FROM golang:1.8.0-alpine

COPY . ~/go/api/
WORKDIR ~/go/api/

EXPOSE 8080
ENTRYPOINT go run api.go