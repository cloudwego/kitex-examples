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
	"github.com/cloudwego/kitex-examples/bizdemo/mall/cmd/mall_user/kitex_gen/cmp/ecom/user"
	"github.com/cloudwego/kitex-examples/bizdemo/mall/pkg/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Init init DB
func Init() {
	var err error
	DB, err = gorm.Open(mysql.Open(conf.MySQLDefaultDSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}

	m := DB.Migrator()
	hasUserTable := m.HasTable(&User{})
	hasUserRoleTable := m.HasTable(&UserRole{})
	if hasUserTable && hasUserRoleTable {
		return
	}
	if !hasUserTable {
		if err = m.CreateTable(&User{}); err != nil {
			panic(err)
		}
	}
	if !hasUserRoleTable {
		if err = m.CreateTable(&UserRole{}); err != nil {
			panic(err)
		}
		// insert admin
		AdminUser := &User{
			UserName: "admin",
			Password: "admin",
		}
		ctx := context.TODO()
		userList, err := QueryUser(ctx, "admin")
		if err != nil {
			panic(err)
		}
		if len(userList) == 0 {
			if err = CreateUser(ctx, []*User{AdminUser}); err != nil {
				panic(err)
			}
		}
		if err = AddUserRole(ctx, "admin", user.Role_Admin); err != nil {
			panic(err)
		}
	}
}
