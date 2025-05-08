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
	"log"
	"sync"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/genericclient"
	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/transport"
)

func main() {
	// 1. Create Thrift file provider
	p, err := generic.NewThriftFileProvider("../idl/streaming.thrift")
	if err != nil {
		log.Fatal(err)
	}

	// 2. Create JSON Thrift generic call
	g, err := generic.JSONThriftGeneric(p)
	if err != nil {
		log.Fatal(err)
	}

	// 3. Create streaming client
	cli, err := genericclient.NewStreamingClient(
		"streaming_service",
		g,
		client.WithTransportProtocol(transport.GRPC),
		client.WithHostPorts("127.0.0.1:8888"),
	)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	// 4.1 Test pingpong mode
	fmt.Println("Testing EchoPingPong...")
	if err := testEchoPingPong(ctx, cli); err != nil {
		log.Printf("EchoPingPong failed: %v", err)
	}

	// 4.2 Test client streaming mode
	fmt.Println("\nTesting EchoClient...")
	if err := testEchoClient(ctx, cli); err != nil {
		log.Printf("EchoClient failed: %v", err)
	}

	// 4.3 Test server streaming mode
	fmt.Println("\nTesting EchoServer...")
	if err := testEchoServer(ctx, cli); err != nil {
		log.Printf("EchoServer failed: %v", err)
	}

	// 4.4 Test bidirectional streaming mode
	fmt.Println("\nTesting Echo (Bidirectional)...")
	if err := testEchoBidirectional(ctx, cli); err != nil {
		log.Printf("Echo failed: %v", err)
	}
}

// Test pingpong mode
func testEchoPingPong(ctx context.Context, cli genericclient.Client) error {
	resp, err := cli.GenericCall(ctx, "EchoPingPong", `{"message": "pingpong request"}`)
	if err != nil {
		return err
	}
	strResp, ok := resp.(string)
	if ok {
		fmt.Printf("EchoPingPong response: %v\n", strResp)
	}
	return nil
}

// Test client streaming mode
func testEchoClient(ctx context.Context, cli genericclient.Client) error {
	streamCli, err := genericclient.NewClientStreaming(ctx, cli, "EchoClient")
	if err != nil {
		return fmt.Errorf("failed to create client streaming: %v", err)
	}

	// Send multiple requests
	for i := 0; i < 3; i++ {
		req := fmt.Sprintf(`{"message": "grpc client streaming generic %dth request"}`, i)
		if err = streamCli.Send(req); err != nil {
			return fmt.Errorf("failed to send: %v", err)
		}
		time.Sleep(time.Second)
	}

	// Receive final response
	resp, err := streamCli.CloseAndRecv()
	if err != nil {
		return fmt.Errorf("failed to receive: %v", err)
	}

	strResp, ok := resp.(string)
	if ok {
		fmt.Printf("EchoClient response: %v\n", strResp)
	}
	return nil
}

// Test server streaming mode
func testEchoServer(ctx context.Context, cli genericclient.Client) error {
	streamCli, err := genericclient.NewServerStreaming(ctx, cli, "EchoServer", `{"message": "grpc server streaming generic request"}`)
	if err != nil {
		return fmt.Errorf("failed to create server streaming: %v", err)
	}

	// Receive multiple responses
	for {
		resp, err := streamCli.Recv()
		if err == io.EOF {
			fmt.Println("Server streaming message receive done. stream is closed")
			break
		} else if err != nil {
			return fmt.Errorf("failed to receive: %v", err)
		}

		strResp, ok := resp.(string)
		if ok {
			fmt.Printf("EchoServer response: %v\n", strResp)
		}
	}
	return nil
}

// Test bidirectional streaming mode
func testEchoBidirectional(ctx context.Context, cli genericclient.Client) error {
	streamCli, err := genericclient.NewBidirectionalStreaming(ctx, cli, "Echo")
	if err != nil {
		return fmt.Errorf("failed to create bidirectional streaming: %v", err)
	}

	wg := &sync.WaitGroup{}
	wg.Add(2)
	var sendErr, recvErr error

	// Send messages
	go func() {
		defer func() {
			if p := recover(); p != nil {
				sendErr = fmt.Errorf("panic: %v", p)
			}
			wg.Done()
		}()
		defer streamCli.Close()

		for i := 0; i < 3; i++ {
			req := fmt.Sprintf(`{"message": "grpc bidirectional streaming generic %dth request"}`, i)
			if err = streamCli.Send(req); err != nil {
				sendErr = fmt.Errorf("bidirectionalStreaming send: failed, err = %v", err)
				break
			}
			klog.Infof("BidirectionalStreamingTest send: req = %+v", req)
		}
	}()

	// Receive messages
	go func() {
		defer func() {
			if p := recover(); p != nil {
				recvErr = fmt.Errorf("panic: %v", p)
			}
			wg.Done()
		}()

		for {
			resp, err := streamCli.Recv()
			if err == io.EOF {
				klog.Infof("bidirectionalStreaming message receive done. stream is closed")
				break
			} else if err != nil {
				recvErr = fmt.Errorf("failed to recv: %v", err)
				break
			}

			strResp, ok := resp.(string)
			if ok {
				klog.Infof("bidirectionalStreaming message received: %+v", strResp)
			}
		}
	}()

	wg.Wait()

	// Check errors
	if sendErr != nil {
		return sendErr
	}
	if recvErr != nil {
		return recvErr
	}
	return nil
}