#!/bin/bash

# Start server
echo "Starting server..."
go run . &
SERVER_PID=$!

# Wait for server to start
sleep 2

# Run client
echo "Running client..."
cd client
go run main.go

# Cleanup
kill -9 $SERVER_PID $(lsof -t -i:8888) 2>/dev/null || true