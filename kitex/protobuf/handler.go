package main

import (
	"context"

	hello "kitex-examples/kitex/protobuf/kitex_gen/hello"
)

// HelloServiceImpl implements the last service interface defined in the IDL.
type HelloServiceImpl struct{}

// Hello implements the HelloServiceImpl interface.
func (s *HelloServiceImpl) Hello(ctx context.Context, req *hello.HelloReq) (resp *hello.HelloResp, err error) {
	resp = new(hello.HelloResp)
	resp.RespBody = "hello " + req.Name
	return
}
