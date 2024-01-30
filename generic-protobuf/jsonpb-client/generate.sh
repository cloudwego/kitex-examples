#!/bin/bash

cd `dirname $0`

kitex=${1:-kitex}

set -x

$kitex -module jsonpb-demo -service jsonpb-demo-service idl/api.proto
