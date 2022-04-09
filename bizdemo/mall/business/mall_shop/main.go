package main

import (
	"github.com/cloudwego/kitex-examples/bizdemo/mall/business/mall_shop/dal"
	shop "github.com/cloudwego/kitex-examples/bizdemo/mall/business/mall_shop/kitex_gen/cmp/ecom/shop/shopservice"
	"log"
)

func Init() {
	dal.Init()
}

func main() {
	svr := shop.NewServer(new(ShopServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
