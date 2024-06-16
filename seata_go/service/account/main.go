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
	"net"

	account "github.com/cloudwego/kitex-examples/seata_go/kitex_gen/account/accountservice"
	"github.com/cloudwego/kitex-examples/seata_go/middleware"
	"github.com/cloudwego/kitex-examples/seata_go/service/account/dal/model"
	"github.com/cloudwego/kitex-examples/seata_go/service/account/dal/mysql"
	"github.com/cloudwego/kitex/server"

	"seata.apache.org/seata-go/pkg/client"
)

func init() {
	client.InitPath("conf/seatago.yml")
	mysql.Init()
	initData()
}

func main() {
	addr, err := net.ResolveTCPAddr("tcp", ":8881")
	if err != nil {
		panic(err)
	}

	var opts []server.Option
	opts = append(opts, server.WithServiceAddr(addr))
	// use seata-go middleware when initializing the server
	opts = append(opts, server.WithMiddleware(middleware.SeataGoServerMiddleware))

	svr := account.NewServer(new(AccountServiceImpl), opts...)

	err = svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}

func initData() {
	_ = model.Insert(context.Background(), mysql.DB, &model.Account{UserId: "user_1", Money: 10000})
}
