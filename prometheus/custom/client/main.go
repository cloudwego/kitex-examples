/*
 * Copyright 2024 CloudWeGo Authors
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
	"net/http"
	"time"

	"github.com/cloudwego/kitex-examples/kitex_gen/pbapi"
	"github.com/cloudwego/kitex-examples/kitex_gen/pbapi/echo"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"
	prometheus "github.com/kitex-contrib/monitor-prometheus"
	promlib "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	registry     = promlib.NewRegistry()
	clientTracer = prometheus.NewClientTracer("", "", prometheus.WithRegistry(registry), prometheus.WithDisableServer(true))
)

func main() {
	http.Handle("/metrics", promhttp.HandlerFor(registry, promhttp.HandlerOpts{}))
	go http.ListenAndServe(":8093", nil) //nolint:errcheck
	cli, err := echo.NewClient("grpc-server", client.WithHostPorts("localhost:8082"), client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{
		ServiceName: "grpc-client",
	}), client.WithMetaHandler(transmeta.ClientHTTP2Handler), client.WithTransportProtocol(transport.GRPC), client.WithTracer(clientTracer))
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		log.Fatal(err)
	}
	for {
		resp, err := cli.Echo(context.Background(), &pbapi.Request{Message: "hello"})
		if err != nil {
			log.Fatal(err)
		}
		klog.Info(resp)
		time.Sleep(time.Second)
	}
}
