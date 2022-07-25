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

package handlers

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/kitex-examples/bizdemo/mall/cmd/mall_api/dal/client"
	"github.com/cloudwego/kitex-examples/bizdemo/mall/cmd/mall_api/kitex_gen/cmp/ecom/shop"
	"github.com/cloudwego/kitex-examples/bizdemo/mall/pkg/conf"
	"github.com/cloudwego/kitex-examples/bizdemo/mall/pkg/errno"
	"github.com/hertz-contrib/jwt"
)

// SettleShop godoc
// @Summary 商家入驻
// @Description 商家入驻，返回ShopID
// @Tags 商家模块
// @Accept json
// @Produce json
// @Param shopSettleParam body handlers.ShopSettleParam true "入驻材料"
// @Security TokenAuth
// @Success 200 {object} handlers.Response
// @Router /shop/settle [post]
func SettleShop(ctx context.Context, c *app.RequestContext) {
	var settleParam ShopSettleParam
	if err := c.BindAndValidate(&settleParam); err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}

	if len(settleParam.ShopName) == 0 {
		SendResponse(c, errno.ParamErr, nil)
		return
	}

	claims := jwt.ExtractClaims(ctx, c)
	userID := int64(claims[conf.IdentityKey].(float64))
	shopId, err := client.SettleShop(ctx, &shop.SettleShopReq{
		UserId:   userID,
		ShopName: settleParam.ShopName,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, map[string]interface{}{"shop_id": shopId})
}
