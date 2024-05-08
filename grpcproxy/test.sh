#!/bin/bash

# 设置工作目录为项目目录
cd ./

REPO_PATH="."

# 初始化状态变量
status=0
project="grpc_proxy"

echo "---------------------------------------"
echo "Running project: $project"

# 启动 server 和 client

cd "$REPO_PATH" || exit
go run main.go > /dev/null 2>&1 &
pid=$!
cd - > /dev/null || exit


# 等待 server 和 client 启动
sleep 1


# 检查 server 和 client 是否仍在运行
if kill -0 $pid ; then
    echo "Project run successfully: $project"
    echo "---------------------------------------"
else
    echo "Project failed to run: $project"
    echo "---------------------------------------"
    status=1
fi

# 杀死 server 和 client
kill -9  $pid  $(lsof -t -i:8888)


# 设置脚本的退出状态
exit $status