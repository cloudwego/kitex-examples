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
package service

import (
	"context"
	"hash/crc32"

	"github.com/cloudwego/kitex-examples/bizdemo/mall/business/mall_shop/dal/db"
	"github.com/cloudwego/kitex-examples/bizdemo/mall/business/mall_shop/kitex_gen/cmp/ecom/shop"
	"github.com/cloudwego/kitex-examples/bizdemo/mall/pkg/errno"
	uuid "github.com/satori/go.uuid"
)

type ShopService struct {
	ctx context.Context
}

func NewShopService(ctx context.Context) *ShopService {
	return &ShopService{ctx: ctx}
}

func (s *ShopService) SettleShop(req *shop.SettleShopReq) (int64, error) {
	shopInfo, err := db.GetShopInfoByUserId(s.ctx, req.GetUserId())
	if err != nil {
		return 0, err
	}
	if shopInfo.ShopId != 0 {
		return 0, errno.ShopAlreadyExistErr
	}

	shopId := GenShopId()
	if err := db.CreateShop(s.ctx, &db.ShopDO{
		ShopId:   shopId,
		ShopName: req.ShopName,
		UserId:   uint(req.UserId),
	}); err != nil {
		return 0, err
	}
	return shopId, nil
}

func (s *ShopService) GetShopIdByUserId(req *shop.GetShopIdByUserIdReq) (int64, error) {
	shopInfo, err := db.GetShopInfoByUserId(s.ctx, req.GetUserId())
	if err != nil {
		return 0, err
	}
	if shopInfo.ShopId == 0 {
		return 0, errno.ShopNotExistErr
	}
	return shopInfo.ShopId, nil
}

func GenShopId() int64 {
	uuid := uuid.NewV4()
	uuidHash := int64(crc32.ChecksumIEEE([]byte(uuid.String())))
	return uuidHash
}
