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

	"github.com/cloudwego/kitex/client"

	"github.com/cloudwego/kitex-examples/thrift_multi_service/kitex_gen/api"
	"github.com/cloudwego/kitex-examples/thrift_multi_service/kitex_gen/api/servicea"
	"github.com/cloudwego/kitex-examples/thrift_multi_service/kitex_gen/api/serviceb"
)

func main() {
	clientA, err := servicea.NewClient("servicea", client.WithHostPorts("[::1]:8888"))
	if err != nil {
		log.Fatal(err)
	}
	req := &api.Request{Message: "my request"}
	resp, err := clientA.EchoA(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)

	clientB, err := serviceb.NewClient("serviceb", client.WithHostPorts("[::1]:8888"))
	if err != nil {
		log.Fatal(err)
	}
	resp, err = clientB.EchoB(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)
}
