FROM golang:1.16.6-alpine

RUN apk update && apk add git

WORKDIR /code

ADD . .
RUN go env -w GOPROXY=https://goproxy.io,direct
RUN go env -w GO111MODULE=on

RUN cd cmd/user/ && sh build.sh

CMD cd cmd/user/output && sh bootstrap.sh