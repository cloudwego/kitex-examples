#!/bin/bash

#
# Copyright 2022 CloudWeGo Authors
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#

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

