package src

import (
	"context"
	"fmt"
	"time"

	"github.com/bytedance/gopkg/cloud/metainfo"
	"github.com/cloudwego/kitex/client"
	xds2 "github.com/cloudwego/kitex/pkg/xds"
	"github.com/cloudwego/kitex/transport"
	"github.com/kitex-contrib/xds"
	"github.com/kitex-contrib/xds/xdssuite"

	"github.com/cloudwego/kitex-proxyless-test/service/codec/thrift/kitex_gen/proxyless"
	"github.com/cloudwego/kitex-proxyless-test/service/codec/thrift/kitex_gen/proxyless/greetservice"
)

type ProxylessClient struct {
	cli greetservice.Client
}

var (
	routeKey   = "stage"
	routeValue = "canary"
)

func routeByStage(ctx context.Context) map[string]string {
	if v, ok := metainfo.GetValue(ctx, routeKey); ok {
		return map[string]string{
			routeKey: v,
		}
	}
	return nil
}

func NewProxylessClient(targetService string) TestService {
	err := xds.Init()
	if err != nil {
		panic(err)
	}

	cli, err := greetservice.NewClient(
		targetService,
		client.WithXDSSuite(xds2.ClientSuite{
			RouterMiddleware: xdssuite.NewXDSRouterMiddleware(
				xdssuite.WithRouterMetaExtractor(routeByStage),
			),
			Resolver: xdssuite.NewXDSResolver(),
		}),
		client.WithTransportProtocol(transport.TTHeader),
	)
	if err != nil {
		panic(err)
	}
	return &ProxylessClient{cli: cli}
}

func (c *ProxylessClient) Run() error {
	for {
		req := &proxyless.HelloRequest{Message: "Hello!"}
		ctx := metainfo.WithValue(context.Background(), routeKey, routeValue) // set route meta for "stage": "canary"
		ctx = metainfo.WithBackwardValues(ctx)
		resp, err := c.cli.SayHello2(ctx, req)
		if err != nil {
			fmt.Println(err)
		} else {
			rip, _ := metainfo.RecvBackwardValue(ctx, PodNameKey)
			fmt.Printf("Received response: %s, from %s\n", resp.Message, rip)
		}
		time.Sleep(time.Second)
	}
}
