# kitex_swagger_gen

## 介绍

一个使用 `Kitex` 和 `thrift-gen-rpc-swagger` 的示例。

- 使用 `thrift` IDL 定义 `RPC` 接口
- 使用 `kitex` 生成代码
- 使用 `Gorm` and `MySQL`
- 使用 `thrift-gen-rpc-swagger` 插件生成 `swagger` 文件和 `swagger ui` 服务

- `/swagger` 提供 `swagger` 文件和 `swagger ui` 服务器
- `/hander.go` 包含更新用户、添加用户、删除用户、查询用户的基础逻辑以及使用`metainfo`的示例

## IDL

该示例使用 `thrift` IDL 来定义 `RPC` 接口。具体的接口定义在 [user.thrift](idl/user.thrift) 中。

## 代码生成工具

该示例使用 `kitex` 来生成代码。`kitex` 的使用可以参考 [kitex](https://www.cloudwego.io/docs/kitex/tutorials/code-gen/)。

使用的 `kitex` 命令可以在 [Makefile](Makefile) 中找到。

## 插件

`thrift-gen-rpc-swagger` 通过代码生成的`swagger`文档和`swagger ui`服务。

详情可参考 [swagger-generate](https://github.com/hertz-contrib/swagger-generate)。

## 如何运行

### 运行 MySQL docker

```bash
cd bizdemo/kitex_swagger_gen && docker compose up
```

### 运行示例

```bash
cd bizdemo/kitex_swagger_gen
go run .

cd bizdemo/kitex_swagger_gen/downstream_server
go run .
```
