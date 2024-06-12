package middleware

import (
	"context"
	"github.com/bytedance/gopkg/cloud/metainfo"
	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/klog"
	"seata.apache.org/seata-go/pkg/constant"

	"seata.apache.org/seata-go/pkg/tm"
)

func SeataGoClientMiddleware(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, req, resp interface{}) (err error) {
		xid := tm.GetXID(ctx)
		// attach the xid to the context
		ctx = metainfo.WithPersistentValue(ctx, constant.XidKey, xid)
		return next(ctx, req, resp)
	}
}

func SeataGoServerMiddleware(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, req, resp interface{}) (err error) {
		// get xid from context
		xid, ok := metainfo.GetPersistentValue(ctx, constant.XidKey)
		if !ok {
			klog.Errorf("the request context does not contain %s, global transaction xid", constant.XidKey)
			return next(ctx, req, resp)
		}
		// initialize this context using tm
		ctx = tm.InitSeataContext(ctx)
		tm.SetXID(ctx, xid)
		return next(ctx, req, resp)
	}
}
