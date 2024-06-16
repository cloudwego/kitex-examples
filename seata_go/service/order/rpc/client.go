// Copyright 2024 CloudWeGo Authors
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

package rpc

import (
	"github.com/cloudwego/kitex-examples/seata_go/kitex_gen/account/accountservice"
	"github.com/cloudwego/kitex-examples/seata_go/kitex_gen/storage/storageservice"
	"github.com/cloudwego/kitex-examples/seata_go/middleware"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/transport"
)

var (
	accountServerAddr = "localhost:8881"
	storageServerAddr = "localhost:8882"
	AccountClient     accountservice.Client
	StorageClient     storageservice.Client
)

func InitClient() {
	AccountClient = accountservice.MustNewClient("account",
		client.WithHostPorts(accountServerAddr),
		client.WithMiddleware(middleware.SeataGoClientMiddleware),
		client.WithTransportProtocol(transport.TTHeaderFramed))
	StorageClient = storageservice.MustNewClient("storage",
		client.WithHostPorts(storageServerAddr),
		client.WithMiddleware(middleware.SeataGoClientMiddleware),
		client.WithTransportProtocol(transport.TTHeaderFramed))
}
