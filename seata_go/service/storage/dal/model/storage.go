// Copyright 2024 CloudWeGo Authors
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

package model

import (
	"context"

	"gorm.io/gorm"
)

type Storage struct {
	ID            uint   `gorm:"id;primarykey" json:"id"`
	CommodityCode string `gorm:"commodity_code;unique" json:"commodity_code"`
	Count         int32  `gorm:"count" json:"count"`
	Price         int32  `gorm:"price" json:"price"`
}

func Insert(ctx context.Context, db *gorm.DB, storage *Storage) error {
	return db.WithContext(ctx).Create(storage).Error
}

func Get(ctx context.Context, db *gorm.DB, commodityCode string) (storage *Storage, err error) {
	storage = new(Storage)
	err = db.Where("commodity_code = ?", commodityCode).First(storage).Error
	return
}

func Update(ctx context.Context, db *gorm.DB, commodityCode string, count int32) error {
	return db.WithContext(ctx).Model(new(Storage)).Where("commodity_code = ?", commodityCode).Update("count", count).Error
}
