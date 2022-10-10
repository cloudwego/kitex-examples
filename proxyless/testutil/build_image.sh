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

# go build
sh ./build.sh

# docker build
minikube image rm kitex-example/proxyless-test-controller # delete old image
minikube image build -t kitex-example/proxyless-test-controller .

# Push to public registry
#docker image rm liamzqh/kitex-proxyless-test-controller # delete old image
#docker build -t liamzqh/kitex-proxyless-test-controller . # build
#docker push liamzqh/kitex-proxyless-test-controller # push
