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
	"log"
	"strings"
	"time"
)

func NewFuture(f func() (interface{}, error)) func() (interface{}, error) {
	var res interface{}
	var err error

	c := make(chan struct{}, 1)
	go func() {
		defer close(c)
		res, err = f()
	}()
	return func() (interface{}, error) {
		<-c
		return res, err
	}
}

func sequentialCall(client echo.Client) {

	for i := 1; i < 5; i++ {
		var req = &api.Request{Message: "my request"}
		resp, err := client.Echo(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(resp)
	}
}

func asyncParallelCall(client echo.Client) {
	var futures []func() (interface{}, error)
	for i := 0; i < 5; i++ {
		var req = &api.Request{Message: "my request"}

		futures = append(futures, NewFuture(func() (interface{}, error) {
			return client.Echo(context.Background(), req)
		}))

	}
	for i := 0; i < 5; i++ {
		resp, err := futures[i]()
		if err != nil {
			log.Fatal(err)
		}
		log.Println(resp)
	}
}

func main() {
	client, err := echo.NewClient("echo", client.WithHostPorts("[::1]:8888"))
	if err != nil {
		log.Fatal(err)
	}

	log.Println(strings.Repeat("=", 10))
	log.Println("sequential call")
	t0 := time.Now()
	sequentialCall(client)
	log.Println("cast time: " + time.Since(t0).String())

	log.Println(strings.Repeat("=", 10))
	log.Println("async parallel call")
	t1 := time.Now()
	asyncParallelCall(client)
	log.Println("cast time: " + time.Since(t1).String())

}
