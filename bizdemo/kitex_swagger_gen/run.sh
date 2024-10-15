#!/bin/bash

# 设置工作目录为项目目录
REPO_PATH="."
DOWNSTREAM_PATH="./downstream_server"

project="kitex_swagger_gen"

echo "---------------------------------------"
echo "Running project: $project"

# 启动当前目录下的服务
cd "$REPO_PATH/" || exit
echo "Starting main server..."
go run . > /dev/null 2>&1 &
main_server_pid=$!
cd - > /dev/null || exit

# 启动 downstream_server 服务
cd "$DOWNSTREAM_PATH/" || exit
echo "Starting downstream server..."
go run . > /dev/null 2>&1 &
downstream_server_pid=$!
cd - > /dev/null || exit

# 当脚本退出时，停止所有服务
trap 'kill $main_server_pid $downstream_server_pid' EXIT

# 等待所有服务结束
wait $main_server_pid $downstream_server_pid
