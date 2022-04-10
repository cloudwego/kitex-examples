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

// BrandCreate godoc
// @Summary 商家绑定品牌
// @Description 商家绑定品牌，返回品牌ID
// @Tags 商品模块-品牌子模块
// @Accept json
// @Produce json
// @Param shopSettleParam body handlers.BrandAddParam true "品牌信息"
// @Param Authorization header string true "Bearer $token"
// @Success 200 {object} handlers.Response
// @Router /product/brand_create [post]
func BrandCreate(c *gin.Context) {
	var brandCreateParam BrandAddParam
	if err := c.ShouldBind(&brandCreateParam); err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}

	claims := jwt.ExtractClaims(c)
	userID := int64(claims[conf.IdentityKey].(float64))
	shopId, err := client.GetShopIdByUserId(c, &shop.GetShopIdByUserIdReq{UserId: userID})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}

	req := &product.AddBrandReq{
		ShopId:     shopId,
		BrandName:  brandCreateParam.BrandName,
		Logo:       brandCreateParam.Logo,
		BrandStroy: brandCreateParam.BrandStory,
	}
	brandId, err := client.CreateBrand(c, req)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, map[string]interface{}{"brand_id": brandId})
}
