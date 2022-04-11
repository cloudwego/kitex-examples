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
	"github.com/cloudwego/thriftgo/pkg/test"
	"testing"
)

func TestProductServiceImpl_AddBrand(t *testing.T) {
	Init()
	ctx := context.TODO()
	impl := new(ProductServiceImpl)
	resp, err := impl.AddBrand(ctx, &product.AddBrandReq{
		ShopId:     3007482826,
		BrandName:  "中国李宁",
		Logo:       "",
		BrandStroy: "",
	})
	test.Assert(t, err == nil)
	test.Assert(t, resp.BaseResp.StatusCode == 0)
	test.Assert(t, resp.BrandId != 0)
}

func TestProductServiceImpl_GetBrandsByShopId(t *testing.T) {
	Init()
	ctx := context.TODO()
	impl := new(ProductServiceImpl)
	resp, err := impl.GetBrandsByShopId(ctx, &product.GetBrandsByShopIdReq{ShopId: 1555281180})
	test.Assert(t, err == nil)
	test.Assert(t, resp.BaseResp.StatusCode == 0)
	test.Assert(t, resp.Brands[0].BrandName == "Adidas")
}

func TestProductServiceImpl_UpdateBrand(t *testing.T) {
	Init()
	ctx := context.TODO()
	impl := new(ProductServiceImpl)
	brandName := "ADIDAS"
	resp, err := impl.UpdateBrand(ctx, &product.UpdateBrandReq{
		BrandId:   1,
		ShopId:    1555281180,
		BrandName: &brandName,
	})
	test.Assert(t, err == nil)
	test.Assert(t, resp.BaseResp.StatusCode == 0)
}

func TestProductServiceImpl_DeleteBrand(t *testing.T) {
	Init()
	ctx := context.TODO()
	impl := new(ProductServiceImpl)
	resp, err := impl.DeleteBrand(ctx, &product.DeleteBrandReq{
		BrandId: 1,
		ShopId:  1555281180,
	})
	test.Assert(t, err == nil)
	test.Assert(t, resp.BaseResp.StatusCode == 0)
}
