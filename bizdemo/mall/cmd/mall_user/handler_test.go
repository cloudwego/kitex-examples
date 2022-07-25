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

package main

import (
	"context"
	"testing"

	"github.com/cloudwego/kitex-examples/bizdemo/mall/cmd/mall_user/kitex_gen/cmp/ecom/user"
	"github.com/cloudwego/thriftgo/pkg/test"
)

func TestUserServiceImpl_CreateUser(t *testing.T) {
	Init()
	ctx := context.TODO()
	impl := new(UserServiceImpl)
	resp, err := impl.CreateUser(ctx, &user.CreateUserReq{
		UserName: "yaoxianjie",
		Password: "yxj12345",
	})
	test.Assert(t, err == nil)
	test.Assert(t, resp.BaseResp.StatusCode == 0)
}

func TestUserServiceImpl_CheckUser(t *testing.T) {
	Init()
	ctx := context.TODO()
	impl := new(UserServiceImpl)
	resp, err := impl.CheckUser(ctx, &user.CheckUserReq{
		UserName: "yaoxianjie",
		Password: "yxj12345",
	})
	test.Assert(t, err == nil)
	test.Assert(t, resp.BaseResp.StatusCode == 0)
}

func TestUserServiceImpl_MGetUser(t *testing.T) {
	Init()
	ctx := context.TODO()
	impl := new(UserServiceImpl)
	ids := make([]int64, 0)
	ids = append(ids, int64(1))
	resp, err := impl.MGetUser(ctx, &user.MGetUserReq{Ids: ids})
	test.Assert(t, err == nil)
	test.Assert(t, resp.BaseResp.StatusCode == 0)
	test.Assert(t, resp.Users[0].UserName == "yaoxianjie")
}

func TestUserServiceImpl_AddUserRole(t *testing.T) {
	Init()
	ctx := context.TODO()
	impl := new(UserServiceImpl)
	userName := "yaoxianjie"
	role := user.Role_Admin
	resp, err := impl.AddUserRole(ctx, &user.AddUserRoleReq{
		UserName: userName,
		Role:     role,
	})
	test.Assert(t, err == nil)
	test.Assert(t, resp.BaseResp.StatusCode == 0)

	role = user.Role_ProductAuditor
	resp, err = impl.AddUserRole(ctx, &user.AddUserRoleReq{
		UserName: userName,
		Role:     role,
	})
	test.Assert(t, err == nil)
	test.Assert(t, resp.BaseResp.StatusCode == 0)
}

func TestUserServiceImpl_DelUserRole(t *testing.T) {
	Init()
	ctx := context.TODO()
	impl := new(UserServiceImpl)
	userName := "yaoxianjie"
	role := user.Role_ProductAuditor
	resp, err := impl.DelUserRole(ctx, &user.DelUserRoleReq{
		UserName: userName,
		Role:     role,
	})
	test.Assert(t, err == nil)
	test.Assert(t, resp.BaseResp.StatusCode == 0)

	role = user.Role_CategoryOperator
	resp, err = impl.DelUserRole(ctx, &user.DelUserRoleReq{
		UserName: userName,
		Role:     role,
	})
	test.Assert(t, err == nil)
	test.Assert(t, resp.BaseResp.StatusCode == 0)
}

func TestUserServiceImpl_ValidateUserRole(t *testing.T) {
	Init()
	ctx := context.TODO()
	impl := new(UserServiceImpl)
	userName := "admin"
	role := []user.Role{user.Role_CategoryOperator}
	resp, err := impl.ValidateUserRole(ctx, &user.ValidateUserRolesReq{
		UserName: userName,
		Roles:    role,
	})
	test.Assert(t, err == nil)
	test.Assert(t, resp.BaseResp.StatusCode == 0)
	test.Assert(t, resp.IsPass == true)
}
