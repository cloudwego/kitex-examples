package db

import "gorm.io/gorm"

type SpecPrice struct {
}
type Product struct {
	gorm.Model
	ProductId   int64    `json:"product_id"`  // 商品ID
	ShopId      int64    `json:"shop_id"`     // 店铺ID
	Name        string   `json:"name"`        // 商品名
	Description string   `json:"description"` // 详情
	Img         []string `json:"img"`         // 主图
	Status      int64    `json:"status"`      // 商品状态
}
