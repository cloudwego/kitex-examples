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
	"log"
	"net"

	order "github.com/cloudwego/kitex-examples/seata_go/kitex_gen/order/orderservice"
	"github.com/cloudwego/kitex-examples/seata_go/middleware"
	"github.com/cloudwego/kitex-examples/seata_go/service/order/dal/mysql"
	"github.com/cloudwego/kitex-examples/seata_go/service/order/rpc"
	"github.com/cloudwego/kitex/server"

	"seata.apache.org/seata-go/pkg/client"
)

func init() {
	client.InitPath("conf/seatago.yml")
	mysql.Init()
	rpc.InitClient()
}

func main() {
	addr, err := net.ResolveTCPAddr("tcp", ":8883")
	if err != nil {
		panic(err)
	}

	var opts []server.Option
	opts = append(opts, server.WithServiceAddr(addr))
	// use seata-go middleware when initializing the server
	opts = append(opts, server.WithMiddleware(middleware.SeataGoServerMiddleware))

	svr := order.NewServer(new(OrderServiceImpl), opts...)

	err = svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
