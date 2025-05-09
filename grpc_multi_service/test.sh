#!/bin/bash

# Set working directory to project directory
cd ./

REPO_PATH="."
status=0
project="grpc_multi_service"

echo "---------------------------------------"
echo "Running project: $project"

# Start server
echo "Starting server..."
cd "$REPO_PATH/server" || exit
go run . &
SERVER_PID=$!
cd - > /dev/null || exit

# Wait for server to start
sleep 1

# Check if server is still running
if kill -0 $SERVER_PID 2>/dev/null; then
    echo "Server started successfully"
else
    echo "Server failed to start"
    echo "---------------------------------------"
    exit 1
fi

# Run client
echo "Running client..."
cd "$REPO_PATH/client" || exit
OUTPUT=$(go run main.go)
echo "$OUTPUT"
client_status=$?
cd - > /dev/null || exit

# Check execution status
if [ $client_status -eq 0 ]; then
    echo "Project run successfully: $project"
    echo "---------------------------------------"
else
    echo "Project failed to run: $project"
    echo "---------------------------------------"
    status=1
fi

# Cleanup processes
kill -9 $SERVER_PID $(lsof -t -i:8888) 2>/dev/null || true

# Set script exit status
exit $status