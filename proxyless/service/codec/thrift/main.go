package main

import (
	"log"

	proxyless "github.com/cloudwego/kitex-proxyless-test/service/codec/thrift/kitex_gen/proxyless/greetservice"
)

func main() {
	svr := proxyless.NewServer(new(GreetServiceImpl))

	err := svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
