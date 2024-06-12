package model

import (
	"context"
	"gorm.io/gorm"
)

type Account struct {
	ID     uint   `gorm:"primarykey"`
	UserId string `gorm:"unique"`
	Money  int32
}

func Insert(ctx context.Context, db *gorm.DB, account *Account) error {
	return db.WithContext(ctx).Create(account).Error
}

func Get(ctx context.Context, db *gorm.DB, userId string) (account *Account, err error) {
	account = new(Account)
	err = db.Where("user_id = ?", userId).First(account).Error
	return
}

func Update(ctx context.Context, db *gorm.DB, userId string, money int32) error {
	return db.WithContext(ctx).Model(new(Account)).Where("user_id = ?", userId).Update("money", money).Error
}
