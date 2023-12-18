## Create a project based on protobuf IDL

English | [中文](./README_CN.md)

### Create a new project

1. Create protobuf idl file in the current directory

```
syntax = "proto3";

package hello;

option go_package = "hello";

message HelloReq {
  string Name = 1;
}

message HelloResp {
  string RespBody = 1;
}

service HelloService {
  rpc Hello(HelloReq) returns(HelloResp);
}
```

2. Create a new project

```
// GOPATH 下执行
kitex -service hello ./idl/hello.proto

// 不在 GOPATH 下执行
kitex -service hello -module kitex-examples/kitex/protobuf ./idl/hello.proto

// 整理 & 拉取依赖
go mod tidy
```

3. Modify the handler and add your own logic

```
// HelloServiceImpl implements the last service interface defined in the IDL.
type HelloServiceImpl struct{}

// Hello implements the HelloServiceImpl interface.
func (s *HelloServiceImpl) Hello(ctx context.Context, req *hello.HelloReq) (resp *hello.HelloResp, err error) {
	resp = new(hello.HelloResp)
	resp.RespBody = "hello " + req.Name
	return
}
```

4. Compile the project

```
go build
```

5. Run the project

Run the project:

```
./main
```

### Update an existing project

1. If your protobuf idl has been updated, for example:

```
syntax = "proto3";

package hello;

option go_package = "hello";

message HelloReq {
  string Name = 1;
}

message HelloResp {
  string RespBody = 1;
}

message ByeReq {
  string Name = 1;
}

message ByeResp {
  string RespBody = 1;
}

service HelloService {
  rpc Hello(HelloReq) returns(HelloResp);
  rpc Bye(ByeReq) returns(ByeResp);
}
```

2. Switch to the directory where the command was initially executed and update the modified protobuf idl

```
// GOPATH 下执行
kitex -service hello ./idl/hello.proto

// 不在 GOPATH 下执行
kitex -service hello -module kitex-examples/kitex/protobuf ./idl/hello.proto

// 整理 & 拉取依赖
go mod tidy
```

Modify the logic, compile and run in the same way