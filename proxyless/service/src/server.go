package src

import (
	"context"
	"fmt"
	"os"

	"github.com/bytedance/gopkg/cloud/metainfo"
	"github.com/cloudwego/kitex-proxyless-test/service/codec/thrift/kitex_gen/proxyless"
	"github.com/cloudwego/kitex-proxyless-test/service/codec/thrift/kitex_gen/proxyless/greetservice"
	"github.com/cloudwego/kitex/server"
)

// GreetServiceImpl implements the last service interface defined in the IDL.
type GreetServiceImpl struct{}

// SayHello1 implements the GreetServiceImpl interface.
func (s *GreetServiceImpl) SayHello1(ctx context.Context, request *proxyless.HelloRequest) (resp *proxyless.HelloResponse, err error) {
	// TODO: Your code here...
	resp = proxyless.NewHelloResponse()
	fmt.Println("SayHello1 Called")
	resp.SetMessage("Hello1!")
	if podName, ok := os.LookupEnv(PodNameKey); ok {
		metainfo.SendBackwardValue(ctx, PodNameKey, podName)
	}
	return
}

// SayHello2 implements the GreetServiceImpl interface.
func (s *GreetServiceImpl) SayHello2(ctx context.Context, request *proxyless.HelloRequest) (resp *proxyless.HelloResponse, err error) {
	// TODO: Your code here...
	resp = proxyless.NewHelloResponse()
	fmt.Println("SayHello2 Called")
	resp.SetMessage("Hello2!")
	if podName, ok := os.LookupEnv(PodNameKey); ok {
		metainfo.SendBackwardValue(ctx, PodNameKey, podName)
	}
	return
}

type ProxylessServer struct {
	svr server.Server
}

func NewProxylessServer() TestService {
	return &ProxylessServer{svr: greetservice.NewServer(&GreetServiceImpl{})}
}

func (s *ProxylessServer) Run() error {
	// :8888 is the default port for the server.
	err := s.svr.Run()
	return err
}
