package handlers

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/cloudwego/kitex-examples/bizdemo/mall/business/mall_api/dal/client"
	"github.com/cloudwego/kitex-examples/bizdemo/mall/business/mall_api/kitex_gen/cmp/ecom/product"
	"github.com/cloudwego/kitex-examples/bizdemo/mall/business/mall_api/kitex_gen/cmp/ecom/shop"
	"github.com/cloudwego/kitex-examples/bizdemo/mall/pkg/conf"
	"github.com/cloudwego/kitex-examples/bizdemo/mall/pkg/errno"
	"github.com/gin-gonic/gin"
)

// GetBrands godoc
// @Summary 商家绑定品牌查询
// @Description 商家绑定品牌查询
// @Tags 商品模块-品牌子模块
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer $token"
// @Success 200 {object} handlers.Response
// @Router /product/get_brands [get]
func GetBrands(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	userID := int64(claims[conf.IdentityKey].(float64))

	shopId, err := client.GetShopIdByUserId(c, &shop.GetShopIdByUserIdReq{UserId: userID})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}

	brands, err := client.GetBrands(c, &product.GetBrandsByShopIdReq{ShopId: shopId})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, map[string]interface{}{"brands": brands})
}
