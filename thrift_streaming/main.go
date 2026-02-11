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

	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/endpoint/sep"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/streaming"
	"github.com/cloudwego/kitex/pkg/utils/kitexutil"
	"github.com/cloudwego/kitex/server"

	echo "github.com/cloudwego/kitex-examples/thrift_streaming/kitex_gen/echo/testservice"
)

func main() {
	svr := echo.NewServer(new(TestServiceImpl),
		server.WithMiddleware(func(next endpoint.Endpoint) endpoint.Endpoint {
			// server middleware
			return func(ctx context.Context, req, resp interface{}) (err error) {
				method, _ := kitexutil.GetMethod(ctx)
				klog.Infof("[%s] server middleware, req = %#v", method, req)
				err = next(ctx, req, resp)
				klog.Infof("[%s] server middleware, err = %#v, resp = %#v", method, err, resp)
				return err
			}
		}),

		server.WithStreamOptions(
			// recv middleware
			// NOTE: message (request from client) will NOT be available until `next` returns
			server.WithStreamRecvMiddleware(func(next sep.StreamRecvEndpoint) sep.StreamRecvEndpoint {
				return func(ctx context.Context, stream streaming.ServerStream, req interface{}) (err error) {
					method, _ := kitexutil.GetMethod(ctx)
					err = next(ctx, stream, req)
					klog.Infof("[%s] server recv middleware, err = %#v, req = %#v", method, err, req)
					return err
				}
			}),

			// send middleware
			server.WithStreamSendMiddleware(func(next sep.StreamSendEndpoint) sep.StreamSendEndpoint {
				return func(ctx context.Context, stream streaming.ServerStream, resp interface{}) (err error) {
					method, _ := kitexutil.GetMethod(ctx)
					err = next(ctx, stream, resp)
					klog.Infof("[%s] server send middleware, err = %#v, resp = %#v", method, err, resp)
					return err
				}
			}),
		),
	)

	err := svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
