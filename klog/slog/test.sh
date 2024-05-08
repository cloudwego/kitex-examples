#!/bin/bash

# 设置工作目录为项目目录
cd ./

REPO_PATH="."

# 初始化状态变量
status=0
project="klog_slog"

echo "---------------------------------------"
echo "Running project: $project"

# 启动 server

cd "$REPO_PATH" || exit
go run . > /dev/null 2>&1 &
server_pid=$!
cd - > /dev/null || exit


# 等待 server  启动
sleep 1


# 检查 server 是否仍在运行
if kill -0 $server_pid ; then
    echo "Project run successfully: $project"
    echo "---------------------------------------"
else
    echo "Project failed to run: $project"
    echo "---------------------------------------"
    status=1
fi

# 杀死 server
kill -9  $server_pid  $(lsof -t -i:8888)


# 设置脚本的退出状态
exit $status