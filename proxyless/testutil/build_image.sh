#!/bin/bash

# go build
sh ./build.sh

# docker build
minikube image rm kitex-example/proxyless-test-controller # delete old image
minikube image build -t kitex-example/proxyless-test-controller .

# Push to public registry
#docker image rm liamzqh/kitex-proxyless-test-controller # delete old image
#docker build -t liamzqh/kitex-proxyless-test-controller . # build
#docker push liamzqh/kitex-proxyless-test-controller # push
