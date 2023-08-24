// Copyright 2023 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package main

import (
	"context"
	"io"
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/cloudwego/kitex-examples/grpcmultiservice/kitex_gen/multiservice"
	"github.com/cloudwego/kitex-examples/grpcmultiservice/kitex_gen/multiservice/serviceb"
	"github.com/cloudwego/kitex-examples/grpcmultiservice/kitex_gen/multiservice/servicex"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/transport"

	"github.com/cloudwego/kitex-examples/grpcmultiservice/kitex_gen/multiservice/servicea"
)

var hostport = "localhost:8888"

func callServiceA() error {
	clia, err := servicea.NewClient("servicea", client.WithTransportProtocol(transport.GRPC), client.WithHostPorts(hostport))
	if err != nil {
		return err
	}

	log.Println("ChatA:「Client」Run Call....")
	stream, err := clia.ChatA(context.Background())
	if err != nil {
		return err
	}
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			req, err := stream.Recv()
			if err == io.EOF {
				log.Println("ChatA:「CLIENT」RECV DONE")
				break
			}
			if err != nil {
				log.Println("ChatA:「CLIENT」Err:", err)
				continue
			}
			log.Println("ChatA:「CLIENT」RECV DATA: " + req.Message)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 2; i++ {
			req := &multiservice.Request{Name: "kitex-" + strconv.Itoa(i)}
			log.Println("ChatA:「CLIENT」SEND DATA: " + req.Name)
			err := stream.Send(req)
			if err != nil {
				log.Println("ChatA:「CLIENT」Err:", err)
			}
			time.Sleep(time.Second)
		}
		err := stream.Close()
		if err != nil {
			log.Println("ChatA:「CLIENT」Err:", err)
		}
		log.Println("ChatA:「CLIENT」SEND DONE")
	}()
	wg.Wait()

	log.Print("ChatA:「Client」Call Done!")
	return nil
}

func callServiceB() error {
	clib, err := serviceb.NewClient("serviceb", client.WithTransportProtocol(transport.GRPC), client.WithHostPorts(hostport))
	if err != nil {
		return err
	}

	log.Println("ChatB:「Client」Run Call....")
	stream, err := clib.ChatB(context.Background())
	if err != nil {
		return err
	}

	for i := 0; i < 10; i++ {
		req := &multiservice.Request{Name: "kitex-" + strconv.Itoa(i)}
		if err := stream.Send(req); err != nil {
			return err
		}
		time.Sleep(time.Second)
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		return err
	}
	log.Printf("message received: %v\n", resp.Message)
	log.Println("ChatB:「Client」Call Done!")
	return nil
}

func callServicX() error {
	clix, err := servicex.NewClient("servicex", client.WithTransportProtocol(transport.GRPC), client.WithHostPorts(hostport))
	if err != nil {
		return err
	}

	log.Println("ChatX:「Client」Run Call....")

	req := &multiservice.RequestX{Name: "kitex-" + strconv.Itoa(0)}
	resp, err := clix.ChatX(context.Background(), req)
	if err != nil {
		return err
	}
	time.Sleep(time.Second)

	log.Printf("message received: %v\n", resp.Message)
	log.Println("ChatX:「Client」Call Done!")
	return nil
}

func main() {
	if err := callServiceA(); err != nil {
		log.Println(err)
	}
	if err := callServiceB(); err != nil {
		log.Println(err)
	}
	if err := callServicX(); err != nil {
		log.Println(err)
	}
}
