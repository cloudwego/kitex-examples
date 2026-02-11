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
	"fmt"
	"io"
	"strconv"
	"sync"

	"github.com/cloudwego/kitex/pkg/klog"

	"github.com/cloudwego/kitex-examples/thrift_streaming/kitex_gen/echo"
)

// TestServiceImpl implements the last service interface defined in the IDL.
type TestServiceImpl struct{}

// Echo is bidirectional streaming
func (s *TestServiceImpl) Echo(ctx context.Context, stream echo.TestService_EchoServer) (err error) {
	// no need to call `stream.Close()` manually
	wg := sync.WaitGroup{}
	wg.Add(2)
	// Recv
	go func() {
		defer func() {
			if p := recover(); p != nil {
				err = fmt.Errorf("panic: %v", p)
			}
			wg.Done()
		}()
		for {
			msg, recvErr := stream.Recv(ctx)
			// make sure you receive and io.EOF or other non-nil error
			// otherwise RPCFinish event will not be recorded
			if recvErr == io.EOF {
				klog.Infof("Echo: stream is closed")
				return
			} else if recvErr != nil {
				err = recvErr
				return
			}
			klog.Infof("Echo: received message = %s", msg)
		}
	}()
	// Send
	go func() {
		defer func() {
			if p := recover(); p != nil {
				err = fmt.Errorf("panic: %v", p)
			}
			wg.Done()
		}()
		for i := 0; i < 3; i++ {
			msg := &echo.Response{Message: "server, " + strconv.Itoa(i)}
			if sendErr := stream.Send(ctx, msg); sendErr != nil {
				err = sendErr
				return
			}
			klog.Infof("Echo: sent message = %s", msg)
		}
	}()
	wg.Wait()
	return
}

// EchoClient is client streaming
func (s *TestServiceImpl) EchoClient(ctx context.Context, stream echo.TestService_EchoClientServer) (err error) {
	for i := 0; i < 3; i++ {
		msg, err := stream.Recv(ctx)
		if err != nil {
			return err
		}
		klog.Infof("EchoClient: recv message = %s", msg)
	}
	return stream.SendAndClose(ctx, &echo.Response{Message: "echoClient"})
}

// EchoServer is server streaming
func (s *TestServiceImpl) EchoServer(ctx context.Context, req *echo.Request, stream echo.TestService_EchoServerServer) (err error) {
	klog.Infof("EchoServer called, req = %+v", req)
	for i := 0; i < 3; i++ {
		msg := &echo.Response{Message: "server, " + strconv.Itoa(i)}
		if err = stream.Send(ctx, msg); err != nil {
			return
		}
		klog.Infof("EchoServer: sent message = %s", msg)
	}
	return
}

// EchoPingPong implements the TestServiceImpl interface.
// You can define both Streaming APIs and PingPong APIs in the same service.
func (s *TestServiceImpl) EchoPingPong(ctx context.Context, req *echo.Request) (resp *echo.Response, err error) {
	klog.Infof("EchoPingPong is called, req = %+v", req)
	resp = &echo.Response{
		Message: req.Message,
	}
	return
}

// EchoUnary is unary api over streaming protocol
// NOT recommended; please just use KitexThrift api (without the annotation in your IDL)
func (s *TestServiceImpl) EchoUnary(ctx context.Context, req *echo.Request) (resp *echo.Response, err error) {
	klog.Infof("EchoUnary is called, req = %+v", req)
	resp = &echo.Response{
		Message: req.Message,
	}
	return
}
