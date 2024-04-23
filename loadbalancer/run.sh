#!/bin/bash

# 设置工作目录为项目目录
cd ./

REPO_PATH="."

project="loadbalancer"

echo "---------------------------------------"
echo "Running project: $project"

# 启动 server

cd "$REPO_PATH/server" || exit
go run main.go > /dev/null 2>&1 &
server_pid=$!
cd - > /dev/null || exit


# 等待 server 启动
sleep 1

# 启动 client_consistbalancer

cd "$REPO_PATH/client/consistbalancer" || exit
go run main.go > /dev/null 2>&1 &
client_pid_1=$!
cd - > /dev/null || exit


# 等待 client 启动
sleep 1

# 启动 client_interleavedweightedroundrobin

cd "$REPO_PATH/client/interleavedweightedroundrobin" || exit
go run main.go > /dev/null 2>&1 &
client_pid_2=$!
cd - > /dev/null || exit


# 等待 client 启动
sleep 1

# 启动 client_weightedrandom

cd "$REPO_PATH/client/weightedrandom" || exit
go run main.go > /dev/null 2>&1 &
client_pid_3=$!
cd - > /dev/null || exit


# 等待 client 启动
sleep 1

# 启动 client_weightedroundrobin

cd "$REPO_PATH/client/weightedroundrobin" || exit
go run main.go > /dev/null 2>&1 &
client_pid_4=$!
cd - > /dev/null || exit


# 等待 client 启动
sleep 1

# 当脚本退出时，停止 server
trap 'kill $server_pid' EXIT

# 等待 server 和 client 结束
wait $server_pid $client_pid_1 $client_pid_2 $client_pid_3 $client_pid_4