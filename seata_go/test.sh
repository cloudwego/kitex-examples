#!/bin/bash

cd ./
REPO_PATH="."

# initialize the status variable
status=0
project="seata_go"

echo "---------------------------------------"
echo "Running project: $project"

cd "$REPO_PATH/dockercompose" || exit
docker-compose up -d
cd - > /dev/null || exit

sleep 10

# start servers
go run $REPO_PATH/service/account > /dev/null 2>&1 &
account_server_pid=$!

go run $REPO_PATH/service/storage > /dev/null 2>&1 &
storage_server_pid=$!

go run $REPO_PATH/service/order > /dev/null 2>&1 &
order_server_pid=$!

sleep 3

# start client
go run main.go > /dev/null 2>&1 &
client_pid=$!

sleep 1

# check if server and client are still running
if kill -0 $account_server_pid && kill -0 $storage_server_pid && kill -0 $order_server_pid && kill -0 $client_pid; then
    echo "Project run successfully: $project"
    echo "---------------------------------------"
else
    echo "Project failed to run: $project"
    echo "---------------------------------------"
    status=1
fi

# kill all servers and client
kill -9  $account_server_pid $storage_server_pid $order_server_pid  $client_pid

cd "$REPO_PATH/dockercompose" || exit
docker-compose down
cd - > /dev/null || exit

exit $status
