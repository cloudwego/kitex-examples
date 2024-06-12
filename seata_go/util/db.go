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
