// Code generated by Kitex v0.9.0. DO NOT EDIT.

package serviceb

import (
	"context"
	"errors"
	service "github.com/cloudwego/kitex-examples/protobuf_multi_service/kitex_gen/multi/service"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"ChatB": kitex.NewMethodInfo(
		chatBHandler,
		newChatBArgs,
		newChatBResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
}

var (
	serviceBServiceInfo                = NewServiceInfo()
	serviceBServiceInfoForClient       = NewServiceInfoForClient()
	serviceBServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return serviceBServiceInfo
}

// for client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return serviceBServiceInfoForStreamClient
}

// for stream client
func serviceInfoForClient() *kitex.ServiceInfo {
	return serviceBServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "ServiceB"
	handlerType := (*service.ServiceB)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "multiservice",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.9.0",
		Extra:           extra,
	}
	return svcInfo
}

func chatBHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(service.RequestB)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(service.ServiceB).ChatB(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *ChatBArgs:
		success, err := handler.(service.ServiceB).ChatB(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*ChatBResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newChatBArgs() interface{} {
	return &ChatBArgs{}
}

func newChatBResult() interface{} {
	return &ChatBResult{}
}

type ChatBArgs struct {
	Req *service.RequestB
}

func (p *ChatBArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(service.RequestB)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *ChatBArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *ChatBArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *ChatBArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *ChatBArgs) Unmarshal(in []byte) error {
	msg := new(service.RequestB)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var ChatBArgs_Req_DEFAULT *service.RequestB

func (p *ChatBArgs) GetReq() *service.RequestB {
	if !p.IsSetReq() {
		return ChatBArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ChatBArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *ChatBArgs) GetFirstArgument() interface{} {
	return p.Req
}

type ChatBResult struct {
	Success *service.Reply
}

var ChatBResult_Success_DEFAULT *service.Reply

func (p *ChatBResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(service.Reply)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *ChatBResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *ChatBResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *ChatBResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *ChatBResult) Unmarshal(in []byte) error {
	msg := new(service.Reply)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ChatBResult) GetSuccess() *service.Reply {
	if !p.IsSetSuccess() {
		return ChatBResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ChatBResult) SetSuccess(x interface{}) {
	p.Success = x.(*service.Reply)
}

func (p *ChatBResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ChatBResult) GetResult() interface{} {
	return p.Success
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) ChatB(ctx context.Context, Req *service.RequestB) (r *service.Reply, err error) {
	var _args ChatBArgs
	_args.Req = Req
	var _result ChatBResult
	if err = p.c.Call(ctx, "ChatB", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
