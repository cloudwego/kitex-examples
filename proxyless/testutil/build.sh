#!/bin/bash

# build
go mod tidy
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o output/bin/test-controller