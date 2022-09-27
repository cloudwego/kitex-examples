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

package main

import (
	"context"
	"log"
	"time"

	"github.com/cloudwego/kitex-examples/kitex_gen/api"
	"github.com/cloudwego/kitex-examples/kitex_gen/api/echo"

	kclient "github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"
	prometheus "github.com/kitex-contrib/monitor-prometheus"
)

func main() {
	client, err := echo.NewClient("echo",
		kclient.WithHostPorts("0.0.0.0:8888"),
		kclient.WithTracer(prometheus.NewClientTracer(":9093", "/metrics")),
		kclient.WithMetaHandler(transmeta.ClientTTHeaderHandler),
		kclient.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "kitex-client"}),
		kclient.WithTransportProtocol(transport.TTHeader),
	)
	if err != nil {
		log.Fatal(err)
	}
	for {
		req := &api.Request{Message: "my request"}
		resp, err := client.Echo(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}
		klog.Info(resp)
		time.Sleep(time.Second)
	}
}
