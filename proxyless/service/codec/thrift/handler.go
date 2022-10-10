package main

import (
	"context"

	proxyless "github.com/cloudwego/kitex-proxyless-test/service/codec/thrift/kitex_gen/proxyless"
)

// GreetServiceImpl implements the last service interface defined in the IDL.
type GreetServiceImpl struct{}

// SayHello1 implements the GreetServiceImpl interface.
func (s *GreetServiceImpl) SayHello1(ctx context.Context, request *proxyless.HelloRequest) (resp *proxyless.HelloResponse, err error) {
	// TODO: Your code here...
	return
}

// SayHello2 implements the GreetServiceImpl interface.
func (s *GreetServiceImpl) SayHello2(ctx context.Context, request *proxyless.HelloRequest) (resp *proxyless.HelloResponse, err error) {
	// TODO: Your code here...
	return
}
