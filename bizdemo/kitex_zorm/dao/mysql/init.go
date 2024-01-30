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
	"gitee.com/chunanyong/zorm"
)

var (
	dsn = "gorm:gorm@tcp(localhost:9910)/gorm?charset=utf8&parseTime=True&loc=Local"
	DB  *zorm.DBDao
)

func Init() {
	var err error
	dbConfig := &zorm.DataSourceConfig{
		DSN:        dsn,
		DriverName: "mysql",
		Dialect:    "mysql",
	}

	DB, err = zorm.NewDBDao(dbConfig)

	if err != nil {
		panic(err)
	}
}
