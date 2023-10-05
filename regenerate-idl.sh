#!/bin/bash

kitex=${1:-kitex}

function gen_thrift() {
    frugal_tag='-thrift frugal_tag'
    module='-module github.com/cloudwego/kitex-examples'
    idl=echo.thrift

    $kitex $frugal_tag $module $idl
}

function gen_protobuf() {
    module='-module github.com/cloudwego/kitex-examples'
    idl=echo.proto
    $kitex $module $idl
}

gen_thrift
gen_protobuf
