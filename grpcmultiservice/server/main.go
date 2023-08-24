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
	"log"
	"net"

	"github.com/cloudwego/kitex/pkg/serviceinfo"
	"github.com/cloudwego/kitex/server"

	"github.com/cloudwego/kitex-examples/grpcmultiservice/kitex_gen/multiservice"
	"github.com/cloudwego/kitex-examples/grpcmultiservice/kitex_gen/multiservice/servicea"
	"github.com/cloudwego/kitex-examples/grpcmultiservice/kitex_gen/multiservice/serviceb"
	"github.com/cloudwego/kitex-examples/grpcmultiservice/kitex_gen/multiservice/servicex"
)

func main() {
	var svcs []serviceinfo.Service
	svca := servicea.BuildServiceAService(new(ServiceAImpl))
	svcb := serviceb.BuildServiceBService(new(ServiceBImpl))
	svcx := servicex.BuildServiceXService(new(ServiceXImpl))
	svcs = append(svcs, svca, svcb, svcx)

	addr, _ := net.ResolveTCPAddr("tcp", "localhost:8888")
	svr := multiservice.NewServerWithMultiServices(svcs, server.WithServiceAddr(addr))
	err := svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
