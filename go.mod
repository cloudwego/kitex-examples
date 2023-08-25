module github.com/cloudwego/kitex-examples

go 1.16

replace github.com/apache/thrift => github.com/apache/thrift v0.13.0

require (
	github.com/apache/thrift v0.16.0
	github.com/bytedance/gopkg v0.0.0-20230728082804-614d0af6619b
	github.com/cloudwego/kitex v0.7.0
	github.com/kitex-contrib/monitor-prometheus v0.1.0
	github.com/kitex-contrib/obs-opentelemetry v0.2.3
	github.com/kitex-contrib/obs-opentelemetry/logging/logrus v0.0.0-20230819133448-76093321aa8e
	github.com/kitex-contrib/tracer-opentracing v0.0.3
	github.com/opentracing/opentracing-go v1.2.0
	github.com/uber/jaeger-client-go v2.29.1+incompatible
	go.opentelemetry.io/otel v1.16.0
	golang.org/x/sync v0.1.0
	google.golang.org/protobuf v1.30.0
)

require (
	github.com/HdrHistogram/hdrhistogram-go v1.1.2 // indirect
	github.com/uber/jaeger-lib v2.4.1+incompatible // indirect
	go.uber.org/atomic v1.10.0 // indirect
)
