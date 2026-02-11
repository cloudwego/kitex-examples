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
	"io"
	"strconv"
	"sync"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/endpoint/cep"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/streaming"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/pkg/utils/kitexutil"
	"github.com/cloudwego/kitex/transport"

	test "github.com/cloudwego/kitex-examples/thrift_streaming/kitex_gen/echo"
	"github.com/cloudwego/kitex-examples/thrift_streaming/kitex_gen/echo/testservice"
)

var (
	streamClient = testservice.MustNewClient(
		"server_name_for_discovery",
		client.WithHostPorts("127.0.0.1:8888"),

		// client middleware
		client.WithMiddleware(func(e endpoint.Endpoint) endpoint.Endpoint {
			return func(ctx context.Context, req, resp interface{}) (err error) {
				method, _ := kitexutil.GetMethod(ctx)
				klog.Infof("[%s] streamclient middleware, req = %v", method, req)
				err = e(ctx, req, resp)
				klog.Infof("[%s] streamclient middleware, err = %v, resp = %v", method, err, resp)
				return err
			}
		}),

		// send middleware
		client.WithStreamOptions(
			client.WithStreamSendMiddleware(func(next cep.StreamSendEndpoint) cep.StreamSendEndpoint {
				return func(ctx context.Context, stream streaming.ClientStream, req interface{}) (err error) {
					method, _ := kitexutil.GetMethod(stream.Context())
					err = next(ctx, stream, req)
					klog.Infof("[%s] streamclient send middleware, err = %v, req = %v", method, err, req)
					return err
				}
			}),

			// recv middleware
			// NOTE: message (response from server) will NOT be available until `next` returns
			client.WithStreamRecvMiddleware(func(next cep.StreamRecvEndpoint) cep.StreamRecvEndpoint {
				return func(ctx context.Context, stream streaming.ClientStream, resp interface{}) (err error) {
					method, _ := kitexutil.GetMethod(stream.Context())
					err = next(ctx, stream, resp)
					klog.Infof("[%s] streamclient recv middleware, err = %v, resp = %v", method, err, resp)
					return err
				}
			}),
		),
	)

	pingPongClient = testservice.MustNewClient(
		"server_name_for_discovery",
		client.WithHostPorts("127.0.0.1:8888"),
		client.WithTransportProtocol(transport.TTHeader),
		client.WithMetaHandler(transmeta.ClientTTHeaderHandler),
		client.WithRPCTimeout(time.Second),
	)
)

func main() {
	// For PingPong APIs: use Client
	echoPingPong(pingPongClient)

	// For Streaming APIs: use StreamClient
	echo(streamClient)       // bidirectional streaming
	echoClient(streamClient) // client streaming
	echoServer(streamClient) // server streaming
	echoUnary(streamClient)  // unary, not recommended (performance issue); please just use PingPong api

	klog.Infof("main exit")
}

func echoPingPong(cli testservice.Client) {
	req := &test.Request{Message: "hello"}
	rsp, err := cli.EchoPingPong(context.Background(), req)
	if err != nil {
		klog.Warnf("echoPingPong: failed, err = " + err.Error())
		return
	}
	klog.Infof("echoPingPong: rsp = %v", rsp)
}

func echoUnary(cli testservice.Client) {
	ctx := context.Background()
	req := &test.Request{Message: "hello"}
	rsp, err := cli.EchoUnary(ctx, req)
	if err != nil {
		klog.Warnf("echoPingPong: failed, err = " + err.Error())
		return
	}
	klog.Infof("echoPingPong: rsp = %v", rsp)
}

func echo(cli testservice.Client) {
	ctx := context.Background()
	stream, err := cli.Echo(ctx)
	if err != nil {
		panic("echo: failed to call, err = " + err.Error())
	}

	wg := &sync.WaitGroup{}
	wg.Add(2)
	// Send
	go func() {
		defer wg.Done()
		defer stream.CloseSend(ctx) // Tell the server there'll be no more message from client
		for i := 0; i < 3; i++ {
			req := &test.Request{Message: "client, " + strconv.Itoa(i)}
			if err = stream.Send(ctx, req); err != nil {
				klog.Warnf("echo.send: failed, err = " + err.Error())
				break
			}
			klog.Infof("echo.sent: req = %+v", req)
		}
	}()

	// Recv
	go func() {
		defer wg.Done()
		for {
			resp, err := stream.Recv(ctx)
			// make sure you receive and io.EOF or other non-nil error
			// otherwise RPCFinish event will not be recorded
			if err == io.EOF {
				klog.Infof("echo.recv: stream is closed")
				break
			} else if err != nil {
				klog.Warnf("echo.recv: failed, err = " + err.Error())
				break
			}
			klog.Infof("echo.recv: resp = %+v", resp)
		}
	}()
	wg.Wait()
}

func echoClient(cli testservice.Client) {
	ctx := context.Background()
	stream, err := cli.EchoClient(context.Background())
	if err != nil {
		panic("failed to call Echo: " + err.Error())
	}
	for i := 0; i < 3; i++ {
		req := &test.Request{Message: "hello, " + strconv.Itoa(i)}
		err := stream.Send(ctx, req)
		if err != nil {
			panic("failed to send Echo: " + err.Error())
		}
		klog.Infof("sent: %+v", req)
	}

	// Recv
	resp, err := stream.CloseAndRecv(ctx)
	if err != nil {
		klog.Warnf("failed to recv Echo: " + err.Error())
	} else {
		klog.Infof("recv: %+v", resp)
	}
}

func echoServer(cli testservice.Client) {
	ctx := context.Background()
	req := &test.Request{Message: "hello"}
	stream, err := cli.EchoServer(context.Background(), req)
	if err != nil {
		panic("failed to call Echo: " + err.Error())
	}
	for {
		resp, err := stream.Recv(ctx)
		// make sure you receive and io.EOF or other non-nil error
		// otherwise RPCFinish event will not be recorded
		if err == io.EOF {
			klog.CtxInfof(stream.Context(), "stream is closed")
			break
		} else if err != nil {
			klog.CtxInfof(stream.Context(), "failed to recv Echo: "+err.Error())
			break
		}
		klog.CtxInfof(stream.Context(), "resp: %+v", resp)
	}
}
