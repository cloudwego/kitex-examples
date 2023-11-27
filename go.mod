module github.com/cloudwego/kitex-examples

go 1.16

replace github.com/apache/thrift => github.com/apache/thrift v0.13.0

require (
	github.com/apache/thrift v0.16.0
	github.com/bytedance/gopkg v0.0.0-20230728082804-614d0af6619b
	github.com/cloudwego/kitex v0.7.3
	github.com/kitex-contrib/monitor-prometheus v0.2.0
	github.com/kitex-contrib/obs-opentelemetry v0.2.5
	github.com/kitex-contrib/obs-opentelemetry/logging/logrus v0.0.0-20231103033707-6f0423a24fdf
	github.com/kitex-contrib/tracer-opentracing v0.0.3
	github.com/opentracing/opentracing-go v1.2.0
	github.com/uber/jaeger-client-go v2.29.1+incompatible
	go.opentelemetry.io/otel v1.19.0
	golang.org/x/sync v0.3.0
	google.golang.org/protobuf v1.31.0
)

require (
	github.com/HdrHistogram/hdrhistogram-go v1.1.2 // indirect
	github.com/bytedance/mockey v1.2.1 // indirect
	github.com/cloudwego/fastpb v0.0.4
	github.com/cloudwego/frugal v0.1.8
	github.com/kitex-contrib/obs-opentelemetry/logging/zap v0.0.0-20231103033707-6f0423a24fdf
	github.com/sirupsen/logrus v1.9.2
	github.com/tidwall/gjson v1.14.4 // indirect
	github.com/uber/jaeger-lib v2.4.1+incompatible // indirect
	go.uber.org/zap v1.26.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
)
