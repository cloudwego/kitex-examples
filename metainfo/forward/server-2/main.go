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

	"github.com/bytedance/gopkg/cloud/metainfo"
	"github.com/cloudwego/kitex-examples/kitex_gen/api"
	"github.com/cloudwego/kitex-examples/kitex_gen/api/echo"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/utils"
	"github.com/cloudwego/kitex/server"
)

var _ api.Echo = &EchoImpl{}

// EchoImpl implements the last service interface defined in the IDL.
type EchoImpl struct{}

// Echo implements the Echo interface.
func (s *EchoImpl) Echo(ctx context.Context, req *api.Request) (resp *api.Response, err error) {
	klog.Info("echo server-2 called")

	// `temp` not exist in server-2 context
	temp, ok1 := metainfo.GetValue(ctx, "temp")
	if ok1 {
		klog.Info(temp)
	} else {
		klog.Warn("`temp` not exist in server-2 context")
	}

	// "12345"
	logid, ok2 := metainfo.GetPersistentValue(ctx, "logid")
	if ok2 {
		klog.Info(logid)
	} else {
		klog.Warn("`logid` not exist in server-2 context")
	}

	return &api.Response{Message: req.Message}, nil
}

func main() {
	svr := echo.NewServer(new(EchoImpl), server.WithServiceAddr(utils.NewNetAddr("tcp", ":8881")))
	if err := svr.Run(); err != nil {
		log.Println("server stopped with error:", err)
	} else {
		log.Println("server stopped")
	}
}
