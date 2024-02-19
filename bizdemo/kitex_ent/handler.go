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

	"github.com/cloudwego/kitex-examples/bizdemo/kitex_ent/dao/mysql"
	user "github.com/cloudwego/kitex-examples/bizdemo/kitex_ent/kitex_gen/user"
	"github.com/cloudwego/kitex-examples/bizdemo/kitex_ent/pack"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// UpdateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdateUser(ctx context.Context, req *user.UpdateUserRequest) (resp *user.UpdateUserResponse, err error) {
	resp = new(user.UpdateUserResponse)

	if err = mysql.UpdateUser(mysql.Client, int(req.UserId), req.Name, int(req.Age)); err != nil {
		resp.Msg = err.Error()
		resp.Code = user.Code_DBErr
		return
	}

	resp.Code = user.Code_Success
	return
}

// DeleteUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) DeleteUser(ctx context.Context, req *user.DeleteUserRequest) (resp *user.DeleteUserResponse, err error) {
	resp = new(user.DeleteUserResponse)

	if err = mysql.DeleteUser(mysql.Client, int(req.UserId)); err != nil {
		resp.Msg = err.Error()
		resp.Code = user.Code_DBErr
		return
	}

	resp.Code = user.Code_Success
	return
}

// QueryUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) QueryUser(ctx context.Context, req *user.QueryUserRequest) (resp *user.QueryUserResponse, err error) {
	resp = new(user.QueryUserResponse)

	users, err := mysql.SearchUsers(mysql.Client, *req.Keyword, int(req.Page), int(req.PageSize))
	if err != nil {
		resp.Msg = err.Error()
		resp.Code = user.Code_DBErr
	}

	resp.Totoal = int64(len(users))
	resp.Users = pack.Users(users)
	resp.Code = user.Code_Success

	return
}

// CreateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CreateUser(ctx context.Context, req *user.CreateUserRequest) (resp *user.CreateUserResponse, err error) {
	resp = new(user.CreateUserResponse)

	if _, err = mysql.CreateUser(mysql.Client, "John Doe", 30); err != nil {
		resp.Msg = err.Error()
		resp.Code = user.Code_DBErr
	}

	resp.Code = user.Code_Success
	return
}
