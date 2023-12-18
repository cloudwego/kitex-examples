## 基于 protobuf IDL 创建项目

[English](./README.md) | 中文

### 创建一个新项目

1.  在当前目录下创建 protobuf idl 文件

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

2.  创建新项目

```
// GOPATH 下执行
kitex -service hello ./idl/hello.proto

// 不在 GOPATH 下执行
kitex -service hello -module kitex-examples/kitex/protobuf ./idl/hello.proto

// 整理 & 拉取依赖
go mod tidy
```

3.  修改handler，添加自己的逻辑

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

4.  编译项目

```
go build
```

5.  运行项目

运行项目：

```
./main
```

### 更新一个已有的项目

1.  如果你的 protobuf idl 有更新，例如：

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

2.  切换到一开始执行命令的目录，更新修改后的 protobuf idl

```
// GOPATH 下执行
kitex -service hello ./idl/hello.proto

// 不在 GOPATH 下执行
kitex -service hello -module kitex-examples/kitex/protobuf ./idl/hello.proto

// 整理 & 拉取依赖
go mod tidy
```

修改逻辑，编译运行同理