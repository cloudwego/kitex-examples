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
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/genericclient"
	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/cloudwego/kitex/pkg/klog"
)

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
	cli, err := genericclient.NewClient("echo", g, client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		klog.Fatalf("new http generic client failed: %v", err)
	}
	for {
		resp, err := cli.GenericCall(context.Background(), "echo", map[string]interface{}{
			"message": "my request",
		})
		if err != nil {
			log.Fatal(err)
		}
		klog.Infof("response: %#v", resp)
		time.Sleep(time.Second)
	}
}
