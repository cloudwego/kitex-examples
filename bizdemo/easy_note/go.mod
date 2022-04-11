module github.com/cloudwego/kitex-examples/bizdemo/easy_note

go 1.16

require (
	github.com/apache/thrift v0.15.0
	github.com/appleboy/gin-jwt/v2 v2.8.0
	github.com/cloudwego/kitex v0.2.1
	github.com/gin-gonic/gin v1.7.7
	github.com/kitex-contrib/obs-opentelemetry v0.0.0-20220323032334-353a6ad0baaa
	github.com/kitex-contrib/obs-opentelemetry/logging/logrus v0.0.0-20220323032334-353a6ad0baaa
	github.com/kitex-contrib/registry-nacos v0.0.0-20220410054737-ef179f188106
	github.com/shirou/gopsutil v3.21.11+incompatible
	github.com/tklauser/go-sysconf v0.3.9 // indirect
	github.com/yusufpapurcu/wmi v1.2.2 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	google.golang.org/protobuf v1.27.1
	gorm.io/driver/mysql v1.2.3
	gorm.io/gorm v1.22.5
	gorm.io/plugin/opentracing v0.0.0-20211220013347-7d2b2af23560
)

replace github.com/apache/thrift => github.com/apache/thrift v0.13.0
