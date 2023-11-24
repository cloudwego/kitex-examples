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
	"io"
	"os"
	"path"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/sirupsen/logrus"
)

type Logger struct {
	l *logrus.Logger
}

func NewLogger() *Logger {
	var logFilePath string
	if dir, err := os.Getwd(); err == nil {
		logFilePath = dir + "/logs/"
	}
	if err := os.MkdirAll(logFilePath, 0o777); err != nil {
		fmt.Println(err.Error())
	}
	// Set filename to date
	logFileName := time.Now().Format("2006-01-02") + ".log"
	fileName := path.Join(logFilePath, logFileName)
	if _, err := os.Stat(fileName); err != nil {
		if _, err := os.Create(fileName); err != nil {
			fmt.Println(err.Error())
		}
	}

	// Write to file
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		fmt.Println(err.Error())
	}

	// Instantiate
	logger := logrus.New()
	// Set output
	logger.Out = src
	// Set log level
	logger.SetLevel(logrus.DebugLevel)
	// Set log format
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	return &Logger{l: logger}
}

func (l *Logger) Trace(v ...interface{}) {
	l.l.Trace(v...)
}

func (l *Logger) Debug(v ...interface{}) {
	l.l.Debug(v...)
}

func (l *Logger) Info(v ...interface{}) {
	l.l.Info(v...)
}

func (l *Logger) Notice(v ...interface{}) {
	l.l.Warn(v...)
}

func (l *Logger) Warn(v ...interface{}) {
	l.l.Warn(v...)
}

func (l *Logger) Error(v ...interface{}) {
	l.l.Error(v...)
}

func (l *Logger) Fatal(v ...interface{}) {
	l.l.Fatal(v...)
}

func (l *Logger) Tracef(format string, v ...interface{}) {
	l.l.Tracef(format, v...)
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.l.Debugf(format, v...)
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.l.Infof(format, v...)
}

func (l *Logger) Noticef(format string, v ...interface{}) {
	l.l.Warnf(format, v...)
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	l.l.Warnf(format, v...)
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.l.Errorf(format, v...)
}

func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.l.Fatalf(format, v...)
}

func (l *Logger) CtxTracef(ctx context.Context, format string, v ...interface{}) {
	l.l.WithContext(ctx).Tracef(format, v...)
}

func (l *Logger) CtxDebugf(ctx context.Context, format string, v ...interface{}) {
	l.l.WithContext(ctx).Debugf(format, v...)
}

func (l *Logger) CtxInfof(ctx context.Context, format string, v ...interface{}) {
	l.l.WithContext(ctx).Infof(format, v...)
}

func (l *Logger) CtxNoticef(ctx context.Context, format string, v ...interface{}) {
	l.l.WithContext(ctx).Warnf(format, v...)
}

func (l *Logger) CtxWarnf(ctx context.Context, format string, v ...interface{}) {
	l.l.WithContext(ctx).Warnf(format, v...)
}

func (l *Logger) CtxErrorf(ctx context.Context, format string, v ...interface{}) {
	l.l.WithContext(ctx).Errorf(format, v...)
}

func (l *Logger) CtxFatalf(ctx context.Context, format string, v ...interface{}) {
	l.l.WithContext(ctx).Fatalf(format, v...)
}

func (l *Logger) SetLevel(level klog.Level) {
	var lv logrus.Level
	switch level {
	case klog.LevelTrace:
		lv = logrus.TraceLevel
	case klog.LevelDebug:
		lv = logrus.DebugLevel
	case klog.LevelInfo:
		lv = logrus.InfoLevel
	case klog.LevelWarn, klog.LevelNotice:
		lv = logrus.WarnLevel
	case klog.LevelError:
		lv = logrus.ErrorLevel
	case klog.LevelFatal:
		lv = logrus.FatalLevel
	default:
		lv = logrus.WarnLevel
	}
	l.l.SetLevel(lv)
}

func (l *Logger) SetOutput(writer io.Writer) {
	l.l.SetOutput(writer)
}
