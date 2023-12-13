package dal

import (
	"github.com/kitex/hello/biz/dal/mysql"
	"github.com/kitex/hello/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
