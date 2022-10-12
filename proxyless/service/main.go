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
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"

	"github.com/cloudwego/kitex-examples/proxyless/config"
	"github.com/cloudwego/kitex-examples/proxyless/service/src"
	"github.com/cloudwego/kitex-examples/proxyless/service/src/benchmark"
)

const (
	serviceNameKey = "MY_SERVICE_NAME"
	pprofAddr      = ":6789"
)

func pprof() {
	_ = http.ListenAndServe(pprofAddr, nil)
}

func main() {
	go pprof()
	serviceName, ok := os.LookupEnv(serviceNameKey)
	if !ok {
		panic("Please specify the service name")
	}

	// use the current namespace
	namespace, ok := os.LookupEnv(config.POD_NAMESPACE_KEY)
	if !ok {
		panic("Please specify the namespace")
	}

	var svc src.TestService
	switch serviceName {
	case config.TestClientSvc:
		svc = src.NewProxylessClient(fmt.Sprintf("%s.%s.%s:%s", config.TestServerSvc, namespace, config.Suffix, config.ServerServicePort))
	case config.TestServerSvc:
		svc = src.NewProxylessServer()
	case config.BenchmarkClient:
		svc = benchmark.NewBenchmarkRunner(fmt.Sprintf("%s.%s.%s:%s", config.TestServerSvc, namespace, config.Suffix, config.ServerServicePort))
	}
	fmt.Println("TEST SERVICE START")
	err := svc.Run()
	if err != nil {
		errMsg := fmt.Errorf("test failed with error: %s", err.Error())
		fmt.Println(errMsg)
	}
}
