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

	"github.com/cloudwego/kitex-examples/kitex_gen/pbapi"
	"github.com/cloudwego/kitex-examples/kitex_gen/pbapi/echo"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/transport"
)

func main() {
	cli := echo.MustNewClient("echo", client.WithTransportProtocol(transport.GRPC),
		client.WithHostPorts("[::1]:8888"))

	req := &pbapi.Request{Message: "my request"}
	resp, err := cli.Echo(context.Background(), req)
	bizErr, isBizErr := kerrors.FromBizStatusError(err)

	if isBizErr {
		log.Println(bizErr.BizMessage())
	} else {
		log.Println(resp)
	}
}
