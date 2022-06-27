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

package handler

import (
	"context"
	"io"
	"log"

	"github.com/cloudwego/kitex-examples/grpcproxy/kitex_gen/grpcproxy"
	"github.com/cloudwego/kitex-examples/grpcproxy/kitex_gen/grpcproxy/servicea"
	"github.com/cloudwego/kitex-examples/grpcproxy/proxy"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/streaming"
	"github.com/cloudwego/kitex/transport"
)

// GRPCStructProxyHandler redirects RPC struct from client to target server.
func GRPCStructProxyHandler(ctx context.Context, methodName string, serverStream streaming.Stream) error {
	log.Println("Proxy Handler is working....")

	_, address := proxy.Resolve(methodName)
	client, _ := servicea.NewClient("destService", client.WithHostPorts(address),
		client.WithTransportProtocol(transport.GRPC))

	clientStream, err := client.Chat(context.Background())
	if err != nil {
		return err
	}

	s2c := redirectStruct(serverStream, clientStream)
	c2s := redirectStruct(clientStream, serverStream)

	for i := 0; i < 2; i++ {
		select {
		case s2cErr := <-s2c:
			if s2cErr != io.EOF {
				return s2cErr
			}
		case c2sErr := <-c2s:
			if c2sErr != io.EOF {
				return c2sErr
			}
		}
	}
	return nil
}

func redirectStruct(from, to streaming.Stream) chan error {
	ret := make(chan error)

	go func() {
		for {
			req := &grpcproxy.Request{}
			err := from.RecvMsg(req)
			if err != nil {
				from.Close()
				ret <- err
				break
			}

			// do your own filter logic here
			//if req.Name==xxx{
			//	continue
			//}

			err = to.SendMsg(req)
			if err != nil {
				ret <- err
				break
			}
		}
	}()

	return ret
}
