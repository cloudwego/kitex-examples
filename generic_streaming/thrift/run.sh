#!/bin/bash

project="generic-streaming-thrift"
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

echo "---------------------------------------"
echo "Running project: $project"

# 启动 server
cd "$SCRIPT_DIR/rpc"
go run . &
pid=$!
cd "$SCRIPT_DIR"

# 等待服务器启动
echo "Waiting for server to start..."
sleep 5  

# 检查服务器是否还在运行
if ! ps -p $pid > /dev/null; then
    echo "Error: Server failed to start"
    exit 1
fi

# 运行客户端
echo "Running client..."
cd "$SCRIPT_DIR/client"
go run main.go
cd "$SCRIPT_DIR"

# 当脚本退出时，停止 server
trap 'kill $pid 2>/dev/null || true' EXIT

# 等待 server 结束
wait $pid 