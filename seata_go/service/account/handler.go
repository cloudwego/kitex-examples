package main

import (
	"context"
	"github.com/cloudwego/kitex-examples/seata_go/service/account/dal/model"
	"github.com/cloudwego/kitex-examples/seata_go/service/account/dal/mysql"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

// AccountServiceImpl implements the last service interface defined in the IDL.
type AccountServiceImpl struct{}

// Deduct implements the AccountServiceImpl interface.
func (s *AccountServiceImpl) Deduct(ctx context.Context, userId string, money int32) (err error) {
	err = mysql.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		account, err := model.Get(ctx, tx, userId)
		if err != nil {
			klog.Errorf("get account failed: %v", err)
			return err
		}
		err = model.Update(ctx, tx, userId, account.Money-money)
		if err != nil {
			klog.Errorf("update account failed: %v", err)
			return err
		}
		return nil
	})
	return
}
