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
	"encoding/json"
	"github.com/cloudwego/kitex-examples/bizdemo/mall/business/mall_user/kitex_gen/cmp/ecom/user"
	"github.com/cloudwego/kitex-examples/bizdemo/mall/pkg/conf"
	"gorm.io/gorm"
)

type UserRole struct {
	gorm.Model
	UserName string `json:"user_name"`
	Roles    string `json:"roles"`
}

func (userRole *UserRole) TableName() string {
	return conf.UserRoleTableName
}

func AddUserRole(ctx context.Context, userName string, role user.Role) error {
	record := &UserRole{}

	insertRecord := &UserRole{
		UserName: userName,
	}
	err := DB.WithContext(ctx).Where("user_name = ?", userName).Find(&record).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			roleList := []int64{int64(role)}
			rolesBytes, _ := json.Marshal(roleList)
			insertRecord.Roles = string(rolesBytes)
			return DB.WithContext(ctx).Create(insertRecord).Error
		} else {
			return err
		}
	}
	roleListString := record.Roles
	roleList := make([]int64, 0)
	_ = json.Unmarshal([]byte((roleListString)), &roleList)
	isExist := false
	for _, roleItem := range roleList {
		if roleItem == int64(role) {
			isExist = true
			break
		}
	}
	if !isExist {
		roleList = append(roleList, int64(role))
		rolesBytes, _ := json.Marshal(roleList)
		insertRecord.Roles = string(rolesBytes)
		return DB.WithContext(ctx).Create(insertRecord).Error
	}
	return nil
}

func DelUserRole(ctx context.Context, userName string, role user.Role) error {
	record := &UserRole{}

	updateRecord := &UserRole{
		UserName: userName,
	}
	err := DB.WithContext(ctx).Where("user_name = ?", userName).Find(&record).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		} else {
			return err
		}
	}

	roleListString := record.Roles
	roleList := make([]int64, 0)
	updateRoleList := make([]int64, 0)
	_ = json.Unmarshal([]byte((roleListString)), &roleList)
	for _, roleItem := range roleList {
		if roleItem != int64(role) {
			updateRoleList = append(updateRoleList, roleItem)
		}
	}
	rolesBytes, _ := json.Marshal(updateRoleList)
	updateRecord.Roles = string(rolesBytes)
	return DB.WithContext(ctx).Create(updateRecord).Error
}

func ValidateUserRole(ctx context.Context, userName string, roles []user.Role) (bool, error) {
	record := &UserRole{}
	err := DB.WithContext(ctx).Where("user_name = ?", userName).Find(&record).Error
	if err != nil {
		return false, err
	}
	roleListString := record.Roles
	roleList := make([]int64, 0)
	_ = json.Unmarshal([]byte((roleListString)), &roleList)
	roleMap := make(map[int64]bool)
	for _, roleItem := range roleList {
		if roleItem == int64(user.Role_Admin) {
			return true, nil
		}
		roleMap[roleItem] = true
	}
	for _, role := range roles {
		if _, ok := roleMap[int64(role)]; !ok {
			return false, nil
		}
	}
	return true, nil
}
