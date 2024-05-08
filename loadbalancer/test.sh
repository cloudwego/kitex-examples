#!/bin/bash

# 设置工作目录为项目目录
cd ./

REPO_PATH="."

# 初始化状态变量
status=0
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

# 检查 server 和 client 是否仍在运行
if kill -0 $server_pid && kill -0 $client_pid_1 && kill -0 $client_pid_2 && kill -0 $client_pid_3 && kill -0 $client_pid_4; then
    echo "Project run successfully: $project"
    echo "---------------------------------------"
else
    echo "Project failed to run: $project"
    echo "---------------------------------------"
    status=1
fi

# 杀死 server 和 client
kill -9 $server_pid $client_pid_1 $client_pid_2 $client_pid_3 $client_pid_4 $(lsof -t -i:8888)


# 设置脚本的退出状态
exit $status