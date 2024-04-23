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

	"github.com/apache/thrift/lib/go/thrift"
	"github.com/cloudwego/kitex-examples/kitex_gen/api"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/genericclient"
	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/utils"
)

func main() {
	genericCli, err := genericclient.NewClient("echo", generic.BinaryThriftGeneric(), client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		log.Fatal(err)
	}
	codec := utils.NewThriftMessageCodec()
	for {
		ctx := context.Background()
		buf, err := codec.Encode("echo", thrift.CALL, 0, &api.EchoEchoArgs{Req: &api.Request{Message: "my request"}})
		if err != nil {
			klog.Fatal(err)
		}
		resp, err := genericCli.GenericCall(ctx, "echo", buf)
		if err != nil {
			klog.Errorf("call echo failed: %w\n", err)
		}
		result := &api.EchoEchoResult{}
		_, _, err = codec.Decode(resp.([]byte), result)
		if err != nil {
			klog.Fatal(err)
		}
		klog.Info(result.Success)
		time.Sleep(time.Second)
	}
}
