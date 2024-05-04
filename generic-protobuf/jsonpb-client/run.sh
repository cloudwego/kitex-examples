#!/bin/bash

# 设置工作目录为项目目录
cd ./

REPO_PATH="."

project="generic-protobuf"

echo "---------------------------------------"
echo "Running project: $project"

# 启动 server 和 client
cd "$REPO_PATH" || exit
go run main.go > /dev/null 2>&1 &
pid=$!
cd - > /dev/null || exit

# 当脚本退出时，停止 server 和 client
trap 'kill $pid ' EXIT

# 等待 server 和 client 结束
wait $pid