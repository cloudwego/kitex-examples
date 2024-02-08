package main

import (
	"context"

	stock "example_shop/kitex_gen/example/shop/stock"
)

// StockServiceImpl implements the last service interface defined in the IDL.
type StockServiceImpl struct{}

// GetItemStock implements the StockServiceImpl interface.
func (s *StockServiceImpl) GetItemStock(ctx context.Context, req *stock.GetItemStockReq) (resp *stock.GetItemStockResp, err error) {
	resp = stock.NewGetItemStockResp()
	resp.Stock = req.GetItemId()
	return
}
