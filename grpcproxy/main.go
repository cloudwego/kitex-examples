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

package main

import (
	"time"

	"github.com/cloudwego/kitex-examples/grpcproxy/client"
	"github.com/cloudwego/kitex-examples/grpcproxy/proxy"
	"github.com/cloudwego/kitex-examples/grpcproxy/proxy/handler"
	"github.com/cloudwego/kitex-examples/grpcproxy/server"
)

func main() {
	// Do proxy by redirecting gRPC frame
	handler := handler.GRPCFrameProxyHandler

	// Do proxy by decoding gRPC frame into struct and send them to target server.
	// handler := handler.GRPCStructProxyHandler

	go server.RunServer(proxy.Ip)
	go proxy.RunProxyServer(proxy.ProxyIp, handler)
	time.Sleep(time.Second)
	client.RunClient(proxy.ProxyIp)
}
