#!/bin/bash

# 设置工作目录为项目目录
cd ./

REPO_PATH="."

# 初始化状态变量
status=0
project="metainfo_forward"

echo "---------------------------------------"
echo "Running project: $project"

# 启动 server-2

cd "$REPO_PATH/server-2" || exit
go run main.go > /dev/null 2>&1 &
server_pid_2=$!
cd - > /dev/null || exit


# 等待 server-2 启动
sleep 1

# 启动 server-1

cd "$REPO_PATH/server-1" || exit
go run main.go > /dev/null 2>&1 &
server_pid_1=$!
cd - > /dev/null || exit


# 等待 server-1 启动
sleep 1

# 启动 client

cd "$REPO_PATH/client" || exit
go run main.go > /dev/null 2>&1 &
client_pid=$!
cd - > /dev/null || exit


# 等待 client 启动
sleep 1

# 检查 server 和 client 是否仍在运行
if kill -0 $server_pid_2 && kill -0 $server_pid_1 && kill -0 $client_pid; then
    echo "Project run successfully: $project"
    echo "---------------------------------------"
else
    echo "Project failed to run: $project"
    echo "---------------------------------------"
    status=1
fi

# 杀死 server 和 client
kill -9  $server_pid_1 $server_pid_2 $client_pid $(lsof -t -i:8888)


# 设置脚本的退出状态
exit $status