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

package benchmark

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"

	kxds "github.com/cloudwego/kitex/pkg/xds"

	"github.com/cloudwego/kitex-benchmark/perf"
	"github.com/cloudwego/kitex-benchmark/runner"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/connpool"
	"github.com/cloudwego/kitex/pkg/klog"
	dns "github.com/kitex-contrib/resolver-dns"
	"github.com/kitex-contrib/xds"
	"github.com/kitex-contrib/xds/xdssuite"

	"github.com/cloudwego/kitex-examples/proxyless/service/codec/thrift/kitex_gen/proxyless"
	"github.com/cloudwego/kitex-examples/proxyless/service/codec/thrift/kitex_gen/proxyless/greetservice"
)

const (
	meshModeKey       = "mesh_mode"
	meshModeProxyless = "kitex-proxyless"
	meshModeProxy     = "kitex-sidecar"
	meshModeDirect    = "kitex-direct"

	totalReqKey    = "total_req"
	concurrencyKey = "concurrency"
)

var (
	totalReq    int64 = 1000000
	concurrency       = 10
	echoSize          = 1024
)

func initEnv() {
	if r, ok := os.LookupEnv(totalReqKey); ok {
		totalReq, _ = strconv.ParseInt(r, 10, 64)
	}

	if c, ok := os.LookupEnv(concurrencyKey); ok {
		concurrency, _ = strconv.Atoi(c)
	}
}

/* Reuse the runner in Kitex-runBenchmark to record the result */

// Client is a client for runBenchmark.
type Client struct {
	cli     greetservice.Client
	reqPool *sync.Pool
}

func NewBenchmarkClient(cli greetservice.Client) *Client {
	return &Client{
		cli: cli,
		reqPool: &sync.Pool{
			New: func() interface{} {
				return &proxyless.HelloRequest{Message: "Hello!"}
			},
		},
	}
}

func (r *Client) Echo(action, msg string) (err error) {
	req := r.reqPool.Get().(*proxyless.HelloRequest)
	defer r.reqPool.Put(req)

	req.Message = msg
	_, err = r.cli.SayHello1(context.Background(), req)
	return err
}

func runBenchmark(name string, cli greetservice.Client) {
	bcli := NewBenchmarkClient(cli)
	r := runner.NewRunner()
	payload := string(make([]byte, echoSize))
	handler := func() error { return bcli.Echo("", payload) }

	// === warmup ===
	r.Warmup(handler, concurrency, 100*100)

	// === benching ===
	recorder := perf.NewRecorder(fmt.Sprintf("%s@Client", name))
	recorder.Begin()
	r.Run(name, handler, concurrency, totalReq, echoSize, 0)

	// == ending ===
	recorder.End()

	// === reporting ===
	recorder.Report() // report client
	fmt.Printf("\n\n")
}

type Runner struct {
	targetService string
}

func NewBenchmarkRunner(target string) *Runner {
	return &Runner{targetService: target}
}

func (r *Runner) directBenchmark() {
	cli, err := greetservice.NewClient(
		r.targetService,
		client.WithResolver(dns.NewDNSResolver()),
		client.WithLongConnection(
			connpool.IdleConfig{MaxIdlePerAddress: 1000, MaxIdleGlobal: 1000, MaxIdleTimeout: time.Minute},
		),
	)
	if err != nil {
		if err != nil {
			klog.Error("[direct runBenchmark] construct client error: %v\n", err)
			return
		}
	}
	klog.Info("Start Direct (DNS) Benchmark")
	for {
		runBenchmark(meshModeDirect, cli)
	}
}

func (r *Runner) sidecarBenchmark() {
	cli, err := greetservice.NewClient(
		r.targetService,
		client.WithHostPorts(r.targetService),
		client.WithLongConnection(
			connpool.IdleConfig{MaxIdlePerAddress: 1000, MaxIdleGlobal: 1000, MaxIdleTimeout: time.Minute},
		),
	)
	if err != nil {
		if err != nil {
			klog.Error("[sidecar runBenchmark] construct client error: %v\n", err)
			return
		}
	}
	klog.Info("Start Sidecar Benchmark")
	for {
		runBenchmark(meshModeProxy, cli)
	}
}

func (r *Runner) proxylessBenchmark() {
	err := xds.Init()
	if err != nil {
		klog.Error("[proxyless runBenchmark] xds init error: %v\n", err)
		return
	}
	cs := kxds.ClientSuite{
		RouterMiddleware: xdssuite.NewXDSRouterMiddleware(),
		Resolver:         xdssuite.NewXDSResolver(),
	}
	cli, err := greetservice.NewClient(
		r.targetService,
		client.WithXDSSuite(cs),
		client.WithLongConnection(
			connpool.IdleConfig{MaxIdlePerAddress: 1000, MaxIdleGlobal: 1000, MaxIdleTimeout: time.Minute},
		),
	)
	if err != nil {
		klog.Error("[proxyless] construct client error: %v\n", err)
		return
	}
	klog.Info("Start Proxyless Benchmark")
	for {
		runBenchmark(meshModeProxyless, cli)
	}
}

func (r *Runner) Run() error {
	meshMode, ok := os.LookupEnv(meshModeKey)
	if !ok {
		return fmt.Errorf("please specify the mesh mode")
	}
	// load the config from ENV
	initEnv()

	switch meshMode {
	case meshModeProxyless:
		r.proxylessBenchmark()
	case meshModeProxy:
		r.sidecarBenchmark()
	case meshModeDirect:
		r.directBenchmark()
	}
	return nil
}
