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
	"log"

	"github.com/cloudwego/kitex-examples/generic/kitex_gen/http"
	"github.com/cloudwego/kitex-examples/generic/kitex_gen/http/bizservice"
	"github.com/cloudwego/kitex/pkg/klog"
)

// BizServiceImpl implements the last service interface defined in the IDL.
type BizServiceImpl struct{}

// BizMethod1 implements the BizServiceImpl interface.
func (s *BizServiceImpl) BizMethod1(ctx context.Context, req *http.BizRequest) (resp *http.BizResponse, err error) {
	klog.Infof("BizMethod1 called, request: %#v", req)
	return &http.BizResponse{HttpCode: 200, Text: "Method1 response", Token: 1111}, nil
}

// BizMethod2 implements the BizServiceImpl interface.
func (s *BizServiceImpl) BizMethod2(ctx context.Context, req *http.BizRequest) (resp *http.BizResponse, err error) {
	klog.Infof("BizMethod2 called, request: %#v", req)
	return &http.BizResponse{HttpCode: 200, Text: "Method2 response", Token: 2222}, nil
}

// BizMethod3 implements the BizServiceImpl interface.
func (s *BizServiceImpl) BizMethod3(ctx context.Context, req *http.BizRequest) (resp *http.BizResponse, err error) {
	klog.Infof("BizMethod3 called, request: %#v", req)
	return &http.BizResponse{HttpCode: 200, Text: "Method3 response", Token: 3333}, nil
}

func main() {
	svr := bizservice.NewServer(new(BizServiceImpl))

	err := svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
