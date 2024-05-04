#!/bin/bash

# 设置工作目录为项目目录
cd ./

REPO_PATH="."

# 初始化状态变量
status=0
project="basic"

echo "---------------------------------------"
echo "Running project: $project"

docker run -d -p 2379:2379 --name etcd quay.io/coreos/etcd:v3.5.0 /usr/local/bin/etcd --advertise-client-urls http://0.0.0.0:2379 --listen-client-urls http://0.0.0.0:2379

# 启动 item server

cd "$REPO_PATH/rpc/item" || exit
go run main.go > /dev/null 2>&1 &
item_server_pid=$!
cd - > /dev/null || exit


# 等待 item server 启动
sleep 1

# 启动 stock server

cd "$REPO_PATH/rpc/stock" || exit
go run main.go > /dev/null 2>&1 &
stock_server_pid=$!
cd - > /dev/null || exit


# 等待 stock server 启动
sleep 1

# 启动 client

cd "$REPO_PATH/api" || exit
go run main.go > /dev/null 2>&1 &
client_pid=$!
cd - > /dev/null || exit


# 等待 client 启动
sleep 1

# 检查 server 和 client 是否仍在运行
if kill -0 $item_server_pid && kill -0 $stock_server_pid && kill -0 $client_pid ; then
    echo "Project run successfully: $project"
    echo "---------------------------------------"
else
    echo "Project failed to run: $project"
    echo "---------------------------------------"
    status=1
fi

# 杀死 server 和 client
kill $item_server_pid $stock_server_pid $client_pid


# 设置脚本的退出状态
exit $status