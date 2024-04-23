#!/bin/bash

# 设置工作目录为项目目录
cd ./

REPO_PATH="."

# 初始化状态变量
status=0
project="business_exception_TTHeader"

echo "---------------------------------------"
echo "Running project: $project"

# 检查端口是否被占用
if lsof -Pi :8888 -sTCP:LISTEN -t >/dev/null ; then
    kill -9 $(lsof -t -i:8888)
fi

# 启动 server

cd "$REPO_PATH/server" || exit
go run main.go > /dev/null 2>&1 &
server_pid=$!
cd - > /dev/null || exit


# 等待 server 启动
sleep 1

# 启动 client

cd "$REPO_PATH/client" || exit
go run main.go > /dev/null 2>&1
client_status=$?
cd - > /dev/null || exit


# 等待 client 启动
sleep 1

# 检查 server 是否仍在运行
if kill -0 $server_pid && [ $client_status -ne 0 ]; then
    echo "Project run successfully: $project"
    echo "---------------------------------------"
else
    echo "Project failed to run: $project"
    echo "---------------------------------------"
    status=1
fi

# 杀死 server
kill $server_pid


# 设置脚本的退出状态
exit $status