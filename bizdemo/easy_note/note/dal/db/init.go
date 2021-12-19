package db

import (
	"context"

	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/note/constant"
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/note/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
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
	if m.HasTable(model.Note{}) {
		return
	}
	if err := m.CreateTable(model.Note{}); err != nil {
		panic(err)
	}
}

func GetDBWriter(ctx context.Context) *gorm.DB {
	return connPool.WithContext(ctx).Clauses(dbresolver.Write)
}

func GetDBReader(ctx context.Context) *gorm.DB {
	return connPool.WithContext(ctx).Clauses(dbresolver.Read)
}
