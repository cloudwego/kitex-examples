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

	"github.com/cloudwego/kitex-examples/thrift_multi_service/kitex_gen/api"
	"github.com/cloudwego/kitex-examples/thrift_multi_service/kitex_gen/api/servicea"
	"github.com/cloudwego/kitex-examples/thrift_multi_service/kitex_gen/api/serviceb"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
)

var _ api.ServiceA = &ServiceAImpl{}
var _ api.ServiceB = &ServiceBImpl{}

// ServiceAImpl implements the last servicea interface defined in the IDL.
type ServiceAImpl struct{}

// EchoA implements the EchoA interface.
func (s *ServiceAImpl) EchoA(ctx context.Context, req *api.Request) (resp *api.Response, err error) {
	klog.Info("echoA called")
	return &api.Response{Message: req.Message}, nil
}

// ServiceBImpl implements the last serviceb interface defined in the IDL.
type ServiceBImpl struct{}

// EchoB implements the EchoB interface.
func (s *ServiceBImpl) EchoB(ctx context.Context, req *api.Request) (resp *api.Response, err error) {
	klog.Info("echoB called")
	return &api.Response{Message: req.Message}, nil
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
