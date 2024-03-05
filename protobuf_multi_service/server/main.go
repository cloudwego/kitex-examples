// Copyright 2024 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package main

import (
	"context"
	"log"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"

	"github.com/cloudwego/kitex-examples/protobuf_multi_service/kitex_gen/multi/service"
	"github.com/cloudwego/kitex-examples/protobuf_multi_service/kitex_gen/multi/service/servicea"
	"github.com/cloudwego/kitex-examples/protobuf_multi_service/kitex_gen/multi/service/serviceb"
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

func main() {
	svr := server.NewServer()
	servicea.RegisterService(svr, &ServiceAImpl{})
	serviceb.RegisterService(svr, &ServiceBImpl{})
	if err := svr.Run(); err != nil {
		log.Println("server stopped with error:", err)
	} else {
		log.Println("server stopped")
	}
}
