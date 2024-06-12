package rpc

import (
	"github.com/cloudwego/kitex-examples/seata_go/kitex_gen/account/accountservice"
	"github.com/cloudwego/kitex-examples/seata_go/kitex_gen/storage/storageservice"
	"github.com/cloudwego/kitex-examples/seata_go/middleware"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/transport"
)

var (
	accountServerAddr = "localhost:8881"
	storageServerAddr = "localhost:8882"
	AccountClient     accountservice.Client
	StorageClient     storageservice.Client
)

func InitClient() {
	AccountClient = accountservice.MustNewClient("account",
		client.WithHostPorts(accountServerAddr),
		client.WithMiddleware(middleware.SeataGoClientMiddleware),
		client.WithTransportProtocol(transport.TTHeaderFramed))
	StorageClient = storageservice.MustNewClient("storage",
		client.WithHostPorts(storageServerAddr),
		client.WithMiddleware(middleware.SeataGoClientMiddleware),
		client.WithTransportProtocol(transport.TTHeaderFramed))
}
