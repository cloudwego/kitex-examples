#!/bin/bash

project="generic-streaming-thrift"
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

echo "---------------------------------------"
echo "Running project: $project"

# Start server
cd "$SCRIPT_DIR/rpc"
go run . &
pid=$!
cd "$SCRIPT_DIR"

# Wait for server to start
echo "Waiting for server to start..."
sleep 5  

# Check if server is still running
if ! ps -p $pid > /dev/null; then
    echo "Error: Server failed to start"
    exit 1
fi

# Run client
echo "Running client..."
cd "$SCRIPT_DIR/client"
go run main.go
cd "$SCRIPT_DIR"

# Stop server when script exits
trap 'kill $pid 2>/dev/null || true' EXIT

# Wait for server to finish
wait $pid 