/*
 * Copyright 2024 CloudWeGo Authors
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

package main

import (
	"context"
	"net"
	"net/http"

	"github.com/cloudwego/kitex-examples/kitex_gen/api"
	"github.com/cloudwego/kitex-examples/kitex_gen/api/echo"
	"github.com/cloudwego/kitex-examples/kitex_gen/pbapi"
	echo1 "github.com/cloudwego/kitex-examples/kitex_gen/pbapi/echo"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	_ "github.com/cloudwego/kitex/pkg/remote/codec/protobuf/encoding/gzip"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/server"
	"github.com/cloudwego/kitex/transport"
	prometheus "github.com/kitex-contrib/monitor-prometheus"
	promlib "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var _ api.Echo = &EchoImpl{}

var (
	registry     = promlib.NewRegistry()
	clientTracer = prometheus.NewClientTracer("", "", prometheus.WithRegistry(registry), prometheus.WithDisableServer(true))
	serverTracer = prometheus.NewServerTracer("", "", prometheus.WithRegistry(registry), prometheus.WithDisableServer(true))
)

// EchoImpl implements the last service interface defined in the IDL.
type EchoImpl struct{}

// Echo implements the Echo interface.
func (s *EchoImpl) Echo(ctx context.Context, req *api.Request) (resp *api.Response, err error) {
	klog.Info(req.Message)
	return &api.Response{Message: req.Message}, nil
}

// PbEchoImpl implements the last service interface defined in the IDL.
type PbEchoImpl struct{}

// Echo implements the Echo interface.
func (p *PbEchoImpl) Echo(ctx context.Context, Req *pbapi.Request) (resp *pbapi.Response, err error) {
	resp = new(pbapi.Response)
	cli, _ := echo.NewClient("thrift-server",
		client.WithMetaHandler(transmeta.ClientTTHeaderHandler),
		client.WithHostPorts("localhost:8081"),
		client.WithTracer(clientTracer),
		client.WithTransportProtocol(transport.TTHeader),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: "thrift-client",
		}))
	resp1, err := cli.Echo(ctx, &api.Request{Message: Req.Message})
	if err != nil {
		klog.Error(err)
		return nil, err
	}

	resp.Message = "hello " + Req.Message + " resp1 " + resp1.Message
	return
}

func main() {
	http.Handle("/metrics", promhttp.HandlerFor(registry, promhttp.HandlerOpts{}))
	go http.ListenAndServe(":8092", nil) //nolint:errcheck

	addr1, _ := net.ResolveTCPAddr("tcp", ":8081")
	svr := echo.NewServer(new(EchoImpl),
		server.WithServiceAddr(addr1),
		server.WithTracer(serverTracer),
		server.WithMetaHandler(transmeta.ServerTTHeaderHandler),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: "thrift-server",
		}),
	)
	addr2, _ := net.ResolveTCPAddr("tcp", ":8082")
	svr2 := echo1.NewServer(
		new(PbEchoImpl), server.WithTracer(serverTracer),
		server.WithServiceAddr(addr2),
		server.WithMetaHandler(transmeta.ServerHTTP2Handler),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: "grpc-server",
		}),
	)
	go svr2.Run() //nolint:errcheck

	if err := svr.Run(); err != nil {
		klog.Error("server stopped with error:", err)
	}
	klog.Info("server stopped")
}
