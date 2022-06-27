// Copyright 2022 CloudWeGo Authors
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

package client

import (
	"context"
	"io"
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/cloudwego/kitex-examples/grpcproxy/kitex_gen/grpcproxy"
	"github.com/cloudwego/kitex-examples/grpcproxy/kitex_gen/grpcproxy/servicea"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/transport"
)

func RunClient(hostport string) {
	client, _ := servicea.NewClient("grpcService", client.WithTransportProtocol(transport.GRPC), client.WithHostPorts(hostport))

	runBidiStream(client)
}

func runBidiStream(client servicea.Client) {
	log.Println("「Client」Run Bidi Stream Call....")

	stream, err := client.Chat(context.Background())
	if err != nil {
		log.Println(err)
	}
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			req, err := stream.Recv()
			if err == io.EOF {
				log.Println("「CLIENT」RECV DONE")
				break
			}
			if err != nil {
				log.Println("「CLIENT」Err:", err)
				continue
			}
			log.Println("「CLIENT」RECV DATA: " + req.Message)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 2; i++ {
			req := &grpcproxy.Request{Name: "kitex-" + strconv.Itoa(i)}
			log.Println("「CLIENT」SEND DATA: " + req.Name)
			err := stream.Send(req)
			if err != nil {
				log.Println("「CLIENT」Err:", err)
			}
			time.Sleep(time.Second)
		}
		err := stream.Close()
		if err != nil {
			log.Println("「CLIENT」Err:", err)
			return
		}
		log.Println("「CLIENT」SEND DONE")
	}()
	wg.Wait()

	log.Print("「Client」Bidi Call Done!")
}
