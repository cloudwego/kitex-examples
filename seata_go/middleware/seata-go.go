// Copyright 2024 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

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
