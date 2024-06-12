package main

import (
	"context"
	"github.com/cloudwego/kitex-examples/seata_go/service/order/dal/model"
	"github.com/cloudwego/kitex-examples/seata_go/service/order/dal/mysql"
	"github.com/cloudwego/kitex-examples/seata_go/service/order/rpc"
	"github.com/cloudwego/kitex/pkg/klog"
)

// OrderServiceImpl implements the last service interface defined in the IDL.
type OrderServiceImpl struct{}

// Create implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) Create(ctx context.Context, userId string, commodityCode string, count int32) (err error) {
	// calculate the money of the order
	orderMoney, err := rpc.StorageClient.Calculate(ctx, commodityCode, count)
	if err != nil {
		klog.Errorf("calculate order money failed: %v", err)
		return err
	}
	// deduct the user's balance
	err = rpc.AccountClient.Deduct(ctx, userId, orderMoney)
	if err != nil {
		klog.Errorf("deduct account money failed: %v", err)
		return err
	}
	// insert new order
	err = model.Insert(ctx, mysql.DB, &model.Order{
		UserId:        userId,
		CommodityCode: commodityCode,
		Count:         count,
	})
	if err != nil {
		klog.Errorf("insert order failed: %v", err)
		return err
	}
	return nil
}
