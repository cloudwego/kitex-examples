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

package db

import (
	"context"

	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/user/model"
)

func MGetUsers(ctx context.Context, userIDs []int64) ([]*model.User, error) {
	res := make([]*model.User, 0)
	if len(userIDs) == 0 {
		return res, nil
	}

	if err := DB.WithContext(ctx).Where("id in ?", userIDs).Find(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}

func CreateUser(ctx context.Context, users []*model.User) error {

	return DB.WithContext(ctx).Create(users).Error
}

func QueryUser(ctx context.Context, userName string) ([]*model.User, error) {
	res := make([]*model.User, 0)
	if err := DB.WithContext(ctx).Where("user_name = ?", userName).Find(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}