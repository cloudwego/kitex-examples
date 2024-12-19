#!/bin/bash

# for kitex version higher than v0.12.0, thriftgo is no need to install manually.
go install github.com/cloudwego/kitex/tool/cmd/kitex@latest

kitex=${1:-kitex}

function gen_thrift() {
    echo "Generate thrift to kitex_gen..."
    frugal_tag='-thrift frugal_tag'
    module='-module github.com/cloudwego/kitex-examples'
    idl=echo.thrift

    $kitex $frugal_tag $module -invoker $idl
}

function gen_thrift_slim() {
    echo "Generate thrift(slim) to kitex_gen/slim..."
    frugal_tag='-thrift frugal_tag'
    slim='-thrift template=slim'
    module='-module github.com/cloudwego/kitex-examples'
    gen_path='-gen-path kitex_gen/slim'
    idl=echo.thrift

    $kitex $frugal_tag $slim $module $gen_path $idl
}

function gen_protobuf() {
    echo "Generate protobuf to kitex_gen..."
    module='-module github.com/cloudwego/kitex-examples'
    idl=echo.proto
    $kitex $module $idl
}

function regenerate_submod() {
    original_dir=$(pwd)
    local path=$1
    local module=$2
    local thrift_file=$3
    cd $path

    echo "Executing kitex command in $path with module: $module and thrift file: $thrift_file"

    # Execute the kitex command
    kitex -module $module $thrift_file
    go get github.com/cloudwego/kitex@latest
    go mod tidy

    cd $original_dir
}

gen_thrift
gen_thrift_slim
gen_protobuf

regenerate_submod "basic/example_shop" "example_shop" "idl/item.thrift"
regenerate_submod "basic/example_shop" "example_shop" "idl/stock.thrift"
regenerate_submod "bizdemo/easy_note" "github.com/cloudwego/kitex-examples/bizdemo/easy_note" "idl/note.thrift"
regenerate_submod "bizdemo/kitex_ent" "github.com/cloudwego/kitex-examples/bizdemo/kitex_ent" "idl/user.thrift"
regenerate_submod "bizdemo/kitex_gorm" "github.com/cloudwego/kitex-examples/bizdemo/kitex_gorm" "idl/user.thrift"
regenerate_submod "bizdemo/kitex_gorm_gen" "github.com/cloudwego/kitex-examples/bizdemo/kitex_gorm_gen" "idl/user.thrift"
regenerate_submod "bizdemo/kitex_swagger_gen" "github.com/cloudwego/kitex-examples/bizdemo/kitex_swagger_gen" "idl/user.thrift"
regenerate_submod "bizdemo/kitex_zorm" "github.com/cloudwego/kitex-examples/bizdemo/kitex_zorm" "idl/user.thrift"
regenerate_submod "generic" " github.com/cloudwego/kitex-examples" "http/http.thrift"
regenerate_submod "hello" " github.com/cloudwego/kitex-examples" "hello.thrift"
regenerate_submod "thrift_multi_service" " github.com/cloudwego/kitex-examples" "idl/demo.thrift"

