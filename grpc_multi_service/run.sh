#!/bin/bash

# 设置工作目录为项目目录
cd ./

REPO_PATH="."
project="grpc_multi_service"

echo "---------------------------------------"
echo "Running project: $project"

# Start server
echo "Starting server..."
cd "$REPO_PATH/server" || exit
go run . > /dev/null 2>&1 &
server_pid=$!
cd - > /dev/null || exit

# Wait for server to start
sleep 1

# Check if server is still running
if ! kill -0 $server_pid 2>/dev/null; then
    echo "Error: Server failed to start"
    exit 1
fi
echo "Server started successfully (PID: $server_pid)"

# Run client
echo "Running client..."
cd "$REPO_PATH/client" || exit
go run main.go
client_status=$?
cd - > /dev/null || exit

# Cleanup processes
kill -9 $server_pid $(lsof -t -i:8888)

# Set script exit status
exit $client_status