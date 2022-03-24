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
	"os"
	"time"

	"github.com/cloudwego/kitex-examples/kitex_gen/api"
	"github.com/cloudwego/kitex-examples/kitex_gen/api/echo"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	"go.opentelemetry.io/otel"
)

func main() {
	klog.SetLogger(kitexlogrus.NewLogger())
	klog.SetLevel(klog.LevelDebug)

	serviceName := "echo-client"

	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(serviceName),
		// Support setting ExportEndpoint via environment variables: OTEL_EXPORTER_OTLP_ENDPOINT
		//provider.WithExportEndpoint("host.docker.internal:4317"),
		provider.WithInsecure(),
	)
	defer p.Shutdown(context.Background())

	demoServerAddr, ok := os.LookupEnv("DEMO_SERVER_ENDPOINT")
	if !ok {
		demoServerAddr = "0.0.0.0:8181"
	}

	c, err := echo.NewClient(
		"echo",
		client.WithHostPorts(demoServerAddr),
		client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: serviceName}),
	)
	if err != nil {
		klog.Fatal(err)
	}

	for {
		call(c)
		<-time.After(time.Second)
	}
}

func call(c echo.Client) {
	ctx, span := otel.Tracer("client").Start(context.Background(), "root")
	defer span.End()

	req := &api.Request{Message: "my request"}
	resp, err := c.Echo(ctx, req)
	if err != nil {
		klog.CtxErrorf(ctx, "err %v", err)
	}

	klog.CtxInfof(ctx, "req:%v, res:%v", req, resp)
}
