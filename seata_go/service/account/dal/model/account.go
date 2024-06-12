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

type Account struct {
	ID     uint   `gorm:"primarykey"`
	UserId string `gorm:"unique"`
	Money  int32
}

func Insert(ctx context.Context, db *gorm.DB, account *Account) error {
	return db.WithContext(ctx).Create(account).Error
}

func Get(ctx context.Context, db *gorm.DB, userId string) (account *Account, err error) {
	account = new(Account)
	err = db.Where("user_id = ?", userId).First(account).Error
	return
}

func Update(ctx context.Context, db *gorm.DB, userId string, money int32) error {
	return db.WithContext(ctx).Model(new(Account)).Where("user_id = ?", userId).Update("money", money).Error
}
