#!/bin/bash

# 启动服务器
echo "Starting server..."
go run . &
SERVER_PID=$!

# 等待服务器启动
sleep 2

# 运行客户端
echo "Running client..."
cd client
go run main.go

# 清理
kill -9 $SERVER_PID $(lsof -t -i:8888) 2>/dev/null || true