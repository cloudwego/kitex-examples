## Create a project based on thrift IDL

English | [中文](./README_CN.md)

### Create a new project

1. Create thrift idl file in the current directory

```
//idl/hello.thrift
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

2. Create a new project

```
// Execute under GOPATH
kitex -service hello ./idl/hello.thrift

// Execute not under GOPATH
kitex -service hello -module kitex-examples/kitex/thrift ./idl/hello.thrift

// Organize & pull dependencies
go mod tidy
```

3. Modify the handler and add your own logic

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

1. If your thrift idl has been updated, for example:

```
//idl/hello.thrift
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

2. Switch to the directory where the command was initially executed and update the modified thrift idl

```
// Execute under GOPATH
kitex -service hello ./idl/hello.thrift

// Execute not under GOPATH
kitex -service hello -module kitex-examples/kitex/thrift ./idl/hello.thrift

// Organize & pull dependencies
go mod tidy
```

Modify the logic, compile and run in the same way