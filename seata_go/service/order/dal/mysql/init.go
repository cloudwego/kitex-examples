package mysql

import (
	"github.com/cloudwego/kitex-examples/seata_go/service/order/dal/model"
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
	err = DB.AutoMigrate(new(model.Order))
	if err != nil {
		panic(err)
	}
}
