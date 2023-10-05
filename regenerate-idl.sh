#!/bin/bash

kitex=${1:-kitex}

function gen_thrift() {
    echo "Generate thrift to kitex_gen..."
    frugal_tag='-thrift frugal_tag'
    module='-module github.com/cloudwego/kitex-examples'
    idl=echo.thrift

    $kitex $frugal_tag $module $idl
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

gen_thrift
gen_thrift_slim
gen_protobuf
