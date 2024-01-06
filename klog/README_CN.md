# klog
[English](./README.md) | 中文

您可以了解如何使用 Kitex klog

1. 默认情况下，默认使用 kitex 实现的 logger。
默认的 logger 输出也可以使用 `SetOutput` 接口进行重定向，后续的中间件和框架的其他部分可以使用 klog 中的全局方法来输出日志。
2. Kitex 提供 `SetLogger` 接口来允许注入您自己的 logger。

## 接口定义

Kitex 中，pkg/klog 中定义了 Logger、CtxLogger、FormatLogger 接口，这些接口用于以不同的方式输出日志，并定义了一个 Control 接口来控制 logger。 如果您想注入自己的记录器实现，则必须实现上述所有接口（即 FullLogger）。Kitex 已经提供了 FullLogger 的默认实现。

```go
// FullLogger 是 Logger、FormatLogger、CtxLogger 和 Control 的组合。
type FullLogger interface {
    Logger
    FormatLogger
    CtxLogger
    Control
}
````

## 日志扩展

扩展提供了 zap 、logrus 和 slog 的使用

klog 拓展耦合在 otel 的拓展里

关于日志扩展的详细信息，[参见](https://cloudwego.cn/zh/docs/kitex/tutorials/observability/logging/)