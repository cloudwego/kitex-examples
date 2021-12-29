module github.com/cloudwego/kitex-examples/bizdemo/easy_note/note

go 1.16

require (
	github.com/apache/thrift v0.15.0
	github.com/cloudwego/kitex v0.1.1
	github.com/kitex-contrib/registry-etcd v0.0.0-20211207110456-45d9f8b00e1f
	github.com/kitex-contrib/tracer-opentracing v0.0.3
	github.com/opentracing/opentracing-go v1.2.0
	github.com/shirou/gopsutil v3.21.11+incompatible
	github.com/uber/jaeger-client-go v2.30.0+incompatible
	github.com/yusufpapurcu/wmi v1.2.2 // indirect
	gorm.io/driver/mysql v1.2.1
	gorm.io/gorm v1.22.4
	gorm.io/plugin/opentracing v0.0.0-20211213123506-2713cb54af60
)

replace github.com/apache/thrift => github.com/apache/thrift v0.13.0
