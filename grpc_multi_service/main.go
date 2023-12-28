package main

import (
	"github.com/cloudwego/kitex/server"
	"grpc_multi_service/kitex_gen/multi/service/servicea"
	"grpc_multi_service/kitex_gen/multi/service/serviceb"
	"log"
	"net"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", "localhost:8888")

	svr := server.NewServer(server.WithServiceAddr(addr))
	err := svr.RegisterService(servicea.NewServiceInfo(), new(ServiceAImpl))
	if err != nil {
		log.Println(err.Error())
	}
	err = svr.RegisterService(serviceb.NewServiceInfo(), new(ServiceBImpl))
	if err != nil {
		log.Println(err.Error())
	}
	err = svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
