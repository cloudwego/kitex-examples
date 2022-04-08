module github.com/cloudwego/kitex-examples/bizdemo/mall

go 1.16

require (
	github.com/apache/thrift v0.16.0
	github.com/appleboy/gin-jwt/v2 v2.8.0
	github.com/cloudwego/kitex v0.2.1
	github.com/cloudwego/thriftgo v0.1.7
	github.com/gin-gonic/gin v1.7.7
	github.com/kitex-contrib/registry-etcd v0.0.0-20220110034026-b1c94979cea3
	github.com/swaggo/files v0.0.0-20210815190702-a29dd2bc99b2 // indirect
	github.com/swaggo/gin-swagger v1.4.1 // indirect
	github.com/swaggo/swag v1.8.1 // indirect
	gorm.io/driver/mysql v1.3.3
	gorm.io/gorm v1.23.4
)

replace github.com/apache/thrift => github.com/apache/thrift v0.13.0
