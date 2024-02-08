package main

import (
	"log"

	item "example_shop/kitex_gen/example/shop/item/itemservice"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func main() {
	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err)
	}

	svr := item.NewServer(new(ItemServiceImpl),
		server.WithRegistry(r),
		server.WithServerBasicInfo(
			&rpcinfo.EndpointBasicInfo{
				ServiceName: "example.shop.item",
			}),
	)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
