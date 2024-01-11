## 基于 grpc multi-service 创建 server

[English](README.md) | 中文

当前这个功能仅支持 gRPC 传输协议

### 创建一个新项目

1.  在当前目录下创建 protobuf idl 文件

```
syntax = "proto3";

option go_package = "multi/service";

package multiservice;

service ServiceA {
  rpc ChatA (RequestA) returns (Reply) {}
}

message RequestA {
  string name = 1;
}

message Reply {
  string message = 1;
}
```

2.  创建新项目

```
// GOPATH 下执行
kitex -service multiservice ./idl/demo.proto

// 不在 GOPATH 下执行
kitex -service multiservice -module grpc_multi_service ./idl/demo.proto

// 整理 & 拉取依赖
go mod tidy
```

3.  修改handler，添加自己的逻辑

```
type ServiceAImpl struct{}

// ChatA implements the ServiceAImpl interface.
func (s *ServiceAImpl) ChatA(ctx context.Context, req *service.RequestA) (resp *service.Reply, err error) {
	klog.Info("ChatA called, req: ", req.Name)
	resp = new(service.Reply)
	resp.Message = "hello " + req.Name
	return
}
```

4. 在 idl 里添加新的 service

```
syntax = "proto3";

option go_package = "multi/service";

package multiservice;

service ServiceA {
  rpc ChatA (RequestA) returns (Reply) {}
}

service ServiceB {
  rpc ChatB (RequestB) returns (Reply) {}
}

message RequestA {
  string name = 1;
}

message RequestB {
  string name = 1;
}


message Reply {
  string message = 1;
}
```

5.  更新项目

```
// GOPATH 下执行
kitex -service multiservice ./idl/demo.proto

// 不在 GOPATH 下执行
kitex -service multiservice -module grpc_multi_service ./idl/demo.proto

// 整理 & 拉取依赖
go mod tidy
```

6. 修改 handler.go 和 main.go

手动添加到 handler.go
```
type ServiceBImpl struct{}
```

修改 ChatB 逻辑
```
// ChatB implements the ServiceBImpl interface.
func (s *ServiceBImpl) ChatB(ctx context.Context, req *service.RequestB) (resp *service.Reply, err error) {
	klog.Info("ChatB called, req: ", req.Name)
	resp = new(service.Reply)
	resp.Message = "hello " + req.Name
	return
}
```

修改 main.go
```
func main() {
	addr, _ := net.ResolveTCPAddr("tcp", "localhost:8888")

	svr := server.NewServer(server.WithServiceAddr(addr))
	err := svr.RegisterService(servicea.NewServiceInfo(), new(ServiceAImpl))
	if err != nil {
		log.Println(err.Error())
	}
	err = svr.RegisterService(serviceb.NewServiceInfo(), new(ServiceBImpl))
	if err != nil {
		log.Println(err.Error())
	}
	err = svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
```