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
	"context"
	"log"

	"github.com/bytedance/gopkg/cloud/metainfo"
	"github.com/cloudwego/kitex-examples/bizdemo/kitex_swagger_gen/kitex_gen/user"
	"github.com/cloudwego/kitex-examples/bizdemo/kitex_swagger_gen/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/pkg/utils"
	"github.com/cloudwego/kitex/server"
)

// UserService2Impl implements the last service interface defined in the IDL.
type UserService2Impl struct{}

// UpdateUser implements the UserServiceImpl interface.
func (s *UserService2Impl) UpdateUser(ctx context.Context, req *user.UpdateUserRequest) (resp *user.UpdateUserResponse, err error) {
	t1, _ := metainfo.GetPersistentValue(ctx, "p_k")
	t2, _ := metainfo.GetValue(ctx, "k")

	return &user.UpdateUserResponse{
		Code: user.Code_Success,
		Msg:  " 持续元信息->" + t1 + " 临时元信息->" + t2,
	}, nil
}

// DeleteUser implements the UserServiceImpl interface.
func (s *UserService2Impl) DeleteUser(ctx context.Context, req *user.DeleteUserRequest) (resp *user.DeleteUserResponse, err error) {
	t1, _ := metainfo.GetPersistentValue(ctx, "p_k")
	t2, _ := metainfo.GetValue(ctx, "k")

	return &user.DeleteUserResponse{
		Code: user.Code_Success,
		Msg:  " 持续元信息->" + t1 + " 临时元信息->" + t2,
	}, nil
}

// QueryUser implements the UserServiceImpl interface.
func (s *UserService2Impl) QueryUser(ctx context.Context, req *user.QueryUserRequest) (resp *user.QueryUserResponse, err error) {
	t1, _ := metainfo.GetPersistentValue(ctx, "p_k")
	t2, _ := metainfo.GetValue(ctx, "k")

	return &user.QueryUserResponse{
		Code: user.Code_Success,
		Msg:  " 持续元信息->" + t1 + " 临时元信息->" + t2,
	}, nil
}

// CreateUser implements the UserServiceImpl interface.
func (s *UserService2Impl) CreateUser(ctx context.Context, req *user.CreateUserRequest) (resp *user.CreateUserResponse, err error) {
	t1, _ := metainfo.GetPersistentValue(ctx, "p_k")
	t2, _ := metainfo.GetValue(ctx, "k")

	return &user.CreateUserResponse{
		Code: user.Code_Success,
		Msg:  " 持续元信息->" + t1 + " 临时元信息->" + t2,
	}, nil
}

func main() {
	svr := userservice.NewServer(new(UserService2Impl), server.WithServiceAddr(utils.NewNetAddr("tcp", ":8889")))

	err := svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
