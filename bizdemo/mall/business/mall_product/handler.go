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

	"github.com/cloudwego/kitex-examples/bizdemo/mall/business/mall_product/kitex_gen/cmp/ecom/product"
	"github.com/cloudwego/kitex-examples/bizdemo/mall/business/mall_product/pack"
	"github.com/cloudwego/kitex-examples/bizdemo/mall/business/mall_product/service"
	"github.com/cloudwego/kitex-examples/bizdemo/mall/pkg/errno"
)

// ProductServiceImpl implements the last service interface defined in the IDL.
type ProductServiceImpl struct{}

// AddBrand implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) AddBrand(ctx context.Context, req *product.AddBrandReq) (resp *product.AddBrandResp, err error) {
	resp = product.NewAddBrandResp()
	if req.BrandName == "" {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	brandId, err := service.NewBrandService(ctx).CreateBrand(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BrandId = brandId
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// UpdateBrand implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) UpdateBrand(ctx context.Context, req *product.UpdateBrandReq) (resp *product.UpdateBrandResp, err error) {
	resp = product.NewUpdateBrandResp()

	brandService := service.NewBrandService(ctx)
	isExist, err := brandService.IsBrandBelongToShop(req.BrandId, req.ShopId)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	if !isExist {
		resp.BaseResp = pack.BuildBaseResp(errno.BrandNotExistErr)
		return resp, nil
	}

	if err = brandService.UpdateBrand(req); err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// DeleteBrand implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) DeleteBrand(ctx context.Context, req *product.DeleteBrandReq) (resp *product.DeleteBrandResp, err error) {
	resp = product.NewDeleteBrandResp()

	brandService := service.NewBrandService(ctx)
	isExist, err := brandService.IsBrandBelongToShop(req.BrandId, req.ShopId)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	if !isExist {
		resp.BaseResp = pack.BuildBaseResp(errno.BrandNotExistErr)
		return resp, nil
	}

	if err = brandService.DeleteBrand(req); err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// GetBrandsByShopId implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) GetBrandsByShopId(ctx context.Context, req *product.GetBrandsByShopIdReq) (resp *product.GetBrandsByShopIdResp, err error) {
	resp = product.NewGetBrandsByShopIdResp()

	brandList, err := service.NewBrandService(ctx).GetBrandsByShopId(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	brands := make([]*product.Brand, 0)
	for _, brand := range brandList {
		brands = append(brands, &product.Brand{
			BrandId:    int64(brand.ID),
			ShopId:     brand.ShopId,
			BrandName:  brand.Name,
			Logo:       brand.Logo,
			BrandStory: brand.BrandStory,
		})
	}
	resp.Brands = brands
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}
