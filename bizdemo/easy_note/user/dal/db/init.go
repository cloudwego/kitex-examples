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

	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/user/constant"
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/user/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
	_ "gorm.io/plugin/dbresolver"
	gormopentracing "gorm.io/plugin/opentracing"
)

var connPool *gorm.DB

func Init() {
	var err error
	dbDSN := constant.MySQLDefaultDSN

	connPool, err = gorm.Open(mysql.Open(dbDSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}
	connPool.Use(gormopentracing.New())
	m := connPool.Migrator()
	if m.HasTable(model.User{}) {
		return
	}
	if err := m.CreateTable(model.User{}); err != nil {
		panic(err)
	}
}

func GetDBWriter(ctx context.Context) *gorm.DB {
	return connPool.WithContext(ctx).Clauses(dbresolver.Write)
}

func GetDBReader(ctx context.Context) *gorm.DB {
	return connPool.WithContext(ctx).Clauses(dbresolver.Read)
}
