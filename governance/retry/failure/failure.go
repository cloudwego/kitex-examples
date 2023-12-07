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

package failure

import (
	"context"
	"errors"
	"time"

	"github.com/bytedance/gopkg/cloud/metainfo"
	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/retry"
)

func NewDelayMW(delay time.Duration) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, req, resp interface{}) (err error) {
			if _, exist := metainfo.GetPersistentValue(ctx, retry.TransitKey); !exist {
				time.Sleep(delay + 10*time.Millisecond)
				return next(ctx, req, resp)
			}
			klog.Infof("this is a retry request")
			return next(ctx, req, resp)
		}
	}
}

func NewFailureMW() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, req, resp interface{}) (err error) {
			if _, exist := metainfo.GetPersistentValue(ctx, retry.TransitKey); !exist {
				return kerrors.ErrRPCTimeout.WithCause(errors.New("you shall not pass"))
			}
			klog.Infof("this is a retry request")
			return next(ctx, req, resp)
		}
	}
}
