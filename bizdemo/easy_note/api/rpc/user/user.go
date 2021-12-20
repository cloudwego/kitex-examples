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

package user

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/api/constant"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"

	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/api/errno"
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/api/kitex_gen/kitex/demo/user"
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/api/kitex_gen/kitex/demo/user/userservice"
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/api/middleware"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
)

var (
	userClient userservice.Client
)

func initJaeger(service string) (client.Suite, io.Closer) {
	cfg, _ := jaegercfg.FromEnv()
	cfg.ServiceName = service
	tracer, closer, err := cfg.NewTracer(jaegercfg.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	opentracing.InitGlobalTracer(tracer)
	return trace.NewDefaultClientSuite(), closer
}

func Init() {

	tracer, _ := initJaeger(constant.ServiceName)

	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		panic(err)
	}

	fp := retry.NewFailurePolicy()

	c, err := userservice.NewClient(
		constant.UserServiceName,
		client.WithMiddleware(middleware.CommonMiddleware),
		client.WithMiddleware(middleware.ClientMiddleware),
		client.WithMuxConnection(1),                    //mux
		client.WithRPCTimeout(3*time.Second),           // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond), // conn timeout
		client.WithFailureRetry(fp),                    // retry
		client.WithSuite(tracer),                       // tracer
		client.WithResolver(r),
	)
	if err != nil {
		panic(err)
	}

	userClient = c
}

func CreateUser(ctx context.Context, req *user.CreateUserRequest) error {
	resp, err := userClient.CreateUser(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrno(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}

	return nil
}

func CheckUser(ctx context.Context, req *user.CheckUserRequest) (int64, error) {
	resp, err := userClient.CheckUser(ctx, req)
	if err != nil {
		return 0, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return 0, errno.NewErrno(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return resp.UserId, nil
}

func MGetUser(ctx context.Context, req *user.MGetUserRequest) ([]*user.User, error) {
	resp, err := userClient.MGetUser(ctx, req)
	if err != nil {
		return nil, err
	}

	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrno(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}

	return resp.Users, nil
}
