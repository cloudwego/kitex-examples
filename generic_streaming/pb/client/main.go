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

package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"sync"

	"github.com/cloudwego/dynamicgo/proto"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/genericclient"
	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"
)

const (
	idlPath = "../idl/streaming.proto"
)

func main() {
	ctx := context.Background()

	// Initialize generic client
	dOpts := proto.Options{}
	p, err := generic.NewPbFileProviderWithDynamicGo(idlPath, ctx, dOpts)
	if err != nil {
		log.Fatal(err)
	}

	// Create JSON generic
	g, err := generic.JSONPbGeneric(p)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize streaming client
	cli, err := genericclient.NewStreamingClient(
		"streaming",
		g,
		client.WithTransportProtocol(transport.GRPC),
		client.WithHostPorts("127.0.0.1:8888"),
		client.WithMetaHandler(transmeta.ClientHTTP2Handler),
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\n=== Testing UnaryEcho (Request-Response Mode) ===")
	if err := testUnaryEcho(ctx, cli); err != nil {
		log.Printf("UnaryEcho failed: %v", err)
	}

	fmt.Println("\n=== Testing StreamRequestEcho (Client Streaming) ===")
	if err := testStreamRequestEcho(ctx, cli); err != nil {
		log.Printf("StreamRequestEcho failed: %v", err)
	}

	fmt.Println("\n=== Testing StreamResponseEcho (Server Streaming) ===")
	if err := testStreamResponseEcho(ctx, cli); err != nil {
		log.Printf("StreamResponseEcho failed: %v", err)
	}

	fmt.Println("\n=== Testing BidirectionalEcho (Bidirectional Streaming) ===")
	if err := testBidirectionalEcho(ctx, cli); err != nil {
		log.Printf("BidirectionalEcho failed: %v", err)
	}
}

func testUnaryEcho(ctx context.Context, cli genericclient.Client) error {
	req := `{"message": "Hello from unary echo"}`
	resp, err := cli.GenericCall(ctx, "UnaryEcho", req)
	if err != nil {
		return err
	}
	fmt.Printf("UnaryEcho response: %v\n", resp)
	return nil
}

func testStreamRequestEcho(ctx context.Context, cli genericclient.Client) error {
	streamCli, err := genericclient.NewClientStreaming(ctx, cli, "StreamRequestEcho")
	if err != nil {
		return fmt.Errorf("failed to create client streaming: %v", err)
	}

	// Send multiple messages
	messages := []string{"message1", "message2", "message3"}
	for _, msg := range messages {
		req := fmt.Sprintf(`{"message": "%s"}`, msg)
		if err := streamCli.Send(req); err != nil {
			return fmt.Errorf("failed to send: %v", err)
		}
		fmt.Printf("Sending message: %s\n", msg)
	}

	// Close send and wait for response
	resp, err := streamCli.CloseAndRecv()
	if err != nil {
		return fmt.Errorf("failed to receive: %v", err)
	}
	fmt.Printf("StreamRequestEcho response: %v\n", resp)
	return nil
}

func testStreamResponseEcho(ctx context.Context, cli genericclient.Client) error {
	req := `{"message": "Hello from stream response echo"}`
	streamCli, err := genericclient.NewServerStreaming(ctx, cli, "StreamResponseEcho", req)
	if err != nil {
		return fmt.Errorf("failed to create server streaming: %v", err)
	}

	fmt.Println("Sending request:", req)

	// Receive multiple responses
	fmt.Println("Starting to receive server responses...")
	for {
		resp, err := streamCli.Recv()
		if err == io.EOF {
			fmt.Println("Server streaming message receive done. stream is closed")
			break
		} else if err != nil {
			return fmt.Errorf("failed to receive: %v", err)
		}
		fmt.Printf("StreamResponseEcho response: %v\n", resp)
	}
	return nil
}

func testBidirectionalEcho(ctx context.Context, cli genericclient.Client) error {
	streamCli, err := genericclient.NewBidirectionalStreaming(ctx, cli, "BidirectionalEcho")
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

		messages := []string{"bidirectional1", "bidirectional2", "bidirectional3"}
		for _, msg := range messages {
			req := fmt.Sprintf(`{"message": "%s"}`, msg)
			if err := streamCli.Send(req); err != nil {
				sendErr = fmt.Errorf("failed to send: %v", err)
				break
			}
			klog.Infof("Sending message: %s", msg)
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
				klog.Infof("Bidirectional streaming message receive done. stream is closed")
				break
			} else if err != nil {
				recvErr = fmt.Errorf("failed to receive: %v", err)
				break
			}
			klog.Infof("Received response: %v", resp)
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