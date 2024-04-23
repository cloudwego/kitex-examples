#!/bin/bash

# 设置工作目录为项目目录
cd ./

REPO_PATH="."

# 初始化状态变量
project="metainfo_backward"

echo "---------------------------------------"
echo "Running project: $project"

# 检查端口是否被占用
if lsof -Pi :8888 -sTCP:LISTEN -t >/dev/null ; then
    kill -9 $(lsof -t -i:8888)
fi

# 启动 server

cd "$REPO_PATH/thrift" || exit
go run main.go > /dev/null 2>&1 &
cd - > /dev/null || exit

