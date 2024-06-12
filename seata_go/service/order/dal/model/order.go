package model

import (
	"context"
	"gorm.io/gorm"
)

type Order struct {
	ID            uint `gorm:"primarykey"`
	UserId        string
	CommodityCode string
	Count         int32
}

func Insert(ctx context.Context, db *gorm.DB, order *Order) error {
	return db.WithContext(ctx).Create(order).Error
}
