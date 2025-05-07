package main

import (
	"context"
	"io"
	"strconv"

	"github.com/cloudwego/kitex/pkg/klog"
	echo "generic_streaming_demo_thrift/kitex_gen/echo"
)

// TestServiceImpl implements the last service interface defined in the IDL.
type TestServiceImpl struct{}

// Echo implements bidirectional streaming
func (s *TestServiceImpl) Echo(stream echo.TestService_EchoServer) (err error) {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			klog.Infof("Echo: stream closed by client")
			return nil
		}
		if err != nil {
			klog.Errorf("Echo: failed to receive request: %v", err)
			return err
		}

		// Echo back the message with a prefix
		resp := &echo.Response{
			Message: "server echo: " + req.Message,
		}
		if err := stream.Send(resp); err != nil {
			klog.Errorf("Echo: failed to send response: %v", err)
			return err
		}
		klog.Infof("Echo: received request: %v, sent response: %v", req, resp)
	}
}

// EchoClient implements client streaming
func (s *TestServiceImpl) EchoClient(stream echo.TestService_EchoClientServer) (err error) {
	var messageCount int
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// Client has finished sending
			resp := &echo.Response{
				Message: "server received " + strconv.Itoa(messageCount) + " messages",
			}
			return stream.SendAndClose(resp)
		}
		if err != nil {
			klog.Errorf("EchoClient: failed to receive request: %v", err)
			return err
		}
		messageCount++
		klog.Infof("EchoClient: received message %d: %v", messageCount, req)
	}
}

// EchoServer implements server streaming
func (s *TestServiceImpl) EchoServer(req *echo.Request, stream echo.TestService_EchoServerServer) (err error) {
	// Send multiple responses for a single request
	for i := 0; i < 3; i++ {
		resp := &echo.Response{
			Message: "server streaming response " + strconv.Itoa(i) + " for request: " + req.Message,
		}
		if err := stream.Send(resp); err != nil {
			klog.Errorf("EchoServer: failed to send response: %v", err)
			return err
		}
		klog.Infof("EchoServer: sent response %d: %v", i, resp)
	}
	return nil
}

// EchoPingPong implements traditional request-response
func (s *TestServiceImpl) EchoPingPong(ctx context.Context, req *echo.Request) (resp *echo.Response, err error) {
	resp = &echo.Response{
		Message: "server pong: " + req.Message,
	}
	klog.Infof("EchoPingPong: received request: %v, sending response: %v", req, resp)
	return resp, nil
}
