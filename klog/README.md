# klog
English | [中文](./README_CN.md)

You can learn about how to use Kitex klog

1. By default, the logger implemented by kitex is used by default.
   The default logger output can also be redirected using the `SetOutput` interface, and subsequent middleware and other parts of the framework can use the global methods in klog to output logs.
2. Kitex provides `SetLogger` interface to allow injection of your own logger.

## Interface Definition

In Kitex, the interfaces Logger, CtxLogger, FormatLogger are defined in pkg/klog, and these interfaces are used to output logs in different ways, and a Control interface is defined to control the logger. If you’d like to inject your own logger implementation, you must implement all the above interfaces (i.e. FullLogger). Kitex already provides a default implementation of FullLogger.

```go
// FullLogger is the combination of Logger, FormatLogger, CtxLogger and Control.
type FullLogger interface {
   Logger
   FormatLogger
   CtxLogger
   Control
}
```

## Log Extension

Extension provides the use of zap , logrus and slog

klog extension are coupled in the otel extension

For details on log extension, [see](https://cloudwego.cn/zh/docs/kitex/tutorials/observability/logging/)