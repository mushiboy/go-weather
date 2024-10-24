FROM golang:alpine3.20 AS builder

WORKDIR /app

COPY . .

RUN cd cmd/goweather && go build -o /goweather

FROM alpine:latest

RUN apk add --no-cache bash

COPY --from=builder /goweather /usr/local/bin/goweather

# Create the directory structure for the config file
RUN mkdir -p /usr/local/etc/goweather/internal/config

# Copy the config file to the correct location
COPY --from=builder /app/internal/config/config.json /usr/local/etc/goweather/internal/config/config.json

EXPOSE 8080

ENTRYPOINT [ "/bin/bash"]