## Create client based on grpc multi-service

English | [中文](README_CN.md)

Currently, this feature only supports the gRPC transport protocol

### Create a new project

1. Create protobuf idl file in the current directory

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

2. Create a new project

```
// Execute under GOPATH
kitex ./idl/demo.proto

//Do not execute under GOPATH
kitex -module grpc_multi_service ./idl/demo.proto

// Organize & pull dependencies
go mod tidy
```

3. Modify main.go
```
func main() {
	clienta := servicea.MustNewClient("servicea", client.WithTransportProtocol(transport.GRPC), client.WithHostPorts("127.0.0.1:8888"))
	clientb := serviceb.MustNewClient("serviceb", client.WithTransportProtocol(transport.GRPC), client.WithHostPorts("127.0.0.1:8888"))
        // The newly created client must add client.WithTransportProtocol(transport.GRPC)
        // Otherwise, there will be a problem that method cannot be found
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