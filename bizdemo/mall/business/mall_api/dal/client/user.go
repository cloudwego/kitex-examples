// Copyright 2022 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package client

import (
	"context"
	"time"

	"github.com/cloudwego/kitex-examples/bizdemo/mall/business/mall_api/kitex_gen/cmp/ecom/user"
	"github.com/cloudwego/kitex-examples/bizdemo/mall/business/mall_api/kitex_gen/cmp/ecom/user/userservice"
	"github.com/cloudwego/kitex-examples/bizdemo/mall/pkg/conf"
	"github.com/cloudwego/kitex-examples/bizdemo/mall/pkg/errno"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var userClient userservice.Client

func initUserRpc() {
	r, err := etcd.NewEtcdResolver([]string{conf.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := userservice.NewClient(
		conf.UserRpcServiceName,
		client.WithRPCTimeout(3*time.Second),              // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond),    // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithResolver(r),                            // resolver
	)
	if err != nil {
		panic(err)
	}
	userClient = c
}

func CreateUser(ctx context.Context, req *user.CreateUserReq) error {
	resp, err := userClient.CreateUser(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(int64(resp.BaseResp.StatusCode), resp.BaseResp.StatusMessage)
	}
	return nil
}

func CheckUser(ctx context.Context, req *user.CheckUserReq) (int64, error) {
	resp, err := userClient.CheckUser(ctx, req)
	if err != nil {
		return 0, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return 0, errno.NewErrNo(int64(resp.BaseResp.StatusCode), resp.BaseResp.StatusMessage)
	}
	return resp.UserId, nil
}
