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

	"github.com/cloudwego/kitex-examples/kitex_gen/api"
	"github.com/cloudwego/kitex-examples/kitex_gen/api/echo"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
)

func main() {
	cli, err := echo.NewClient("echo", client.WithHostPorts("127.0.0.1:8888"))
	if err != nil {
		klog.Warnf("failed to new client: %s", err)
		return
	}
	req := &api.Request{Message: "my request"}
	resp, err := cli.Echo(context.Background(), req)
	if err != nil {
		klog.Warnf("failed to call: %s", err)
		return
	}
	klog.Infof("resp: %s", resp.Message)
}
