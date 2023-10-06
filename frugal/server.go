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

	"github.com/cloudwego/kitex/pkg/remote/codec/thrift"
	"github.com/cloudwego/kitex/server"

	"github.com/cloudwego/kitex-examples/kitex_gen/api"
	"github.com/cloudwego/kitex-examples/kitex_gen/api/echo"
	// If you choose the slim template, make sure the client have `Framed` enabled
	// "github.com/cloudwego/kitex-examples/kitex_gen/slim/api"
	// "github.com/cloudwego/kitex-examples/kitex_gen/slim/api/echo"
)

type EchoImpl struct{}

func (e EchoImpl) Echo(ctx context.Context, req *api.Request) (r *api.Response, err error) {
	return &api.Response{Message: req.Message}, nil
}

func frugalServer() {
	code := thrift.NewThriftCodecWithConfig(thrift.FrugalRead | thrift.FrugalWrite)
	svr := echo.NewServer(new(EchoImpl), server.WithPayloadCodec(code))
	err := svr.Run()
	if err != nil {
		panic(err)
	}
}
