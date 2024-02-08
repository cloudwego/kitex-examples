package main

import (
	"context"
	"log"

	item "example_shop/kitex_gen/example/shop/item"
	"example_shop/kitex_gen/example/shop/stock"
	"example_shop/kitex_gen/example/shop/stock/stockservice"

	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
)

// ItemServiceImpl implements the last service interface defined in the IDL.
type ItemServiceImpl struct{}

func NewStockClient(addr string) (stockservice.Client, error) {
	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err)
	}
	return stockservice.NewClient("example.shop.stock", client.WithResolver(r))
}

// GetItem implements the ItemServiceImpl interface.
func (s *ItemServiceImpl) GetItem(ctx context.Context, req *item.GetItemReq) (resp *item.GetItemResp, err error) {
	resp = item.NewGetItemResp()
	resp.Item = item.NewItem()
	resp.Item.Id = req.GetId()
	resp.Item.Title = "Kitex"
	resp.Item.Description = "Kitex is an excellent framework!"

	stockCli, err := NewStockClient("0.0.0.0:8890")
	if err != nil {
		log.Fatal(err)
	}
	stockReq := stock.NewGetItemStockReq()
	stockReq.ItemId = req.GetId()
	stockResp, err := stockCli.GetItemStock(context.Background(), stockReq)
	if err != nil {
		log.Println(err)
		stockResp.Stock = 0
	}
	resp.Item.Stock = stockResp.GetStock()
	return
}
