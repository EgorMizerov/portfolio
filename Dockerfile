FROM golang:latest

RUN mkdir -p /usr/src/app/
WORKDIR /usr/src/app

COPY . /usr/src/app/

RUN go mod download
EXPOSE 80
ENTRYPOINT go run cmd/main.go