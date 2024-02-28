package service

import (
	"context"
	"testing"

	example "github.com/kitex/hello/kitex_gen/hello/example"
)

func TestHelloMethod_Run(t *testing.T) {
	ctx := context.Background()
	s := NewHelloMethodService(ctx)
	// init req and assert value

	request := &example.HelloReq{}
	resp, err := s.Run(request)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test
}
