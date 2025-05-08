#!/bin/bash

# Exit on error
set -e

# Set working directory to project directory
cd ./

REPO_PATH="."
project="generic-streaming-thrift"

echo "---------------------------------------"
echo "Running project: $project"

# Start server
echo "Starting server..."
cd "$REPO_PATH/rpc" || exit
go run . &
SERVER_PID=$!
cd - > /dev/null || exit

# Wait for server to start
echo "Waiting for server to start..."
sleep 3

# Check if server is running
if ! kill -0 $SERVER_PID 2>/dev/null; then
    echo "Error: Server failed to start"
    exit 1
fi

# Run client and check output
echo "Running client..."
cd "$REPO_PATH/client" || exit
OUTPUT=$(go run main.go)
cd - > /dev/null || exit

# Cleanup
echo "Cleaning up..."
if kill -0 $SERVER_PID 2>/dev/null; then
    kill $SERVER_PID
fi

if lsof -i :8888 > /dev/null 2>&1; then
    lsof -t -i:8888 | xargs kill -9
fi

# Set script exit status
exit $status 