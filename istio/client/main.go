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
	"os"
	"time"

	"github.com/cloudwego/kitex-examples/istio/kitex_gen/hello"
	api "github.com/cloudwego/kitex-examples/istio/kitex_gen/hello/hello"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/transport"
)

const (
	serviceNameKey = "SERVICE_NAME"
)

func servicehost() string {
	// the Kubernetes Service Name, if the Service to be accessed
	// is not in the same namespace, the Namespace should be added
	// as the suffix, format: <ServiceName>.<Namespace>
	serviceName := os.Getenv(serviceNameKey)
	return serviceName + ":8888"
}

func main() {
	client, err := api.NewClient("hello",
		client.WithHostPorts(servicehost()),
		// should use grpc protocol, thrift is not well compatible in istio.
		client.WithTransportProtocol(transport.GRPC),
	)
	if err != nil {
		log.Fatal(err)
	}
	for {
		req := &hello.Request{Message: "my request"}
		resp, err := client.Echo(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(resp)
		time.Sleep(time.Second)
	}
}
