#!/bin/bash
# Copyright 2024 CloudWeGo Authors
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

cd `dirname $0`
DIR=`pwd`

module='-module github.com/cloudwego/kitex-examples'
service='-service demo-server'
idl=api.thrift

kitex=${1:-kitex}

thriftgo_version=`thriftgo -version 2>&1`
kitex_version=`$kitex -version 2>&1`

echo "thriftgo: using `which thriftgo` v${thriftgo_version/thriftgo /}"
echo "kitex: using `which $kitex` $kitex_version"
echo -e "\nMake sure you're using kitex >= v0.9.0 && thriftgo >= v0.3.6\n"

set -x
$kitex -streamx $verbose $module $service $idl
