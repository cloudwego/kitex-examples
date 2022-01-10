module github.com/cloudwego/kitex-examples/bizdemo/easy_note/api

go 1.16

require (
	github.com/apache/thrift v0.15.0
	github.com/appleboy/gin-jwt/v2 v2.7.0
	github.com/cloudwego/kitex v0.1.3
	github.com/gin-gonic/gin v1.7.7
	github.com/kitex-contrib/registry-etcd v0.0.0-20211207110456-45d9f8b00e1f
	github.com/kitex-contrib/tracer-opentracing v0.0.3
	github.com/opentracing/opentracing-go v1.2.0
	github.com/uber/jaeger-client-go v2.30.0+incompatible
	google.golang.org/protobuf v1.27.1
)

replace github.com/apache/thrift => github.com/apache/thrift v0.13.0
