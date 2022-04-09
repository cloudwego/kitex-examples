package main

import (
	"github.com/cloudwego/kitex-examples/bizdemo/mall/business/mall_shop/dal"
	shop "github.com/cloudwego/kitex-examples/bizdemo/mall/business/mall_shop/kitex_gen/cmp/ecom/shop/shopservice"
	"github.com/cloudwego/kitex-examples/bizdemo/mall/pkg/conf"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"net"
)

func Init() {
	dal.Init()
}

func main() {
	r, err := etcd.NewEtcdRegistry([]string{conf.EtcdAddress})
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8890")
	if err != nil {
		panic(err)
	}
	Init()
	svr := shop.NewServer(new(ShopServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.ShopRpcServiceName}), // server name
		server.WithServiceAddr(addr), // address
		server.WithRegistry(r),       // registry
	)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
