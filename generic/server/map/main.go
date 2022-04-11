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

	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server/genericserver"
)

var _ generic.Service = &GenericServiceImpl{}

// EchoImpl implements the last service interface defined in the IDL.
type GenericServiceImpl struct{}

// GenericCall implements the Echo interface.
func (s *GenericServiceImpl) GenericCall(ctx context.Context, method string, request interface{}) (response interface{}, err error) {
	m := request.(map[string]interface{})
	klog.Infof("Recv: %v\n", m)
	return map[string]interface{}{
		"message": m["message"],
	}, nil
}

func main() {
	path := "../echo.thrift" // depends on current directory
	p, err := generic.NewThriftFileProvider(path)
	if err != nil {
		klog.Fatalf("new thrift file provider failed: %v", err)
	}
	g, err := generic.MapThriftGeneric(p)
	if err != nil {
		klog.Fatalf("new map thrift generic failed: %v", err)
	}
	svr := genericserver.NewServer(new(GenericServiceImpl), g)
	if err := svr.Run(); err != nil {
		klog.Error("server stopped with error:", err)
	} else {
		klog.Error("server stopped")
	}
}
