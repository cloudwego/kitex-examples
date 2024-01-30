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

package mysql

import (
	"context"

	"gitee.com/chunanyong/zorm"
	"github.com/cloudwego/kitex-examples/bizdemo/kitex_zorm/model"
)

func CreateUser(ctx context.Context, users []*model.User) error {
	_, err := zorm.Transaction(ctx, func(ctx context.Context) (interface{}, error) {
		// The type stored in the slice is zorm.IEntityStruct !!!
		// Uses the IEntityStruct interface and is compatible with the Struct entity class
		userSlice := make([]zorm.IEntityStruct, 0)

		for _, v := range users {
			userSlice = append(userSlice, v)
		}

		_, err := zorm.InsertSlice(ctx, userSlice)

		// If the returned err is not nil, the transaction will be rolled back
		return nil, err
	})
	return err
}

func DeleteUser(ctx context.Context, userId int64) error {
	_, err := zorm.Transaction(ctx, func(ctx context.Context) (interface{}, error) {
		demo := &model.User{}
		demo.Id = userId

		_, err := zorm.Delete(ctx, demo)

		// If the returned err is not nil, the transaction will be rolled back
		return nil, err
	})
	return err
}

func UpdateUser(ctx context.Context, user *model.User) error {
	_, err := zorm.Transaction(ctx, func(ctx context.Context) (interface{}, error) {
		_, err := zorm.Update(ctx, user)

		// If the returned err is not nil, the transaction will be rolled back
		return nil, err
	})
	return err
}

func QueryUser(ctx context.Context, keyword *string, pageNum, pageSize int64) ([]*model.User, int64, error) {
	var res []*model.User

	finder := zorm.NewFinder().Append("SELECT id FROM " + res[0].GetTableName())
	finder.Append("introduce LIKE ?", "%"+*keyword+"%")

	page := zorm.NewPage()
	page.PageNo = int(pageNum)
	page.PageSize = int(pageSize)

	err := zorm.Query(ctx, finder, &res, page)
	if err != nil {
		return nil, 0, err
	}

	return res, int64(page.TotalCount), nil
}
