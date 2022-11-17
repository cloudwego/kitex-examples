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
	"log"
	"math"
	"strconv"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/profiler"
	"github.com/cloudwego/kitex/pkg/remote"
	"github.com/cloudwego/kitex/server"

	"github.com/cloudwego/kitex-examples/kitex_gen/api"
	"github.com/cloudwego/kitex-examples/kitex_gen/api/echo"
)

var _ api.Echo = &EchoImpl{}

// EchoImpl implements the last service interface defined in the IDL.
type EchoImpl struct{}

// Echo implements the Echo interface.
func (s *EchoImpl) Echo(ctx context.Context, req *api.Request) (resp *api.Response, err error) {
	sum := 0
	for i := 0; i < 10000000; i++ {
		sum++
	}
	return &api.Response{Message: req.Message}, nil
}

func LogProcessor(profiles []*profiler.TagsProfile) error {
	if len(profiles) == 0 {
		return nil
	}
	klog.Infof("KITEX: profiler collect %d records", len(profiles))
	for _, p := range profiles {
		klog.Infof("KITEX: profiler - %s %.2f%% %d", p.Key, p.Percent*100, p.Value)
	}
	klog.Info("---------------------------------")
	return nil
}

type ArgsGetter interface {
	GetReq() *api.Request
}

var sizeClassTagging remote.MessageTagging = func(ctx context.Context, msg remote.Message) (context.Context, []string) {
	if data := msg.Data(); data == nil {
		return ctx, nil
	}
	var tags = make([]string, 0, 2)
	var reqType int32
	if args, ok := msg.Data().(ArgsGetter); ok {
		if req := args.GetReq(); req != nil {
			sizeClass := int(math.Floor(math.Log2(float64(len(req.Message)))))
			tags = append(tags, "size_class", strconv.Itoa(sizeClass))
		}
	}
	// if you don't need to get the tags after middleware, not need to change ctx
	return context.WithValue(ctx, "size_class", reqType), tags
}

func main() {
	interval := time.Duration(0)
	window := 1 * time.Second
	pr := profiler.NewProfiler(LogProcessor, interval, window)
	svr := echo.NewServer(
		new(EchoImpl),
		server.WithProfiler(pr),
		server.WithProfilerMessageTagging(sizeClassTagging),
	)
	if err := svr.Run(); err != nil {
		log.Println("server stopped with error:", err)
	} else {
		log.Println("server stopped")
	}
}
