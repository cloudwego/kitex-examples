#!/bin/bash

set -e
CURDIR=$(cd $(dirname $0); pwd)

# go build
sh $CURDIR/build.sh

# docker build
# For minikube
minikube image rm kitex-example/proxyless # delete old image
minikube image build -t kitex-example/proxyless $CURDIR/ # build

# Push to public registry
#docker image rm liamzqh/kitex-proxyless-example # delete old image
#docker build -t liamzqh/kitex-proxyless-example . # build
#docker push liamzqh/kitex-proxyless-example # push

