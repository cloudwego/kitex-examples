package main

import (
	"context"
	account "github.com/cloudwego/kitex-examples/seata_go/kitex_gen/account/accountservice"
	"github.com/cloudwego/kitex-examples/seata_go/middleware"
	"github.com/cloudwego/kitex-examples/seata_go/service/account/dal/model"
	"github.com/cloudwego/kitex-examples/seata_go/service/account/dal/mysql"
	"github.com/cloudwego/kitex/server"
	"log"
	"net"
	"seata.apache.org/seata-go/pkg/client"
)

func init() {
	client.InitPath("conf/seatago.yml")
	mysql.Init()
	initData()
}

func main() {
	addr, err := net.ResolveTCPAddr("tcp", ":8881")
	if err != nil {
		panic(err)
	}

	var opts []server.Option
	opts = append(opts, server.WithServiceAddr(addr))
	// use seata-go middleware when initializing the server
	opts = append(opts, server.WithMiddleware(middleware.SeataGoServerMiddleware))

	svr := account.NewServer(new(AccountServiceImpl), opts...)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}

func initData() {
	_ = model.Insert(context.Background(), mysql.DB, &model.Account{UserId: "user_1", Money: 10000})
}
