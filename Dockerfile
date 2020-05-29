FROM golang:latest

WORKDIR /go/src/app

RUN go get github.com/go-sql-driver/mysql

COPY . .

ENTRYPOINT go run be-svc/main.go