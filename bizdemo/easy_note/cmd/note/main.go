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
	"fmt"
	"io"
	"net"

	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/pkg/constants"
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/pkg/control"
	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/pkg/middleware"

	"github.com/cloudwego/kitex/pkg/klog"

	"github.com/opentracing/opentracing-go"

	"github.com/cloudwego/kitex-examples/bizdemo/easy_note/cmd/note/dal"
	note "github.com/cloudwego/kitex-examples/bizdemo/easy_note/kitex_gen/notedemo/noteservice"
	"github.com/cloudwego/kitex/pkg/acl"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

func InitJaeger(service string) (server.Suite, io.Closer) {
	cfg, _ := jaegercfg.FromEnv()
	cfg.ServiceName = service
	tracer, closer, err := cfg.NewTracer(jaegercfg.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	opentracing.InitGlobalTracer(tracer)
	return trace.NewDefaultServerSuite(), closer
}

func Init() {
	dal.Init()
}

func main() {
	tracer, closer := InitJaeger(constants.NoteServiceName)
	defer closer.Close()

	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"}) // r should not be reused.
	if err != nil {
		panic(err)
	}

	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8888")

	svr := note.NewServer(new(NoteServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.NoteServiceName}), // server name
		server.WithMiddleware(middleware.CommonMiddleware),                                             // middleWare
		server.WithMiddleware(middleware.ServerMiddleware),
		server.WithServiceAddr(addr),                                                     // address
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}),               // limit
		server.WithMuxTransport(),                                                        // Multiplex
		server.WithSuite(tracer),                                                         // tracer
		server.WithMiddleware(acl.NewACLMiddleware([]acl.RejectFunc{control.MemReject})), // access_control
		server.WithRegistry(r),                                                           // registry
	)
	Init()
	err = svr.Run()
	if err != nil {
		klog.Fatal(err)
	}
}