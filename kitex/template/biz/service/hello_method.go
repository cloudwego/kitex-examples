package service

import (
	"context"

	example "github.com/kitex/hello/kitex_gen/hello/example"
)

type HelloMethodService struct {
	ctx context.Context
} // NewHelloMethodService new HelloMethodService
func NewHelloMethodService(ctx context.Context) *HelloMethodService {
	return &HelloMethodService{ctx: ctx}
}

// Run create note info
func (s *HelloMethodService) Run(request *example.HelloReq) (resp *example.HelloResp, err error) {
	// Finish your business logic.

	return
}
