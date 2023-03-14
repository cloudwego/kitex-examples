module github.com/cloudwego/kitex-examples/bizdemo/easy_note

go 1.16

require (
	github.com/apache/thrift v0.15.0
	github.com/cloudwego/hertz v0.6.0
	github.com/cloudwego/kitex v0.5.0
	github.com/hertz-contrib/jwt v1.0.2
	github.com/kitex-contrib/registry-etcd v0.1.0
	github.com/kitex-contrib/tracer-opentracing v0.0.3
	github.com/opentracing/opentracing-go v1.2.0
	github.com/shirou/gopsutil v3.21.11+incompatible
	github.com/tklauser/go-sysconf v0.3.9 // indirect
	github.com/uber/jaeger-client-go v2.30.0+incompatible
	github.com/yusufpapurcu/wmi v1.2.2 // indirect
	google.golang.org/protobuf v1.28.1
	gorm.io/driver/mysql v1.4.7
	gorm.io/gorm v1.24.6
	gorm.io/plugin/opentracing v0.0.0-20211220013347-7d2b2af23560
)

replace github.com/apache/thrift => github.com/apache/thrift v0.13.0
