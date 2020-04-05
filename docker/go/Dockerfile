FROM golang:1.13.7 as build

RUN apt-get update && \
    apt-get install -y --no-install-recommends git

WORKDIR /go/src
COPY ./src /go/src

RUN go build -o go \
    && go get github.com/oxequa/realize \