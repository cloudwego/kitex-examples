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

package service

import (
	"context"
	"github.com/cloudwego/kitex-examples/bizdemo/mall/business/mall_user/dal/db"
	"github.com/cloudwego/kitex-examples/bizdemo/mall/business/mall_user/kitex_gen/cmp/ecom/user"
)

type UserRoleService struct {
	ctx context.Context
}

func NewUserRoleService(ctx context.Context) *UserRoleService {
	return &UserRoleService{
		ctx: ctx,
	}
}

func (s *UserRoleService) AddUserRole(req *user.AddUserRoleReq) error {
	return db.AddUserRole(s.ctx, req.UserName, req.Role)
}

func (s *UserRoleService) DelUserRole(req *user.DelUserRoleReq) error {
	return db.DelUserRole(s.ctx, req.GetUserName(), req.GetRole())
}

func (s *UserRoleService) ValidateUserRole(req *user.ValidateUserRolesReq) (bool, error) {
	return db.ValidateUserRole(s.ctx, req.GetUserName(), req.GetRoles())
}
