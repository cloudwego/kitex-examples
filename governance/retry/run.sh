#!/bin/bash

# 设置工作目录为项目目录
cd ./

REPO_PATH="."

project="governance-retry"

echo "---------------------------------------"
echo "Running project: $project"

# 启动 server
cd "$REPO_PATH/server" || exit
go run main.go > /dev/null 2>&1 &
server_pid=$!
cd - > /dev/null || exit

# 启动 client
cd "$REPO_PATH/client" || exit
go run main.go > /dev/null 2>&1 &
cd - > /dev/null || exit

# 当脚本退出时，停止 server
trap 'kill $server_pid' EXIT

# 等待 server 结束
wait $server_pid