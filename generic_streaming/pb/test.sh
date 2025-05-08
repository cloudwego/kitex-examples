#!/bin/bash

# Exit on error
set -e

# Start server
echo "Starting server..."
go run . &
SERVER_PID=$!

# Wait for server to start
sleep 2

# Run client and check output
echo "Running client..."
cd client
OUTPUT=$(go run main.go)

# Cleanup
if kill -0 $SERVER_PID 2>/dev/null; then
    kill $SERVER_PID
fi
# If port 8888 is still in use, force release (use with caution)
if lsof -i :8888 > /dev/null 2>&1; then
    lsof -t -i:8888 | xargs kill -9
fi