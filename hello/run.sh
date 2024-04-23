#!/bin/bash

# 设置工作目录为项目目录
cd ./

REPO_PATH="."

project="hello"

echo "---------------------------------------"
echo "Running project: $project"

# 启动 server
cd "$REPO_PATH" || exit
go run . > /dev/null 2>&1 &
server_pid=$!
cd - > /dev/null || exit

# 启动 client
cd "$REPO_PATH/client" || exit
go run main.go > /dev/null 2>&1 &
client_pid=$!
cd - > /dev/null || exit

# 当脚本退出时，停止 server 和 client
trap 'kill $server_pid $client_pid' EXIT

# 等待 server 和 client 结束
wait $server_pid $client_pid