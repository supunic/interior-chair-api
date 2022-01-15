FROM golang:1.17.2-alpine

WORKDIR /go/src

RUN go get github.com/cosmtrek/air@v1.27.3 && \
    go get -v github.com/rubenv/sql-migrate/...
