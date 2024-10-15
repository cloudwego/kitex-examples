/*
 * Copyright 2024 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"log"

	"github.com/cloudwego/kitex-examples/bizdemo/kitex_swagger_gen/dao/mysql"
	user "github.com/cloudwego/kitex-examples/bizdemo/kitex_swagger_gen/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex-examples/bizdemo/kitex_swagger_gen/swagger"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/server"
	"github.com/cloudwego/kitex/transport"
)

var svr2Cli user.Client

func main() {
	mysql.Init()

	var err error
	endpoints := []string{"127.0.0.1:8889"}

	svr2Cli, err = user.NewClient("userservice", client.WithHostPorts(endpoints...), client.WithTransportProtocol(transport.TTHeader))
	if err != nil {
		log.Fatal(err)
	}

	svr := user.NewServer(new(UserServiceImpl), server.WithTransHandlerFactory(&swagger.MixTransHandlerFactory{}))

	err = svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
