#!/bin/bash

# 设置工作目录为项目目录
cd ./

REPO_PATH="."

# 初始化状态变量
status=0
project="easy_note"

echo "---------------------------------------"
echo "Running project: $project"

cd "$REPO_PATH" || exit
docker-compose up -d
cd - > /dev/null || exit

# 启动 rpc note server

cd "$REPO_PATH/cmd/note" || exit
go run . > /dev/null 2>&1 &
note_server_pid=$!
cd - > /dev/null || exit


# 等待 server 启动
sleep 1

# 启动 rpc stock server

cd "$REPO_PATH/cmd/user" || exit
go run . > /dev/null 2>&1 &
user_server_pid=$!
cd - > /dev/null || exit


# 等待 server 启动
sleep 1

# 启动 hertx_server

cd "$REPO_PATH/cmd/api" || exit
go run main.go > /dev/null 2>&1 &
server_pid=$!
cd - > /dev/null || exit


# 等待 client 启动
sleep 1

# 检查 server 和 client 是否仍在运行
if kill -0 $note_server_pid && kill -0 $user_server_pid && kill -0 $server_pid; then
    echo "Project run successfully: $project"
    echo "---------------------------------------"
else
    echo "Project failed to run: $project"
    echo "---------------------------------------"
    status=1
fi

# 杀死 server 和 client
kill -9 $user_server_pid $note_server_pid $server_pid $(lsof -t -i:8888)

# 停止并删除所有容器
cd "$REPO_PATH" || exit
docker-compose down
cd - > /dev/null || exit


# 设置脚本的退出状态
exit $status