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

package main

import (
	"context"
	"errors"

	"github.com/cloudwego/kitex/pkg/endpoint"
)

type ctxKey int

const (
	ctxPass ctxKey = iota
)

var (
	noPass  = "NoPass"
	pass    = "Pass"
	errFail = errors.New("you shall not pass")
)

func failMW(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, req, resp interface{}) (err error) {
		if status, _ := ctx.Value(ctxPass).(string); status == noPass {
			return errFail
		}
		return nil
	}
}
