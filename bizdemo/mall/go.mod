module github.com/cloudwego/kitex-examples/bizdemo/mall

go 1.16

require (
	github.com/apache/thrift v0.16.0
	github.com/cloudwego/hertz v0.1.0
	github.com/cloudwego/kitex v0.2.1
	github.com/cloudwego/thriftgo v0.1.7
	github.com/hertz-contrib/jwt v1.0.0
	github.com/hertz-contrib/swagger v0.0.0-20220622161205-642b89a21ef2
	github.com/kitex-contrib/registry-etcd v0.0.0-20220110034026-b1c94979cea3
	github.com/satori/go.uuid v1.2.0
	github.com/swaggo/files v0.0.0-20210815190702-a29dd2bc99b2
	github.com/swaggo/swag v1.8.2
	gorm.io/driver/mysql v1.3.3
	gorm.io/gorm v1.23.4
)

replace github.com/apache/thrift => github.com/apache/thrift v0.13.0
