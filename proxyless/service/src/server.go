/*
 * Copyright 2022 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package src

import (
	"context"
	"fmt"
	"os"

	"github.com/bytedance/gopkg/cloud/metainfo"
	"github.com/cloudwego/kitex/server"

	"github.com/cloudwego/kitex-examples/proxyless/service/codec/thrift/kitex_gen/proxyless"
	"github.com/cloudwego/kitex-examples/proxyless/service/codec/thrift/kitex_gen/proxyless/greetservice"
)

// GreetServiceImpl implements the last service interface defined in the IDL.
type GreetServiceImpl struct{}

// SayHello1 implements the GreetServiceImpl interface.
func (s *GreetServiceImpl) SayHello1(ctx context.Context, request *proxyless.HelloRequest) (resp *proxyless.HelloResponse, err error) {
	// TODO: Your code here...
	resp = proxyless.NewHelloResponse()
	fmt.Println("SayHello1 Called")
	resp.SetMessage("Hello1!")
	if podName, ok := os.LookupEnv(PodNameKey); ok {
		metainfo.SendBackwardValue(ctx, PodNameKey, podName)
	}
	return
}

// SayHello2 implements the GreetServiceImpl interface.
func (s *GreetServiceImpl) SayHello2(ctx context.Context, request *proxyless.HelloRequest) (resp *proxyless.HelloResponse, err error) {
	// TODO: Your code here...
	resp = proxyless.NewHelloResponse()
	fmt.Println("SayHello2 Called")
	resp.SetMessage("Hello2!")
	if podName, ok := os.LookupEnv(PodNameKey); ok {
		metainfo.SendBackwardValue(ctx, PodNameKey, podName)
	}
	return
}

type ProxylessServer struct {
	svr server.Server
}

func NewProxylessServer() TestService {
	return &ProxylessServer{svr: greetservice.NewServer(&GreetServiceImpl{})}
}

func (s *ProxylessServer) Run() error {
	// :8888 is the default port for the server.
	err := s.svr.Run()
	return err
}
