// Copyright 2021 CloudWeGo Authors
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
	"log"
	"time"

	"github.com/cloudwego/kitex-examples/kitex_gen/api"
	"github.com/cloudwego/kitex-examples/kitex_gen/api/echo"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/loadbalance"
)

type ctxKey int

const (
	ctxConsistentKey ctxKey = iota
)

func main() {
	opt := loadbalance.NewConsistentHashOption(func(ctx context.Context, request interface{}) string {
		key, _ := ctx.Value(ctxConsistentKey).(string)
		return key
	})
	lb := loadbalance.NewConsistBalancer(opt)
	client, err := echo.NewClient("echo", client.WithHostPorts("0.0.0.0:8888", "0.0.0.0:8889"), client.WithLoadBalancer(lb))
	if err != nil {
		log.Fatal(err)
	}
	for {
		// call a server
		ctx := context.Background()
		ctx = context.WithValue(ctx, ctxConsistentKey, "my key0")
		req := &api.Request{Message: "my request0"}
		resp, err := client.Echo(ctx, req)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("first call", resp)
		time.Sleep(time.Second)
		// call another server
		ctx = context.Background()
		ctx = context.WithValue(ctx, ctxConsistentKey, "my key1")
		req = &api.Request{Message: "my request1"}
		resp, err = client.Echo(ctx, req)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("second call", resp)
		time.Sleep(time.Second)
	}
}
