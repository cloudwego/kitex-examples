// Copyright 2021 CloudWeGo Authors
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
	"fmt"
	"io"
	"log"
	"strings"
	"time"

	"github.com/cloudwego/kitex-examples/streaming/kitex_gen/pbapi"
	"github.com/cloudwego/kitex-examples/streaming/kitex_gen/pbapi/echo"
	"github.com/cloudwego/kitex/pkg/klog"
	"golang.org/x/sync/errgroup"
)

var _ pbapi.Echo = &EchoImpl{}

// EchoImpl implements the last service interface defined in the IDL.
type EchoImpl struct{}

// StreamRequestEcho implements the Echo interface.
func (s *EchoImpl) StreamRequestEcho(stream pbapi.Echo_StreamRequestEchoServer) (err error) {
	klog.Info("StreamRequestEcho called")
	var msgs []string
	for {
		req, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		klog.Infof("message received: %v\n", req.Message)
		msgs = append(msgs, req.Message)
		time.Sleep(time.Second)
	}
	return stream.SendAndClose(&pbapi.Response{Message: "all message: " + strings.Join(msgs, ", ")})
}

// StreamResponseEcho implements the Echo interface.
func (s *EchoImpl) StreamResponseEcho(req *pbapi.Request, stream pbapi.Echo_StreamResponseEchoServer) (err error) {
	klog.Info("StreamResponseEcho called")
	resp := &pbapi.Response{}
	for i := 0; i < 10; i++ {
		resp.Message = fmt.Sprintf("%v -> %dth response", req.Message, i)
		err := stream.Send(resp)
		if err != nil {
			return err
		}
		time.Sleep(time.Second)
	}
	return stream.Close()
}

// BidirectionalEcho implements the Echo interface.
func (s *EchoImpl) BidirectionalEcho(stream pbapi.Echo_BidirectionalEchoServer) (err error) {
	klog.Info("BidirectionalEcho called")
	var eg errgroup.Group
	eg.Go(func() error {
		for {
			req, err := stream.Recv()
			if err != nil {
				if err == io.EOF {
					return nil
				}
				return err
			}
			klog.Infof("message received: %v\n", req.Message)
			time.Sleep(time.Second)
		}
	})
	eg.Go(func() error {
		resp := &pbapi.Response{}
		var cnt int
		for {
			resp.Message = fmt.Sprintf("%dth response", cnt)
			err := stream.Send(resp)
			if err != nil {
				return err
			}
			cnt++
			time.Sleep(time.Second)
		}
	})
	return eg.Wait()
}

func main() {
	svr := echo.NewServer(new(EchoImpl))
	if err := svr.Run(); err != nil {
		log.Println("server stopped with error:", err)
	} else {
		log.Println("server stopped")
	}
}
