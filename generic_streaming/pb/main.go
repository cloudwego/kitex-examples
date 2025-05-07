package main

import (
	"log"
	pb "pb_generic_streaming_demo/kitex_gen/pb/streamingservice"
)

func main() {
	svr := pb.NewServer(new(StreamingServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
