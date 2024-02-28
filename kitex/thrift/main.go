package main

import (
	"log"

	example "kitex-examples/kitex/thrift/kitex_gen/hello/example/helloservice"
)

func main() {
	svr := example.NewServer(new(HelloServiceImpl))

	err := svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
