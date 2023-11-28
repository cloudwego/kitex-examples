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
	"fmt"
	"log"
	"os"
	"path"
	"runtime"
	"time"

	"github.com/cloudwego/kitex-examples/kitex_gen/api"
	"github.com/cloudwego/kitex-examples/kitex_gen/api/echo"
	"github.com/cloudwego/kitex/pkg/klog"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
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

	// Customizable output directory.
	var logFilePath string
	dir := "./klog"
	logFilePath = dir + "/logs/"
	if err := os.MkdirAll(logFilePath, 0o777); err != nil {
		log.Println(err.Error())
		return
	}

	// Set filename to date
	logFileName := time.Now().Format("2006-01-02") + ".log"
	fileName := path.Join(logFilePath, logFileName)
	if _, err := os.Stat(fileName); err != nil {
		if _, err := os.Create(fileName); err != nil {
			log.Println(err.Error())
			return
		}
	}

	logger := kitexlogrus.NewLogger()
	logger.Logger().SetReportCaller(true)
	// klog will warp a layer of logrus, so you need to calculate the depth of the caller file separately.
	logger.Logger().AddHook(NewCustomHook(10))
	// Provides compression and deletion
	lumberjackLogger := &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    20,   // A file can be up to 20M.
		MaxBackups: 5,    // Save up to 5 files at the same time.
		MaxAge:     10,   // A file can exist for a maximum of 10 days.
		Compress:   true, // Compress with gzip.
	}

	logger.SetOutput(lumberjackLogger)
	logger.SetLevel(klog.LevelDebug)
	// if you want to output the log to the file and the stdout at the same time, you can use the following codes

	// fileWriter := io.MultiWriter(lumberjackLogger, os.Stdout)
	// logger.SetOutput(fileWriter)
	klog.SetLogger(logger)

	if err := svr.Run(); err != nil {
		klog.CtxErrorf(context.Background(), "server stopped with error:", err)
	} else {
		klog.CtxDebugf(context.Background(), "server stopped")
	}
}

// CustomHook Custom Hook for processing logs
type CustomHook struct {
	CallerDepth int
}

func NewCustomHook(depth int) *CustomHook {
	return &CustomHook{
		CallerDepth: depth,
	}
}

func (hook *CustomHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (hook *CustomHook) Fire(entry *logrus.Entry) error {
	// Get caller information and specify depth
	pc, file, line, ok := runtime.Caller(hook.CallerDepth)
	if ok {
		funcName := runtime.FuncForPC(pc).Name()
		entry.Data["caller"] = fmt.Sprintf("%s:%d %s", file, line, funcName)
	}

	return nil
}
