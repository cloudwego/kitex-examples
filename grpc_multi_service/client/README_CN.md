## 基于 grpc multi-service 创建 client

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
kitex ./idl/demo.proto

// 不在 GOPATH 下执行
kitex -module grpc_multi_service ./idl/demo.proto

// 整理 & 拉取依赖
go mod tidy
```

3. 修改 main.go
```
func main() {
	clienta := servicea.MustNewClient("servicea", client.WithTransportProtocol(transport.GRPC), client.WithHostPorts("127.0.0.1:8888"))
	clientb := serviceb.MustNewClient("serviceb", client.WithTransportProtocol(transport.GRPC), client.WithHostPorts("127.0.0.1:8888"))
        // 新建的客户端必须添加 client.WithTransportProtocol(transport.GRPC)
        // 否则会出现找不到 method 的问题 
	resa, err := clienta.ChatA(context.Background(), &service.RequestA{Name: "hello,a"})
	if err != nil {
		klog.Error(err)
		return
	}

	resb, err := clientb.ChatB(context.Background(), &service.RequestB{Name: "hello,b"})
	if err != nil {
		klog.Error(err)
		return
	}

	klog.Info("resa: ", resa.Message)
	klog.Info("resb: ", resb.Message)
}

```