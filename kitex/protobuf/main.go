package main

import (
	hello "kitex-examples/kitex/protobuf/kitex_gen/hello/helloservice"
	"log"
)

func main() {
	svr := hello.NewServer(new(HelloServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
