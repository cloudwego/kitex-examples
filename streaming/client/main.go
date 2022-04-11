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
	"context"
	"fmt"
	"io"
	"time"

	"github.com/cloudwego/kitex-examples/streaming/kitex_gen/pbapi"
	"github.com/cloudwego/kitex-examples/streaming/kitex_gen/pbapi/echo"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"golang.org/x/sync/errgroup"
)

func doStreamRequestEcho(client echo.Client) error {
	// StreamRequestEcho
	streamCli, err := client.StreamRequestEcho(context.Background())
	if err != nil {
		return err
	}
	for i := 0; i < 10; i++ {
		req := &pbapi.Request{Message: fmt.Sprintf("doStreamRequestEcho %dth request", i)}
		if err := streamCli.Send(req); err != nil {
			return err
		}
		time.Sleep(time.Second)
	}
	resp, err := streamCli.CloseAndRecv()
	if err != nil {
		return err
	}
	klog.Infof("doStreamRequestEcho message received: %v\n", resp.Message)
	return nil
}

func doStreamResponseEcho(client echo.Client) error {
	req := &pbapi.Request{Message: "doStreamResponseEcho request"}
	// StreamResponseEcho
	streamCli, err := client.StreamResponseEcho(context.Background(), req)
	if err != nil {
		return err
	}
	for {
		if resp, err := streamCli.Recv(); err != nil {
			if err == io.EOF {
				klog.Info("doStreamResponseEcho receive done")
				return nil
			}
			return err
		} else {
			klog.Infof("doStreamResponseEcho message received: %v\n", resp.Message)
		}
		time.Sleep(time.Second)
	}
}

func doBidirectionalEcho(client echo.Client) error {
	// BidirectionalEcho
	streamCli, err := client.BidirectionalEcho(context.Background())
	if err != nil {
		return err
	}
	var eg errgroup.Group
	eg.Go(func() error {
		var cnt int
		for {
			req := &pbapi.Request{Message: fmt.Sprintf("doBidirectionalEcho %dth request", cnt)}
			if err := streamCli.Send(req); err != nil {
				return err
			}
			cnt++
			time.Sleep(time.Second)
		}
	})
	eg.Go(func() error {
		for {
			if resp, err := streamCli.Recv(); err != nil {
				if err == io.EOF {
					klog.Info("doBidirectionalEcho receive done")
					return nil
				}
				return err
			} else {
				klog.Infof("doBidirectionalEcho message received: %v\n", resp.Message)
			}
			time.Sleep(time.Second)
		}
	})
	return eg.Wait()
}

func main() {
	client, err := echo.NewClient("echo", client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		klog.Fatal(err)
	}
	klog.Info("doStreamRequestEcho start")
	if err := doStreamRequestEcho(client); err != nil {
		klog.Fatal(err)
	}
	klog.Info("doStreamRequestEcho finish")
	klog.Info("doStreamResponseEcho start")
	if err := doStreamResponseEcho(client); err != nil {
		klog.Fatal(err)
	}
	klog.Info("doStreamResponseEcho finish")
	klog.Info("doBidirectionalEcho start")
	if err := doBidirectionalEcho(client); err != nil {
		klog.Fatal(err)
	}
	klog.Info("doBidirectionalEcho finish")
}
