## 基于 thrift IDL 创建项目

[English](./README.md) | 中文

### 创建一个新项目

1.  在当前目录下创建 thrift idl 文件

```
// idl/hello.thrift
namespace go hello.example

struct HelloReq {
    1: string Name;
}

struct HelloResp {
    1: string RespBody;
}


service HelloService {
    HelloResp HelloMethod(1: HelloReq request);
}
```

2.  创建新项目

```
// GOPATH 下执行
kitex -service hello ./idl/hello.thrift

// 不在 GOPATH 下执行
kitex -service hello -module kitex-examples/kitex/thrift ./idl/hello.thrift

// 整理 & 拉取依赖
go mod tidy
```

3.  修改handler，添加自己的逻辑

```
// HelloServiceImpl implements the last service interface defined in the IDL.
type HelloServiceImpl struct{}

// HelloMethod implements the HelloServiceImpl interface.
func (s *HelloServiceImpl) HelloMethod(ctx context.Context, request *example.HelloReq) (resp *example.HelloResp, err error) {
	resp = new(example.HelloResp)
	resp.RespBody = "hello " + request.Name
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

1.  如果你的 thrift idl 有更新，例如：

```
// idl/hello.thrift
namespace go hello.example

struct HelloReq {
    1: string Name;
}

struct HelloResp {
    1: string RespBody;
}

struct OtherReq {
    1: string Other;
}

struct OtherResp {
    1: string Resp;
}


service HelloService {
    HelloResp HelloMethod(1: HelloReq request);
    OtherResp OtherMethod(1: OtherReq request);
}

service NewService {
    HelloResp NewMethod(1: HelloReq request);
}
```

2.  切换到一开始执行命令的目录，更新修改后的 thrift idl

```
// GOPATH 下执行
kitex -service hello ./idl/hello.thrift

// 不在 GOPATH 下执行
kitex -service hello -module kitex-examples/kitex/thrift ./idl/hello.thrift

// 整理 & 拉取依赖
go mod tidy
```

修改逻辑，编译运行同理