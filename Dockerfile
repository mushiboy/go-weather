FROM golang:alpine3.20 AS builder

WORKDIR /app

COPY . .

RUN cd cmd/goweather && go build -o /goweather

FROM alpine:latest

RUN apk add --no-cache bash

COPY --from=builder /goweather /usr/local/bin/goweather

ENTRYPOINT [ "/bin/bash" ]