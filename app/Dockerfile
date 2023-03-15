FROM golang:1.20-alpine

RUN go install github.com/cosmtrek/air@latest

WORKDIR /go/app

COPY ./app .

CMD go build main.go && ./main
