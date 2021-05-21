FROM golang:1.15-alpine

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

RUN apk update && apk add --no-cache musl-dev libpcap-dev gcc

WORKDIR /app

COPY . .
