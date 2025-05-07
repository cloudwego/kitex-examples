#!/bin/bash

# 设置错误时退出
set -e

# 设置工作目录为项目目录
cd ./

REPO_PATH="."
project="generic-streaming-thrift"

echo "---------------------------------------"
echo "Running project: $project"

# 启动服务器
echo "Starting server..."
cd "$REPO_PATH/rpc" || exit
go run . &
SERVER_PID=$!
cd - > /dev/null || exit

# 等待服务器启动
echo "Waiting for server to start..."
sleep 3

# 检查服务器是否正在运行
if ! kill -0 $SERVER_PID 2>/dev/null; then
    echo "Error: Server failed to start"
    exit 1
fi

# 运行客户端并检查输出
echo "Running client..."
cd "$REPO_PATH/client" || exit
OUTPUT=$(go run main.go)
cd - > /dev/null || exit


# 清理
echo "Cleaning up..."
kill -9 $SERVER_PID $(lsof -t -i:8888) 2>/dev/null || true

# 设置脚本的退出状态
exit $status 