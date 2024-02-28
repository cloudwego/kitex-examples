// Copyright 2023 CloudWeGo Authors
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

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/transport"
	"grpc_multi_service_client/kitex_gen/multi/service"
	"grpc_multi_service_client/kitex_gen/multi/service/serviceb"

	"grpc_multi_service_client/kitex_gen/multi/service/servicea"
)

func main() {
	clienta := servicea.MustNewClient("servicea", client.WithTransportProtocol(transport.GRPC), client.WithHostPorts("127.0.0.1:8888"))
	clientb := serviceb.MustNewClient("serviceb", client.WithTransportProtocol(transport.GRPC), client.WithHostPorts("127.0.0.1:8888"))

	resa, err := clienta.ChatA(context.Background(), &service.RequestA{Name: "hello,a"})
	if err != nil {
		klog.Error(err)
		return
	}

	resb, err := clientb.ChatB(context.Background(), &service.RequestB{Name: "hello,b"})
	if err != nil {
		klog.Error(err)
		return
	}

	klog.Info("resa: ", resa.Message)
	klog.Info("resb: ", resb.Message)
}
