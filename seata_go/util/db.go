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

package util

import (
	"database/sql"
	"fmt"

	sql2 "seata.apache.org/seata-go/pkg/datasource/sql"
)

const dsn = "root:root@tcp(localhost:3306)/seata_go?multiStatements=true&interpolateParams=true"

func GetAtMySqlDb() *sql.DB {
	db, err := sql.Open(sql2.SeataATMySQLDriver, dsn)
	if err != nil {
		panic(fmt.Errorf("init seata at mysql driver error: %w", err))
	}
	return db
}
