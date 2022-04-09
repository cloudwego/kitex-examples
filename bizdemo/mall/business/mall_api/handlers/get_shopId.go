package handlers

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/cloudwego/kitex-examples/bizdemo/mall/business/mall_api/dal/client"
	"github.com/cloudwego/kitex-examples/bizdemo/mall/business/mall_api/kitex_gen/cmp/ecom/shop"
	"github.com/cloudwego/kitex-examples/bizdemo/mall/pkg/conf"
	"github.com/cloudwego/kitex-examples/bizdemo/mall/pkg/errno"
	"github.com/gin-gonic/gin"
)

// GetShopId godoc
// @Summary 商家ID查询
// @Description 通过用户ID查询商家ID
// @Tags 商家模块
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer $token"
// @Success 200 {object} handlers.Response
// @Router /shop/id [get]
func GetShopId(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	userID := int64(claims[conf.IdentityKey].(float64))

	shopId, err := client.GetShopIdByUserId(c, &shop.GetShopIdByUserIdReq{UserId: userID})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, map[string]interface{}{"shop_id": shopId})
}
