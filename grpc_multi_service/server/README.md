## Create server based on grpc multi-service

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
kitex -service multiservice ./idl/demo.proto

//Do not execute under GOPATH
kitex -service multiservice -module grpc_multi_service ./idl/demo.proto

// Organize & pull dependencies
go mod tidy
```

3. Modify main.go
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