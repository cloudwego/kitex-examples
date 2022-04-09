package main

import (
	"context"
	"github.com/cloudwego/kitex-examples/bizdemo/mall/business/mall_shop/kitex_gen/cmp/ecom/shop"
	"github.com/cloudwego/thriftgo/pkg/test"
	"testing"
)

func TestShopServiceImpl_SettleShop(t *testing.T) {
	Init()
	ctx := context.TODO()
	impl := new(ShopServiceImpl)
	resp, err := impl.SettleShop(ctx, &shop.SettleShopReq{
		UserId:   0,
		ShopName: "阿迪达斯旗舰店",
	})
	test.Assert(t, err == nil)
	test.Assert(t, resp.ShopId != 0)
}

func TestShopServiceImpl_GetShopIdByUserId(t *testing.T) {
	Init()
	ctx := context.TODO()
	impl := new(ShopServiceImpl)
	resp, err := impl.GetShopIdByUserId(ctx, &shop.GetShopIdByUserIdReq{UserId: 1})
	test.Assert(t, err == nil)
	test.Assert(t, resp.ShopId != 0)
}
