// Copyright 2021 CloudWeGo Authors
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

package mymiddleware

import (
	"context"

	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
)

var _ endpoint.Middleware = CommonMiddleware

type args interface {
	GetFirstArgument() interface{}
}

type result interface {
	GetResult() interface{}
}

func CommonMiddleware(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, req, resp interface{}) (err error) {
		ri := rpcinfo.GetRPCInfo(ctx)
		// get real request
		klog.Infof("real request: %+v\n", req.(args).GetFirstArgument())
		// get local service information
		klog.Infof("local service name: %v\n", ri.From().ServiceName())
		// get remote service information
		klog.Infof("remote service name: %v, remote method: %v\n", ri.To().ServiceName(), ri.To().Method())
		if err := next(ctx, req, resp); err != nil {
			return err
		}
		// get real response
		klog.Infof("real response: %+v\n", resp.(result).GetResult())
		return nil
	}
}

func ClientMiddleware(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, req, resp interface{}) (err error) {
		ri := rpcinfo.GetRPCInfo(ctx)
		// get timeout information
		klog.Infof("rpc timeout: %v, readwrite timeout: %v\n", ri.Config().RPCTimeout(), ri.Config().ConnectTimeout())
		if err := next(ctx, req, resp); err != nil {
			return err
		}
		// get server information
		klog.Infof("server address: %v\n", ri.To().Address())
		return nil
	}
}

func ServerMiddleware(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, req, resp interface{}) (err error) {
		ri := rpcinfo.GetRPCInfo(ctx)
		// get client information
		klog.Infof("client address: %v\n", ri.From().Address())
		if err := next(ctx, req, resp); err != nil {
			return err
		}
		return nil
	}
}
