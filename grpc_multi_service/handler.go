package main

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	service "grpc_multi_service/kitex_gen/multi/service"
)

type ServiceAImpl struct{}

type ServiceBImpl struct{}

// ChatA implements the ServiceAImpl interface.
func (s *ServiceAImpl) ChatA(ctx context.Context, req *service.RequestA) (resp *service.Reply, err error) {
	klog.Info("ChatA called, req: ", req.Name)
	resp = new(service.Reply)
	resp.Message = "hello " + req.Name
	return
}

// ChatB implements the ServiceBImpl interface.
func (s *ServiceBImpl) ChatB(ctx context.Context, req *service.RequestB) (resp *service.Reply, err error) {
	klog.Info("ChatB called, req: ", req.Name)
	resp = new(service.Reply)
	resp.Message = "hello " + req.Name
	return
}
