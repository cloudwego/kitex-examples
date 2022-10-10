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
	"log"
	"time"

	"github.com/bytedance/gopkg/cloud/metainfo"
	"github.com/cloudwego/kitex-examples/kitex_gen/api"
	"github.com/cloudwego/kitex-examples/kitex_gen/api/echo"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/transport"
)

func main() {
	ctx := context.Background()

	// must mark the context to receive backward meta information
	ctx = metainfo.WithBackwardValues(ctx)

	cli, err := echo.NewClient(
		"echo",
		client.WithHostPorts("[::1]:8888"),
		// must use the underlying transport protocol that supports metainfo, such as TTHeader, HTTP
		client.WithTransportProtocol(transport.TTHeader),
	)
	if err != nil {
		log.Fatal(err)
	}
	for {
		req := &api.Request{Message: "my request"}
		_, err = cli.Echo(ctx, req)
		if err != nil {
			log.Fatal(err)
		}
		if err == nil {
			// receive the meta information from server side
			val, ok := metainfo.RecvBackwardValue(ctx, "something-from-server")
			if ok {
				klog.Infof("something-from-server:%s", val)
			} else {
				klog.Warn("`something-from-server` not exist")
			}
		}
		time.Sleep(time.Second)
	}
}
