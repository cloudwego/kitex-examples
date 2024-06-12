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

package main

import (
	"context"
	"time"

	"github.com/cloudwego/kitex-examples/seata_go/kitex_gen/order/orderservice"
	"github.com/cloudwego/kitex-examples/seata_go/kitex_gen/storage/storageservice"
	"github.com/cloudwego/kitex-examples/seata_go/middleware"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/transport"
	"seata.apache.org/seata-go/pkg/tm"

	seataclient "seata.apache.org/seata-go/pkg/client"
)

var (
	storageServerAddr = "localhost:8882"
	orderServerAddr   = "localhost:8883"
	storageClient     storageservice.Client
	orderClient       orderservice.Client
)

func init() {
	seataclient.InitPath("conf/seatago.yml")
	initClient()
}

func main() {
	// refer to [use case](https://seata.apache.org/docs/user/quickstart/#use-case)
	// to simply simulate create order
	err := tm.WithGlobalTx(context.Background(), &tm.GtxConfig{
		Name:    "CreateOrderTx",
		Timeout: time.Second * 30,
	}, createOrder)
	if err != nil {
		panic(err)
	}

	klog.Info("create order successfully")
}

func createOrder(ctx context.Context) error {
	userId := "user_1"
	commodityCode := "commodity_1"
	count := int32(1)

	// deduct the count of commodity in storage
	err := storageClient.Deduct(ctx, commodityCode, count)
	if err != nil {
		klog.Errorf("deduct commodity err: %v", err)
		return err
	}

	// create order
	err = orderClient.Create(ctx, userId, commodityCode, count)
	if err != nil {
		klog.Errorf("create order err: %v", err)
		return err
	}

	return nil
}

func initClient() {
	// use seata-go middleware when initializing the client
	storageClient = storageservice.MustNewClient("storage",
		client.WithHostPorts(storageServerAddr),
		client.WithMiddleware(middleware.SeataGoClientMiddleware),
		client.WithTransportProtocol(transport.TTHeaderFramed))
	orderClient = orderservice.MustNewClient("order",
		client.WithHostPorts(orderServerAddr),
		client.WithMiddleware(middleware.SeataGoClientMiddleware),
		client.WithTransportProtocol(transport.TTHeaderFramed))
}
