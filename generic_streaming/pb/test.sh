#!/bin/bash

# Set working directory to project directory
cd ./

REPO_PATH="."
status=0
project="generic-streaming-pb"

echo "---------------------------------------"
echo "Running project: $project"

# Start server
echo "Starting server..."
go run . &
SERVER_PID=$!

# Wait for server to start
sleep 1

# Check if server is still running
if kill -0 $SERVER_PID 2>/dev/null; then
    echo "Project run successfully: $project"
    echo "---------------------------------------"
else
    echo "Project failed to run: $project"
    echo "---------------------------------------"
    status=1
fi

# Run client
echo "Running client..."
cd client
OUTPUT=$(go run main.go)
echo "$OUTPUT"
cd - > /dev/null || exit

# Cleanup processes
kill -9 $SERVER_PID $(lsof -t -i:8888)

# Set script exit status
exit $status