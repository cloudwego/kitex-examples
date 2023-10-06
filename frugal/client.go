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
	"context"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/remote/codec/thrift"
	"github.com/cloudwego/kitex/transport"

	"github.com/cloudwego/kitex-examples/kitex_gen/api"
	"github.com/cloudwego/kitex-examples/kitex_gen/api/echo"
)

func frugalClient() {
	codec := thrift.NewThriftCodecWithConfig(thrift.FrugalRead | thrift.FrugalWrite)
	framed := client.WithTransportProtocol(transport.Framed)
	server := client.WithHostPorts("127.0.0.1:8888")
	cli := echo.MustNewClient("a.b.c", server, client.WithPayloadCodec(codec), framed)
	rsp, err := cli.Echo(context.Background(), &api.Request{Message: "Hello"})
	klog.Infof("resp: %v, err: %v", rsp, err)
}
