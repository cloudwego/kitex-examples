package conf

const (
	UserTableName = "t_user"

	MySQLDefaultDSN = "gorm:gorm@tcp(localhost:3306)/gorm?charset=utf8&parseTime=True&loc=Local"
	EtcdAddress     = "127.0.0.1:2379"

	UserRpcServiceName = "cmp.ecom.user"
)
