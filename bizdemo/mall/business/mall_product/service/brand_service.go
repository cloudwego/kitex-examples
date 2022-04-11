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
	"github.com/cloudwego/kitex-examples/bizdemo/mall/business/mall_product/dal/db"
	"github.com/cloudwego/kitex-examples/bizdemo/mall/business/mall_product/kitex_gen/cmp/ecom/product"
)

type BrandService struct {
	ctx context.Context
}

func NewBrandService(ctx context.Context) *BrandService {
	return &BrandService{ctx: ctx}
}

func (b *BrandService) CreateBrand(req *product.AddBrandReq) (int64, error) {
	brandId, err := db.CreateBrand(b.ctx, &db.BrandDO{
		ShopId:     req.ShopId,
		Name:       req.BrandName,
		Logo:       req.Logo,
		BrandStory: req.BrandStroy,
	})
	if err != nil {
		return 0, err
	}
	return brandId, nil
}

func (b *BrandService) UpdateBrand(req *product.UpdateBrandReq) error {
	return db.UpdateBrand(b.ctx, req.BrandId, req.ShopId, req.BrandName, req.Logo, req.BrandStory)
}

func (b *BrandService) DeleteBrand(req *product.DeleteBrandReq) error {
	return db.DeleteBrand(b.ctx, req.BrandId, req.ShopId)
}

func (b *BrandService) GetBrandsByShopId(req *product.GetBrandsByShopIdReq) ([]*db.BrandDO, error) {
	brandList, err := db.GetBrandInfoByShopId(b.ctx, req.ShopId)
	if err != nil {
		return nil, err
	}
	return brandList, nil
}

func (b *BrandService) IsBrandBelongToShop(brandId, shopId int64) (bool, error) {
	brandList, err := b.GetBrandsByShopId(&product.GetBrandsByShopIdReq{ShopId: shopId})
	if err != nil {
		return false, err
	}
	for _, brand := range brandList {
		if int64(brand.ID) == brandId {
			return true, nil
		}
	}
	return false, nil
}
