module github.com/cloudwego/kitex-examples

go 1.16

require (
	github.com/apache/thrift v0.13.0
	github.com/bytedance/gopkg v0.0.0-20220817015305-b879a72dc90f
	github.com/cloudwego/kitex v0.5.0
	github.com/kitex-contrib/monitor-prometheus v0.1.0
	github.com/kitex-contrib/obs-opentelemetry v0.2.1
	github.com/kitex-contrib/obs-opentelemetry/logging/logrus v0.0.0-20230219092456-5e6c84962323
	github.com/kitex-contrib/tracer-opentracing v0.0.3
	github.com/opentracing/opentracing-go v1.2.0
	github.com/uber/jaeger-client-go v2.29.1+incompatible
	go.opentelemetry.io/otel v1.13.0
	golang.org/x/sync v0.1.0
	google.golang.org/protobuf v1.28.1
)

require (
	github.com/HdrHistogram/hdrhistogram-go v1.1.2 // indirect
	github.com/uber/jaeger-lib v2.4.1+incompatible // indirect
)
