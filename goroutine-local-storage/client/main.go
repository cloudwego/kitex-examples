// Copyright 2024 CloudWeGo Authors
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
)

var expectedKey = "K_METHOD"

func main() {
	client, err := echo.NewClient("echo", client.WithHostPorts("[::1]:8888"), client.WithContextBackup(func(prev, cur context.Context) (ctx context.Context, backup bool) {
		if v := cur.Value(expectedKey); v != nil {
			// expectedKey exists, no need for recover context
			return cur, false
		}
		// expectedKey doesn't exists, need recover context from prev
		ctx = context.WithValue(cur, expectedKey, prev.Value(expectedKey))
		return ctx, true
	}))
	if err != nil {
		log.Fatal(err)
	}
	for {
		req := &api.Request{Message: "my request"}
		resp, err := client.Echo(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(resp)
		time.Sleep(time.Second)
	}
}
