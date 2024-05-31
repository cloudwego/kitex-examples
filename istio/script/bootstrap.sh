#! /usr/bin/env bash
export PSM=${PSM:-a.b.c}
CURDIR=$(cd $(dirname $0); pwd)

if [ "X$1" != "X" ]; then
    RUNTIME_ROOT=$1
else
    RUNTIME_ROOT=${CURDIR}
fi

export KITEX_RUNTIME_ROOT=$RUNTIME_ROOT
export KITEX_CONF_FILE="kitex.yml"
export KITEX_CONF_DIR="$CURDIR/conf"
export KITEX_LOG_DIR="$RUNTIME_ROOT/log"

if [ ! -d "$KITEX_LOG_DIR/app" ]; then
    mkdir -p "$KITEX_LOG_DIR/app"
fi

if [ ! -d "$KITEX_LOG_DIR/rpc" ]; then
    mkdir -p "$KITEX_LOG_DIR/rpc"
fi

exec "$CURDIR/bin/a.b.c"