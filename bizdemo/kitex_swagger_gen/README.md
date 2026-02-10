# kitex_swagger_gen

## Introduction

An example using `Kitex` and `thrift-gen-rpc-swagger`.

- Defines `RPC` interfaces using `thrift` IDL
- Generates code using `Kitex`
- Utilizes `Gorm` and `MySQL`
- Uses the `thrift-gen-rpc-swagger` plugin to generate `Swagger` files and a `Swagger UI` service

- `/swagger` provides the `Swagger` files and `Swagger UI` server
- `/handler.go` contains the basic logic for updating, adding, deleting, querying users and examples using `metainfo`

## IDL

This example defines `RPC` interfaces using `thrift` IDL. The specific interface definitions can be found in [user.thrift](idl/user.thrift).

## Code Generation Tools

This example uses `Kitex` for code generation. For more information on how to use `Kitex`, refer to the [Kitex documentation](https://www.cloudwego.io/docs/kitex/tutorials/code-gen/).

The `Kitex` commands used can be found in the [Makefile](Makefile).

## Plugins

`thrift-gen-rpc-swagger` generates `Swagger` documentation and a `Swagger UI` service through code generation.

For more details, refer to [swagger-generate](https://github.com/hertz-contrib/swagger-generate).

## How to Run

### Running MySQL Docker

```bash
cd bizdemo/kitex_swagger_gen && docker compose up
```

### Running the Example

```bash
cd bizdemo/kitex_swagger_gen
go run .

cd bizdemo/kitex_swagger_gen/downstream_server
go run .
```
