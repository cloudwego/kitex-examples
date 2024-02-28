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
	"fmt"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/genericclient"
	"github.com/cloudwego/kitex/pkg/generic"
)

func main() {
	// Local file idl parsing
	// YOUR_IDL_PATH thrift file path: example ./idl/example.thrift
	// includeDirs: Specify the include path. By default, the relative path of the current file is used to find include.
	p, err := generic.NewThriftFileProvider("./example_service.thrift")
	if err != nil {
		panic(err)
	}
	// Generic calls to construct JSON requests and return types
	g, err := generic.JSONThriftGeneric(p)
	if err != nil {
		panic(err)
	}
	cli, err := genericclient.NewClient("destServiceName", g, client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		panic(err)
	}
	// 'ExampleMethod' method name must be included in the idl definition
	resp, err := cli.GenericCall(context.Background(), "ExampleMethod", "{\"Msg\": \"hello\"}")
	if err != nil {
		panic(err)
	}
	// resp is a JSON string
	fmt.Println(resp)
}
