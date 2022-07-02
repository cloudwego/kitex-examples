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

package db

import (
	"context"

	"github.com/cloudwego/kitex-examples/bizdemo/mall/pkg/conf"
	"gorm.io/gorm"
)

type ShopDO struct {
	gorm.Model
	ShopId   int64  `json:"shop_id"`
	ShopName string `json:"shop_name"`
	UserId   uint   `json:"user_id"`
}

func (shop *ShopDO) TableName() string {
	return conf.ShopTableName
}

func GetShopInfoByUserId(ctx context.Context, userId int64) (*ShopDO, error) {
	ret := &ShopDO{}
	if err := DB.WithContext(ctx).Where("user_id = ?", uint(userId)).Find(ret).Error; err != nil {
		return nil, err
	}
	return ret, nil
}

func CreateShop(ctx context.Context, shop *ShopDO) error {
	return DB.WithContext(ctx).Create(shop).Error
}
