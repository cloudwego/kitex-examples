package main

import (
	"context"
	"github.com/cloudwego/kitex-examples/bizdemo/mall/business/mall_shop/kitex_gen/cmp/ecom/shop"
)

// ShopServiceImpl implements the last service interface defined in the IDL.
type ShopServiceImpl struct{}

// SettleShop implements the ShopServiceImpl interface.
func (s *ShopServiceImpl) SettleShop(ctx context.Context, req *shop.SettleShopReq) (resp *shop.SettleShopResp, err error) {
	// TODO: Your code here...
	return
}

// GetShopIdByName implements the ShopServiceImpl interface.
func (s *ShopServiceImpl) GetShopIdByName(ctx context.Context, req *shop.GetShopIdByNameReq) (resp *shop.GetShopIdByNameResp, err error) {
	// TODO: Your code here...
	return
}
