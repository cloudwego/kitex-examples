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
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/flow"
	"github.com/cloudwego/kitex-examples/hello/kitex_gen/api"
	"github.com/cloudwego/kitex-examples/hello/kitex_gen/api/hello"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	kitexSentinel "github.com/kitex-contrib/opensergo/sentinel"
)

const FakeErrorMsg = "fake error for testing"

type HelloImpl struct{}

func (s *HelloImpl) Echo(ctx context.Context, req *api.Request) (resp *api.Response, err error) {
	resp = &api.Response{Message: req.Message}
	return
}

func initSentinel() {
	err := sentinel.InitDefault()
	if err != nil {
		log.Fatalf("Unexpected error: %+v", err)
	}
	_, err = flow.LoadRules([]*flow.Rule{
		{
			Resource:               "hello/echo",
			Threshold:              1.0,
			TokenCalculateStrategy: flow.Direct,
			ControlBehavior:        flow.Reject,
		},
	})
	if err != nil {
		log.Fatalf("Unexpected error: %+v", err)
		return
	}
}

func main() {
	initSentinel()
	srv := hello.NewServer(new(HelloImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "hello"}),
		// The extensions are implemented as middleware,
		// but Kitex provides additional interfaces to support fusing and flow limiting,
		// and Sentinel's extensions do not necessarily take effect
		// when both are used at the same time
		server.WithMiddleware(kitexSentinel.SentinelServerMiddleware(
			// customize resource extractor if required
			// method_path by default
			kitexSentinel.WithResourceExtract(func(ctx context.Context, req, resp interface{}) string {
				rpcInfo := rpcinfo.GetRPCInfo(ctx)
				return rpcInfo.To().ServiceName() + "/" + rpcInfo.To().Method()
			}),
			// customize block fallback error message if required
			// abort with blockErr by default
			kitexSentinel.WithBlockFallback(func(ctx context.Context, req, resp interface{}, blockErr error) error {
				return errors.New(FakeErrorMsg)
			}),
		)),
	)
	defer srv.Stop()
	go srv.Run()
	time.Sleep(time.Second)

	c, err := hello.NewClient("hello", client.WithHostPorts(":8888"))
	if err != nil {
		log.Fatalf("Unexcepted err: %v", err)
	}
	req := &api.Request{}
	if _, err := c.Echo(context.Background(), req); err != nil {
		log.Fatalf("first request fail: %v", err)
	}
	if _, err := c.Echo(context.Background(), req); err != nil {
		fmt.Println("block err:", err.Error())
		fmt.Printf("blockErr equals FakeErrorMessage: %v", strings.Contains(err.Error(), FakeErrorMsg))
	}
}
