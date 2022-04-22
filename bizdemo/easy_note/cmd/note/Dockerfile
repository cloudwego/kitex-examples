FROM golang:1.16.6-alpine

RUN apk update && apk add git

WORKDIR /code

ADD . .
RUN go env -w GOPROXY=https://goproxy.io,direct
RUN go env -w GO111MODULE=on

RUN cd cmd/note/ && sh build.sh

CMD cd cmd/note/output && sh bootstrap.sh