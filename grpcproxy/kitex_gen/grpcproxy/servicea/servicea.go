// Code generated by Kitex v0.3.2. DO NOT EDIT.

package servicea

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex-examples/grpcproxy/kitex_gen/grpcproxy"
	"github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	"github.com/cloudwego/kitex/pkg/streaming"
	"google.golang.org/protobuf/proto"
)

func serviceInfo() *kitex.ServiceInfo {
	return serviceAServiceInfo
}

var serviceAServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "ServiceA"
	handlerType := (*grpcproxy.ServiceA)(nil)
	methods := map[string]kitex.MethodInfo{
		"Chat": kitex.NewMethodInfo(chatHandler, newChatArgs, newChatResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "grpcproxy",
	}
	extra["streaming"] = true
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.3.2",
		Extra:           extra,
	}
	return svcInfo
}

func chatHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	st := arg.(*streaming.Args).Stream
	stream := &serviceAChatServer{st}
	return handler.(grpcproxy.ServiceA).Chat(stream)
}

type serviceAChatClient struct {
	streaming.Stream
}

func (x *serviceAChatClient) Send(m *grpcproxy.Request) error {
	return x.Stream.SendMsg(m)
}
func (x *serviceAChatClient) Recv() (*grpcproxy.Reply, error) {
	m := new(grpcproxy.Reply)
	return m, x.Stream.RecvMsg(m)
}

type serviceAChatServer struct {
	streaming.Stream
}

func (x *serviceAChatServer) Send(m *grpcproxy.Reply) error {
	return x.Stream.SendMsg(m)
}

func (x *serviceAChatServer) Recv() (*grpcproxy.Request, error) {
	m := new(grpcproxy.Request)
	return m, x.Stream.RecvMsg(m)
}

func newChatArgs() interface{} {
	return &ChatArgs{}
}

func newChatResult() interface{} {
	return &ChatResult{}
}

type ChatArgs struct {
	Req *grpcproxy.Request
}

func (p *ChatArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in ChatArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *ChatArgs) Unmarshal(in []byte) error {
	msg := new(grpcproxy.Request)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var ChatArgs_Req_DEFAULT *grpcproxy.Request

func (p *ChatArgs) GetReq() *grpcproxy.Request {
	if !p.IsSetReq() {
		return ChatArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ChatArgs) IsSetReq() bool {
	return p.Req != nil
}

type ChatResult struct {
	Success *grpcproxy.Reply
}

var ChatResult_Success_DEFAULT *grpcproxy.Reply

func (p *ChatResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in ChatResult")
	}
	return proto.Marshal(p.Success)
}

func (p *ChatResult) Unmarshal(in []byte) error {
	msg := new(grpcproxy.Reply)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ChatResult) GetSuccess() *grpcproxy.Reply {
	if !p.IsSetSuccess() {
		return ChatResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ChatResult) SetSuccess(x interface{}) {
	p.Success = x.(*grpcproxy.Reply)
}

func (p *ChatResult) IsSetSuccess() bool {
	return p.Success != nil
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Chat(ctx context.Context) (ServiceA_ChatClient, error) {
	streamClient, ok := p.c.(client.Streaming)
	if !ok {
		return nil, fmt.Errorf("client not support streaming")
	}
	res := new(streaming.Result)
	err := streamClient.Stream(ctx, "Chat", nil, res)
	if err != nil {
		return nil, err
	}
	stream := &serviceAChatClient{res.Stream}
	return stream, nil
}
