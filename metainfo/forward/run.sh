#!/bin/bash

# 设置工作目录为项目目录
cd ./

REPO_PATH="."

project="metainfo_backward"

echo "---------------------------------------"
echo "Running project: $project"

# 启动 server2
cd "$REPO_PATH/server-2" || exit
go run main.go > /dev/null 2>&1 &
server_pid_1=$!
cd - > /dev/null || exit

# 启动 server1
cd "$REPO_PATH/server-1" || exit
go run main.go > /dev/null 2>&1 &
server_pid_2=$!
cd - > /dev/null || exit

# 启动 client
cd "$REPO_PATH/client" || exit
go run main.go > /dev/null 2>&1 &
client_pid=$!
cd - > /dev/null || exit

# 当脚本退出时，停止 server 和 client
trap 'kill $server_pid_2 $server_pid_1 $client_pid' EXIT

# 等待 server 和 client 结束
wait $server_pid_1 $server_pid_2 $client_pid