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
	"log"
	"time"

	"github.com/bytedance/gopkg/cloud/circuitbreaker"
	"github.com/cloudwego/kitex-examples/kitex_gen/api"
	"github.com/cloudwego/kitex-examples/kitex_gen/api/echo"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/circuitbreak"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

func main() {
	rpcTimeout := client.WithRPCTimeout(3 * time.Second)
	opt := circuitbreaker.Options{
		ShouldTrip: circuitbreaker.RateTripFunc(0.5, 20),
	}
	cbPanel, err := circuitbreaker.NewPanel(changeHandler, opt)
	if err != nil {
		log.Fatal(err)
	}
	cbCtrl := circuitbreak.Control{GetKey: getKey, GetErrorType: getErrorType, DecorateError: decorateError}
	cbMW := circuitbreak.NewCircuitBreakerMW(cbCtrl, cbPanel)
	client, err := echo.NewClient("echo", client.WithHostPorts("0.0.0.0:8888"), rpcTimeout, client.WithMiddleware(cbMW), client.WithMiddleware(failMW))
	if err != nil {
		log.Fatal(err)
	}
	for {
		req := &api.Request{Message: "my request"}
		ctx := context.Background()
		for i := 0; i < 20; i++ {
			ctx = context.WithValue(ctx, ctxPass, noPass)
			_, err := client.Echo(ctx, req)
			if err != nil {
				if errors.Is(err, errFail) {
					log.Println("fail")
				} else if errors.Is(err, kerrors.ErrCircuitBreak) {
					log.Println("circuitbreak")
				} else {
					log.Fatal(err)
				}
			} else {
				log.Println("success")
			}
			time.Sleep(200 * time.Millisecond)
		}
		for i := 0; i < 50; i++ {
			ctx = context.WithValue(ctx, ctxPass, pass)
			_, err := client.Echo(ctx, req)
			if err != nil {
				if errors.Is(err, errFail) {
					log.Println("fail")
				} else if errors.Is(err, kerrors.ErrCircuitBreak) {
					log.Println("circuitbreak")
				} else {
					log.Fatal(err)
				}
			} else {
				log.Println("success")
			}
			time.Sleep(200 * time.Millisecond)
		}
	}
}
