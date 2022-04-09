package db

import (
	"context"
	"github.com/cloudwego/kitex-examples/bizdemo/mall/pkg/conf"
	"gorm.io/gorm"
)

type ShopDO struct {
	gorm.Model
	ShopId   int64  `json:"shop_id"`
	ShopName string `json:"shop_name"`
	UserId   uint   `json:"user_id"`
}

func (shop *ShopDO) TableName() string {
	return conf.ShopTableName
}

func GetShopInfoByUserId(ctx context.Context, userId int64) (*ShopDO, error) {
	ret := &ShopDO{}
	if err := DB.WithContext(ctx).Where("user_id = ?", uint(userId)).Find(ret).Error; err != nil {
		return nil, err
	}
	return ret, nil
}

func CreateShop(ctx context.Context, shop *ShopDO) error {
	return DB.WithContext(ctx).Create(shop).Error
}
