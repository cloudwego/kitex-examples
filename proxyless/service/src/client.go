/*
 * Copyright 2022 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package src

import (
	"context"
	"fmt"
	"time"

	"github.com/bytedance/gopkg/cloud/metainfo"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	xds2 "github.com/cloudwego/kitex/pkg/xds"
	"github.com/cloudwego/kitex/transport"
	"github.com/kitex-contrib/xds"
	"github.com/kitex-contrib/xds/core/manager"
	"github.com/kitex-contrib/xds/xdssuite"

	"github.com/cloudwego/kitex-examples/proxyless/service/codec/thrift/kitex_gen/proxyless"
	"github.com/cloudwego/kitex-examples/proxyless/service/codec/thrift/kitex_gen/proxyless/greetservice"
)

type ProxylessClient struct {
	cli1 greetservice.Client
	cli2 greetservice.Client
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

func NewProxylessClient(targetService1, targetService2 string) TestService {
	err := xds.Init(xds.WithXDSServerConfig(&manager.XDSServerConfig{SvrAddr: "istiod.istio-system.svc:15012", XDSAuth: true}))
	if err != nil {
		panic(err)
	}

	klog.Infof("service1: %s, service2: %s\n", targetService1, targetService2)
	cli1, err := greetservice.NewClient(
		targetService1,
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

	cli2, err := greetservice.NewClient(
		targetService2,
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

	return &ProxylessClient{cli1: cli1, cli2: cli2}
}

func (c *ProxylessClient) Run() error {
	for {
		req := &proxyless.HelloRequest{Message: "Hello!"}
		ctx := metainfo.WithValue(context.Background(), routeKey, routeValue) // set route meta for "stage": "canary"
		ctx = metainfo.WithBackwardValues(ctx)
		resp, err := c.cli1.SayHello2(ctx, req)
		if err != nil {
			fmt.Println(err)
		} else {
			rip, _ := metainfo.RecvBackwardValue(ctx, PodNameKey)
			fmt.Printf("Received response: %s, from %s\n", resp.Message, rip)
		}

		resp, err = c.cli2.SayHello2(ctx, req)
		if err != nil {
			fmt.Println(err)
		} else {
			rip, _ := metainfo.RecvBackwardValue(ctx, PodNameKey)
			fmt.Printf("Received response: %s, from %s\n", resp.Message, rip)
		}
		time.Sleep(time.Second)
	}
}
