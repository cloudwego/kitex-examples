#!/bin/bash

set -e
CURDIR=$(cd $(dirname $0); pwd)

# build
go mod tidy
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $CURDIR/output/bin/kitex $CURDIR
#go build -o output/bin/kitex
cp $CURDIR/bootstrap.sh $CURDIR/output/