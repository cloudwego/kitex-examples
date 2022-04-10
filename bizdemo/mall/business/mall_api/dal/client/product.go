package client

import (
	"context"
	"github.com/cloudwego/kitex-examples/bizdemo/mall/business/mall_api/kitex_gen/cmp/ecom/product"
	"github.com/cloudwego/kitex-examples/bizdemo/mall/business/mall_api/kitex_gen/cmp/ecom/product/productservice"
	"github.com/cloudwego/kitex-examples/bizdemo/mall/pkg/conf"
	"github.com/cloudwego/kitex-examples/bizdemo/mall/pkg/errno"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	"time"
)

var productClient productservice.Client

func initProductRpc() {
	r, err := etcd.NewEtcdResolver([]string{conf.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := productservice.NewClient(
		conf.ProductRpcServiceName,
		client.WithRPCTimeout(3*time.Second),              // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond),    // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithResolver(r),                            // resolver
	)
	if err != nil {
		panic(err)
	}
	productClient = c
}

func GetBrands(ctx context.Context, req *product.GetBrandsByShopIdReq) ([]*product.Brand, error) {
	resp, err := productClient.GetBrandsByShopId(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(int64(resp.BaseResp.StatusCode), resp.BaseResp.StatusMessage)
	}
	return resp.Brands, nil
}

func CreateBrand(ctx context.Context, req *product.AddBrandReq) (int64, error) {
	resp, err := productClient.AddBrand(ctx, req)
	if err != nil {
		return 0, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return 0, errno.NewErrNo(int64(resp.BaseResp.StatusCode), resp.BaseResp.StatusMessage)
	}
	return resp.GetBrandId(), nil
}

func UpdateBrand(ctx context.Context, req *product.UpdateBrandReq) error {
	resp, err := productClient.UpdateBrand(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(int64(resp.BaseResp.StatusCode), resp.BaseResp.StatusMessage)
	}
	return nil
}

func DelBrand(ctx context.Context, req *product.DeleteBrandReq) error {
	resp, err := productClient.DeleteBrand(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(int64(resp.BaseResp.StatusCode), resp.BaseResp.StatusMessage)
	}
	return nil
}
