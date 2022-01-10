FROM golang:1.17.2-alpine

RUN apk update && apk add git

RUN go get github.com/labstack/echo/v4 && \
  go get -u gorm.io/gorm && \
  go get -u gorm.io/driver/mysql
