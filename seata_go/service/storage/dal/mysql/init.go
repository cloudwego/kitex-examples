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

package mysql

import (
	"github.com/cloudwego/kitex-examples/seata_go/service/storage/dal/model"
	"github.com/cloudwego/kitex-examples/seata_go/util"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	DB, err = gorm.Open(
		mysql.New(mysql.Config{
			Conn: util.GetAtMySqlDb(),
		}),
		&gorm.Config{
			SkipDefaultTransaction: true,
			Logger:                 logger.Default.LogMode(logger.Info),
		},
	)
	if err != nil {
		panic(err)
	}

	migrate()
}

func migrate() {
	err = DB.AutoMigrate(new(model.Storage))
	if err != nil {
		panic(err)
	}
}
