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
	"net"
	"time"

	dproto "github.com/cloudwego/dynamicgo/proto"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/genericclient"
	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	"github.com/cloudwego/kitex/transport"

	"jsonpb-demo/kitex_gen/api"
	"jsonpb-demo/kitex_gen/api/echo"
)

const serverHostPort = "127.0.0.1:9999"

func main() {
	go runGenericServer()
	time.Sleep(time.Second) // wait for server to start
	runClient()
}

func runClient() {
	normalCall(context.Background())
	genericJsonCall(context.Background())
}

func normalCall(ctx context.Context) {
	cli := echo.MustNewClient("server_name_for_discovery", client.WithHostPorts(serverHostPort))
	rsp, err := cli.EchoPB(ctx, &api.Request{
		Message: "hello",
	})
	klog.CtxInfof(ctx, "rsp: %v, err: %v", rsp, err)
}

func genericJsonCall(ctx context.Context) {
	jReq := `{"message": "hello"}`
	cli := initJSONGenericClient()
	jRsp, err := cli.GenericCall(ctx, "EchoPB", jReq)
	klog.CtxInfof(ctx, "genericJsonCall: jRsp(%T) = %s, err = %v", jRsp, jRsp, err)
}

func initJSONGenericClient() genericclient.Client {
	var err error

	path := "./idl/api.proto"

	// initialise DynamicGo proto.ServiceDescriptor
	dOpts := dproto.Options{}
	p, err := generic.NewPbFileProviderWithDynamicGo(path, context.Background(), dOpts)
	if err != nil {
		panic(err)
	}
	g, err := generic.JSONPbGeneric(p)
	if err != nil {
		panic(err)
	}

	var opts []client.Option
	opts = append(opts, client.WithHostPorts(serverHostPort))
	opts = append(opts, client.WithTransportProtocol(transport.TTHeader))

	cli, err := genericclient.NewClient("server_name_for_discovery", g, opts...)
	if err != nil {
		panic(err)
	}
	return cli
}

func WithServiceAddr(hostPort string) server.Option {
	addr, _ := net.ResolveTCPAddr("tcp", hostPort)
	return server.WithServiceAddr(addr)
}

type EchoImpl struct{}

// EchoPB implements the EchoImpl interface.
func (s *EchoImpl) EchoPB(ctx context.Context, req *api.Request) (resp *api.Response, err error) {
	resp = &api.Response{Message: req.Message}
	return
}

func runGenericServer() {
	var opts []server.Option
	opts = append(opts, WithServiceAddr(serverHostPort))

	svr := echo.NewServer(new(EchoImpl), opts...)
	if err := svr.Run(); err != nil {
		klog.Infof(err.Error())
	}
}
