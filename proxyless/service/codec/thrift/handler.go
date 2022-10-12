/*
 * Copyright 2022 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"context"

	proxyless "github.com/cloudwego/kitex-examples/proxyless/service/codec/thrift/kitex_gen/proxyless"
)

// GreetServiceImpl implements the last service interface defined in the IDL.
type GreetServiceImpl struct{}

// SayHello1 implements the GreetServiceImpl interface.
func (s *GreetServiceImpl) SayHello1(ctx context.Context, request *proxyless.HelloRequest) (resp *proxyless.HelloResponse, err error) {
	// TODO: Your code here...
	return
}

// SayHello2 implements the GreetServiceImpl interface.
func (s *GreetServiceImpl) SayHello2(ctx context.Context, request *proxyless.HelloRequest) (resp *proxyless.HelloResponse, err error) {
	// TODO: Your code here...
	return
}
