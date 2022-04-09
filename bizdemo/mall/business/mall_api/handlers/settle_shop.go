package handlers

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/cloudwego/kitex-examples/bizdemo/mall/business/mall_api/dal/client"
	"github.com/cloudwego/kitex-examples/bizdemo/mall/business/mall_api/kitex_gen/cmp/ecom/shop"
	"github.com/cloudwego/kitex-examples/bizdemo/mall/pkg/conf"
	"github.com/cloudwego/kitex-examples/bizdemo/mall/pkg/errno"
	"github.com/gin-gonic/gin"
)

// SettleShop godoc
// @Summary 商家入驻
// @Description 商家入驻，返回ShopID
// @Tags 商家模块
// @Accept json
// @Produce json
// @Param shopSettleParam body handlers.ShopSettleParam true "入驻材料"
// @Param Authorization header string true "Bearer $token"
// @Success 200 {object} handlers.Response
// @Router /shop/settle [post]
func SettleShop(c *gin.Context) {
	var settleParam ShopSettleParam
	if err := c.ShouldBind(&settleParam); err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}

	if len(settleParam.ShopName) == 0 {
		SendResponse(c, errno.ParamErr, nil)
		return
	}

	claims := jwt.ExtractClaims(c)
	userID := int64(claims[conf.IdentityKey].(float64))
	shopId, err := client.SettleShop(c, &shop.SettleShopReq{
		UserId:   userID,
		ShopName: settleParam.ShopName,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, map[string]interface{}{"shop_id": shopId})
}
