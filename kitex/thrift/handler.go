package main

import (
	"context"

	example "kitex-examples/kitex/thrift/kitex_gen/hello/example"
)

// HelloServiceImpl implements the last service interface defined in the IDL.
type HelloServiceImpl struct{}

// HelloMethod implements the HelloServiceImpl interface.
func (s *HelloServiceImpl) HelloMethod(ctx context.Context, request *example.HelloReq) (resp *example.HelloResp, err error) {
	resp = new(example.HelloResp)
	resp.RespBody = "hello " + request.Name
	return
}
