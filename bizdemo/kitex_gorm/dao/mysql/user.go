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

import "github.com/cloudwego/kitex-examples/bizdemo/kitex_gorm/model"

func CreateUser(users []*model.User) error {
	return DB.Create(users).Error
}

func DeleteUser(userId int64) error {
	return DB.Where("id = ?", userId).Delete(&model.User{}).Error
}

func UpdateUser(user *model.User) error {
	return DB.Updates(user).Error
}

func QueryUser(keyword *string, page, pageSize int64) ([]*model.User, int64, error) {
	db := DB.Model(model.User{})
	if keyword != nil && len(*keyword) != 0 {
		db = db.Where(DB.Or("name like ?", "%"+*keyword+"%").
			Or("introduce like ?", "%"+*keyword+"%"))
	}
	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var res []*model.User
	if err := db.Limit(int(pageSize)).Offset(int(pageSize * (page - 1))).Find(&res).Error; err != nil {
		return nil, 0, err
	}
	return res, total, nil
}
