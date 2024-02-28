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
	"strconv"

	"github.com/cloudwego/kitex-examples/bizdemo/kitex_gorm_gen/model/model"

	user "github.com/cloudwego/kitex-examples/bizdemo/kitex_gorm_gen/kitex_gen/user"
	"github.com/cloudwego/kitex-examples/bizdemo/kitex_gorm_gen/model/query"
	"github.com/cloudwego/kitex-examples/bizdemo/kitex_gorm_gen/pack"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// UpdateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdateUser(ctx context.Context, req *user.UpdateUserRequest) (resp *user.UpdateUserResponse, err error) {
	resp = new(user.UpdateUserResponse)
	u := &model.User{
		Name:      req.Name,
		Gender:    strconv.FormatInt(int64(req.Gender), 10),
		Age:       int32(req.Age),
		Introduce: req.Introduce,
	}

	_, err = query.User.WithContext(ctx).Updates(u)
	if err != nil {
		resp.Code = user.Code_DBErr
		resp.Msg = err.Error()
		return
	}

	resp.Code = user.Code_Success
	return
}

// DeleteUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) DeleteUser(ctx context.Context, req *user.DeleteUserRequest) (resp *user.DeleteUserResponse, err error) {
	resp = new(user.DeleteUserResponse)

	_, err = query.User.WithContext(ctx).Where(query.User.ID.Eq(int32(req.UserId))).Delete()
	if err != nil {
		resp.Code = user.Code_DBErr
		resp.Msg = err.Error()
		return
	}

	resp.Code = user.Code_Success
	return
}

// QueryUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) QueryUser(ctx context.Context, req *user.QueryUserRequest) (resp *user.QueryUserResponse, err error) {
	resp = new(user.QueryUserResponse)

	u, m := query.User, query.User.WithContext(ctx)
	if *req.Keyword != "" {
		m = m.Where(u.Introduce.Like("%" + *req.Keyword + "%"))
	}

	var total int64
	total, err = m.Count()
	if err != nil {
		resp.Code = user.Code_DBErr
		resp.Msg = err.Error()
		return
	}

	var users []*model.User
	if total > 0 {
		users, err = m.Limit(int(req.PageSize)).Offset(int(req.PageSize * (req.Page - 1))).Find()
		if err != nil {
			resp.Code = user.Code_DBErr
			resp.Msg = err.Error()
			return
		}
	}

	resp.Code = user.Code_Success
	resp.Totoal = total
	resp.Users = pack.Users(users)

	return
}

// CreateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CreateUser(ctx context.Context, req *user.CreateUserRequest) (resp *user.CreateUserResponse, err error) {
	resp = new(user.CreateUserResponse)

	err = query.User.WithContext(ctx).Create(&model.User{
		Name:      req.Name,
		Gender:    strconv.FormatInt(int64(req.Gender), 10),
		Age:       int32(req.Age),
		Introduce: req.Introduce,
	})
	if err != nil {
		resp.Msg = err.Error()
		resp.Code = user.Code_DBErr
		return nil, err
	}

	resp.Code = user.Code_Success
	return
}
