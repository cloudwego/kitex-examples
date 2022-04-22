FROM golang:1.16.6-alpine

RUN apk update && apk add git

WORKDIR /code

ADD . .
RUN go env -w GOPROXY=https://goproxy.io,direct
RUN go env -w GO111MODULE=on
RUN cd cmd/api && go build -o api-server

CMD cd cmd/api && ./api-server