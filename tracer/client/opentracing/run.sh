#!/bin/sh
export JAEGER_DISABLED=false
export JAEGER_SAMPLER_TYPE="const"
export JAEGER_SAMPLER_PARAM=1
export JAEGER_REPORTER_LOG_SPANS=true
export JAEGER_AGENT_HOST="127.0.0.1"
export JAEGER_AGENT_PORT=6831
go run ./client/opentracing/main.go