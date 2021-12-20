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

package service

import (
	"context"
	"crypto/md5"
	"fmt"
	"io"

	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/user/dal/db"
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/user/errno"
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/user/kitex_gen/kitex/demo/user"
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/user/model"
)

type CreateUserService struct {
	ctx context.Context
}

func NewCreateUserService(ctx context.Context) *CreateUserService {
	return &CreateUserService{ctx: ctx}
}

func (s *CreateUserService) CreateUser(req *user.CreateUserRequest) error {
	users, err := db.QueryUser(s.ctx, req.UserName)
	if err != nil {
		return err
	}

	if len(users) != 0 {
		return errno.UserAlreadyExistErr
	}

	h := md5.New()
	io.WriteString(h, req.Password)
	passWord := fmt.Sprintf("%x", h.Sum(nil))

	return db.CreateUser(s.ctx, []*model.User{{
		UserName: req.UserName,
		Password: passWord,
	}})
}
