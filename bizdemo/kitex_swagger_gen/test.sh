#!/bin/bash

# 设置工作目录为项目目录
cd ./

REPO_PATH="."
DOWNSTREAM_PATH="./downstream_server"

# 初始化状态变量
status=0
project="kitex_swagger_gen"

echo "---------------------------------------"
echo "Running project: $project"

# 启动主服务 (端口 8888)
cd "$REPO_PATH" || exit
go run . > /dev/null 2>&1 &
main_server_pid=$!
cd - > /dev/null || exit

# 启动下游服务 (端口 8889)
cd "$DOWNSTREAM_PATH" || exit
go run . > /dev/null 2>&1 &
downstream_server_pid=$!
cd - > /dev/null || exit

# 等待服务启动
sleep 1

# 检查主服务 (8888) 是否仍在运行
if kill -0 $main_server_pid; then
    echo "Main server (8888) run successfully."
else
    echo "Failed to start main server (8888)."
    status=1
fi

# 检查下游服务 (8889) 是否仍在运行
if kill -0 $downstream_server_pid; then
    echo "Downstream server (8889) run successfully."
else
    echo "Failed to start downstream server (8889)."
    status=1
fi

# 杀死两个服务
kill -9 $main_server_pid $(lsof -t -i:8888)
kill -9 $downstream_server_pid $(lsof -t -i:8889)

# 设置脚本的退出状态
exit $status
