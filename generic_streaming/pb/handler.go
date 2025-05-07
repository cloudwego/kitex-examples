package main

import (
	"context"
	"fmt"
	pb "pb_generic_streaming_demo/kitex_gen/pb"
)

// StreamingServiceImpl implements the last service interface defined in the IDL.
type StreamingServiceImpl struct{}

// StreamRequestEcho implements client streaming:
// - Client sends multiple messages
// - Server returns a single response
func (s *StreamingServiceImpl) StreamRequestEcho(stream pb.StreamingService_StreamRequestEchoServer) (err error) {
	var messages []string
	for {
		req, err := stream.Recv()
		if err != nil {
			break
		}
		messages = append(messages, req.Message)
		fmt.Printf("Received message: %s\n", req.Message)
	}
	
	// Return summary of all received messages
	resp := &pb.Response{
		Message: fmt.Sprintf("Received %d messages: %v", len(messages), messages),
	}
	return stream.SendAndClose(resp)
}

// StreamResponseEcho implements server streaming:
// - Client sends a single request
// - Server returns multiple responses
func (s *StreamingServiceImpl) StreamResponseEcho(req *pb.Request, stream pb.StreamingService_StreamResponseEchoServer) (err error) {
	// Simulate LLM scenario, return multiple responses
	for i := 0; i < 3; i++ {
		resp := &pb.Response{
			Message: fmt.Sprintf("Response %d for: %s", i+1, req.Message),
		}
		if err := stream.Send(resp); err != nil {
			return err
		}
	}
	return nil
}

// BidirectionalEcho implements bidirectional streaming:
// - Both client and server can send multiple messages
// - Messages can be sent in any order
func (s *StreamingServiceImpl) BidirectionalEcho(stream pb.StreamingService_BidirectionalEchoServer) (err error) {
	for {
		req, err := stream.Recv()
		if err != nil {
			break
		}
		fmt.Printf("Received: %s\n", req.Message)
		
		// Send response
		resp := &pb.Response{
			Message: fmt.Sprintf("Echo: %s", req.Message),
		}
		if err := stream.Send(resp); err != nil {
			return err
		}
	}
	return nil
}

// UnaryEcho implements traditional request-response:
// - Client sends a single request
// - Server returns a single response
func (s *StreamingServiceImpl) UnaryEcho(ctx context.Context, req *pb.Request) (resp *pb.Response, err error) {
	return &pb.Response{
		Message: fmt.Sprintf("Echo: %s", req.Message),
	}, nil
}
