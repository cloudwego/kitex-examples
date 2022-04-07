package main

import (
	shop "github.com/cloudwego/kitex-examples/bizdemo/mall/business/mall_shop/kitex_gen/cmp/ecom/shop/shopservice"
	"log"
)

func main() {
	svr := shop.NewServer(new(ShopServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
