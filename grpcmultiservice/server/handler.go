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
	"fmt"
	"io"
	"log"
	"strings"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"golang.org/x/sync/errgroup"

	"github.com/cloudwego/kitex-examples/grpcmultiservice/kitex_gen/multiservice"
)

type ServiceAImpl struct{}

func (s *ServiceAImpl) ChatA(stream multiservice.ServiceA_ChatAServer) (err error) {
	log.Println("ChatA called")
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
		resp := &multiservice.Reply{}
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
	log.Print("ChatA call Done!")
	return eg.Wait()
}

type ServiceBImpl struct{}

func (s *ServiceBImpl) ChatB(stream multiservice.ServiceB_ChatBServer) (err error) {
	log.Println("ChatB called")
	var msgs []string
	for {
		req, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		klog.Infof("message received: %v\n", req.Name)
		msgs = append(msgs, req.Name)
		time.Sleep(time.Second)
	}
	return stream.SendAndClose(&multiservice.Reply{Message: "all message: " + strings.Join(msgs, ", ")})
}

type ServiceXImpl struct{}

func (s *ServiceXImpl) ChatX(ctx context.Context, req *multiservice.RequestX) (res *multiservice.ReplyX, err error) {
	log.Println("ChatX called")
	return &multiservice.ReplyX{Message: req.Name}, nil
}
