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

package server

import (
	"fmt"
	"io"
	"log"
	"time"

	"github.com/cloudwego/kitex-examples/grpcproxy/kitex_gen/grpcproxy"
	"golang.org/x/sync/errgroup"
)

type ServiceAImpl struct{}

func (s *ServiceAImpl) Chat(stream grpcproxy.ServiceA_ChatServer) (err error) {
	log.Println("「SERVER」Handle Bidi Stream Call!")
	var eg errgroup.Group
	eg.Go(func() error {
		for {
			req, err := stream.Recv()
			if err != nil {
				if err == io.EOF {
					return nil
				}
				log.Println("「SERVER」Err:", err)
				return err
			}
			log.Println("「SERVER」Recv Data:" + req.Name)
			time.Sleep(time.Second)
		}
	})
	eg.Go(func() error {
		resp := &grpcproxy.Reply{}
		var cnt int
		for cnt < 2 {
			resp.Message = fmt.Sprintf("%dth response", cnt)
			log.Println("「SERVER」Send Data:" + resp.Message)
			err := stream.Send(resp)
			if err != nil {
				if err != io.EOF {
					log.Println("「SERVER」Err:", err)
				}
				return err
			}
			cnt++
			time.Sleep(time.Second)
		}
		return nil
	})
	log.Print("「Server」Bidi Call Done!")
	return eg.Wait()
}
