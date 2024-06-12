package main

import (
	order "github.com/cloudwego/kitex-examples/seata_go/kitex_gen/order/orderservice"
	"github.com/cloudwego/kitex-examples/seata_go/middleware"
	"github.com/cloudwego/kitex-examples/seata_go/service/order/dal/mysql"
	"github.com/cloudwego/kitex-examples/seata_go/service/order/rpc"
	"github.com/cloudwego/kitex/server"
	"log"
	"net"
	"seata.apache.org/seata-go/pkg/client"
)

func init() {
	client.InitPath("conf/seatago.yml")
	mysql.Init()
	rpc.InitClient()
}

func main() {
	addr, err := net.ResolveTCPAddr("tcp", ":8883")
	if err != nil {
		panic(err)
	}

	var opts []server.Option
	opts = append(opts, server.WithServiceAddr(addr))
	// use seata-go middleware when initializing the server
	opts = append(opts, server.WithMiddleware(middleware.SeataGoServerMiddleware))

	svr := order.NewServer(new(OrderServiceImpl), opts...)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
