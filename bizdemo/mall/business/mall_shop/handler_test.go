// Copyright 2022 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
package main

import (
	"context"
	"testing"

	"github.com/cloudwego/kitex-examples/bizdemo/mall/business/mall_shop/kitex_gen/cmp/ecom/shop"
	"github.com/cloudwego/thriftgo/pkg/test"
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
