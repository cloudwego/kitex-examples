/*
 * Copyright 2021 CloudWeGo
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

package service

import (
	"context"

	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/user/dal/db"
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/user/kitex_gen/kitex/demo/user"
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/user/pack"
)

type MGetUserService struct {
	ctx context.Context
}

func NewMGetUserService(ctx context.Context) *MGetUserService {
	return &MGetUserService{ctx: ctx}
}

func (s *MGetUserService) MGetUser(req *user.MGetUserRequest) ([]*user.User, error) {
	modelUsers, err := db.MGetUsers(s.ctx, req.UserIds)
	if err != nil {
		return nil, err
	}
	return pack.Users(modelUsers), nil
}
