#!/bin/bash

# 设置工作目录为项目目录
cd ./

REPO_PATH="."

# 初始化状态变量
project="metainfo_backward"

echo "---------------------------------------"
echo "Running project: $project"


# 启动 server

cd "$REPO_PATH/thrift" || exit
go run main.go > /dev/null 2>&1 &
cd - > /dev/null || exit

