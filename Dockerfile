FROM ubuntu:latest

RUN apt-get update && apt-get install -y \
    golang-go \
    git \
    && rm -rf /var/lib/apt/lists/*


WORKDIR /app

COPY . /app

RUN cd cmd/goweather && go build -o /usr/local/bin/goweather

ENTRYPOINT [ "/bin/bash" ]

