module github.com/cloudwego/kitex-examples

go 1.16

require (
	github.com/apache/thrift v0.13.0
	github.com/bytedance/gopkg v0.0.0-20220531084716-665b4f21126f
	github.com/cloudwego/kitex v0.4.4
	github.com/kitex-contrib/monitor-prometheus v0.0.0-20210817080809-024dd7bd51e1
	github.com/kitex-contrib/obs-opentelemetry v0.0.0-20220601144657-c60210e3c928
	github.com/kitex-contrib/obs-opentelemetry/logging/logrus v0.0.0-20220601144657-c60210e3c928
	github.com/kitex-contrib/tracer-opentracing v0.0.2
	github.com/opentracing/opentracing-go v1.2.0
	github.com/uber/jaeger-client-go v2.29.1+incompatible
	go.opentelemetry.io/otel v1.7.0
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	google.golang.org/protobuf v1.28.1
)

require (
	github.com/HdrHistogram/hdrhistogram-go v1.1.2 // indirect
	github.com/uber/jaeger-lib v2.4.1+incompatible // indirect
	go.uber.org/atomic v1.8.0 // indirect
)
