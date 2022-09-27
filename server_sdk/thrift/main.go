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
	"context"
	"net"

	"github.com/cloudwego/kitex-examples/kitex_gen/api"
	"github.com/cloudwego/kitex-examples/kitex_gen/api/echo"

	"github.com/apache/thrift/lib/go/thrift"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/utils"
	"github.com/cloudwego/kitex/server/invoke"
)

var _ api.Echo = &EchoImpl{}

// EchoImpl implements the last service interface defined in the IDL.
type EchoImpl struct{}

// Echo implements the Echo interface.
func (s *EchoImpl) Echo(ctx context.Context, req *api.Request) (resp *api.Response, err error) {
	return &api.Response{Message: req.Message}, nil
}

func main() {
	var reqPayload, respPayload []byte
	var local, remote net.Addr
	// init local/remote
	local = utils.NewNetAddr("tcp", "127.0.0.1:8889")
	remote = utils.NewNetAddr("tcp", "127.0.0.1:8888")
	ivk := echo.NewInvoker(new(EchoImpl))
	msg := invoke.NewMessage(local, remote)

	// setup request payload
	codec := utils.NewThriftMessageCodec()
	args := api.NewEchoEchoArgs()
	args.SetReq(&api.Request{
		Message: "hello",
	})
	reqPayload, err := codec.Encode("echo", thrift.CALL, 0, args)
	if err != nil {
		klog.Fatal(err)
	}

	if err = msg.SetRequestBytes(reqPayload); err != nil {
		klog.Fatal(err)
	}

	// start a call
	if err = ivk.Call(msg); err != nil {
		klog.Fatal(err)
	}

	// get response
	respPayload, err = msg.GetResponseBytes()
	if err != nil {
		klog.Fatal(err)
	}

	res := api.NewEchoEchoResult()
	method, _, err := codec.Decode(respPayload, res)
	if err != nil {
		klog.Fatal(err)
	}

	klog.Infof("method: %s, res: %#v", method, res)
}
