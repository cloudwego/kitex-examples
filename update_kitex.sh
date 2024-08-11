#!/bin/bash

version=$1
exclude=$2

# find go.mod and update dependencies
find . -name go.mod -print0 | while IFS= read -r -d '' script; do
    # get absolute dir path
    pwd=$(pwd)
    script_dir=$(dirname "$script")
    if [ "$script_dir" == "$exclude" ]; then
        continue
    fi
    echo $script_dir
    cd "$script_dir" 
    go get -v "github.com/cloudwego/kitex@$version"
    go mod tidy
    cd "$pwd"
done

# find . -name test.sh -print0 | while IFS= read -r -d '' script; do
#     script_dir=$(dirname "$script")
#     if [ "$script_dir" == "$exclude" ]; then
#         continue
#     fi
#     chmod +x "$script"
#     (cd "$script_dir" && bash "./$(basename "$script")")
# done