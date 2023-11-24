// Copyright 2023 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"os"

	"github.com/cloudwego/kitex-examples/kitex_gen/api"
	"github.com/cloudwego/kitex-examples/kitex_gen/api/echo"
	"github.com/cloudwego/kitex/pkg/klog"
)

var _ api.Echo = &EchoImpl{}

// EchoImpl implements the last service interface defined in the IDL.
type EchoImpl struct{}

// Echo implements the Echo interface.
func (s *EchoImpl) Echo(ctx context.Context, req *api.Request) (resp *api.Response, err error) {
	klog.Info("echo called")
	return &api.Response{Message: req.Message}, nil
}

func main() {
	svr := echo.NewServer(new(EchoImpl))

	klog.SetLevel(klog.LevelDebug)
	f, err := os.Create("kitex.log")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// SetOutput sets the output of default logger. By default, it is stderr.
	klog.SetOutput(f)
	// if you want to output the log to the file and the stdout at the same time, you can use the following codes

	// fileWriter := io.MultiWriter(f, os.Stdout)
	// klog.SetOutput(fileWriter)

	if err := svr.Run(); err != nil {
		klog.Error("server stopped with error:", err)
	} else {
		klog.Debug("server stopped")
	}
}
