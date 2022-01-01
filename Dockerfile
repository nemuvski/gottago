FROM golang:1.17-alpine

RUN apk add --no-cache git

RUN mkdir /go/src/workspace
WORKDIR /go/src/workspace
ADD ./workspace /go/src/workspace
