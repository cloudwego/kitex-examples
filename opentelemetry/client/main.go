package main

import (
	"context"
	"time"

	"github.com/cloudwego/kitex-examples/kitex_gen/api"
	"github.com/cloudwego/kitex-examples/kitex_gen/api/echo"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
)

func main() {
	klog.SetLogger(kitexlogrus.NewLogger())
	klog.SetLevel(klog.LevelDebug)

	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName("echo-client"),
		provider.WithExportEndpoint("host.docker.internal:4317"),
	)
	defer p.Shutdown(context.Background())

	c, err := echo.NewClient(
		"echo",
		client.WithHostPorts("0.0.0.0:8181"),
		client.WithSuite(tracing.NewClientSuite(
			tracing.WithStackTrace(true),
		)),
	)
	if err != nil {
		klog.Fatal(err)
	}

	for {
		req := &api.Request{Message: "my request"}
		resp, err := c.Echo(context.Background(), req)
		if err != nil {
			klog.Fatal(err)
		}
		klog.Info(resp)
		time.Sleep(time.Second)
	}
}
