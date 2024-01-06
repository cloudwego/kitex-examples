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
	"github.com/cloudwego/kitex/server"
	"grpc_multi_service/kitex_gen/multi/service/servicea"
	"grpc_multi_service/kitex_gen/multi/service/serviceb"
	"log"
	"net"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", "localhost:8888")

	svr := server.NewServer(server.WithServiceAddr(addr))
	err := svr.RegisterService(servicea.NewServiceInfo(), new(ServiceAImpl))
	if err != nil {
		log.Println(err.Error())
	}
	err = svr.RegisterService(serviceb.NewServiceInfo(), new(ServiceBImpl))
	if err != nil {
		log.Println(err.Error())
	}
	err = svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
