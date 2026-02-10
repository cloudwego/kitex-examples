# Opentelemetry 示例

[English](./README.md) | 中文

## 如何运行

### Docker

确保 Docker 已安装。

### 运行 opentelemetry-collector、jaeger、victoriametrics、grafana

```
docker compose up -d
```

### 运行 Kitex 服务端

```
go run server/main.go
```

### 运行 Kitex 客户端

```
go run client/main.go
```

## 监控

### 查看链路追踪

你可以前往 http://localhost:16686 来访问 Jaeger UI。（详情可访问Monitor Jaeger）
![img.png](static/jaeger.png)

### 查看指标

你可以前往 http://localhost:3000 来访问 Grafana UI。（您可以访问 Monitor Grafana 了解指标）

#### 添加数据源

HTTP 网址：

```
http://victoriametrics:8428/
```

![img_1.png](static/grafana.png)

#### 添加仪表板和面板

![img.png](static/panel.png)

#### 支持指标

- RPC Metrics
- Runtime Metrics

## 链路追踪相关日志

### 设置 logger 实现

```go
import (
    kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
)

func init()  {
    klog.SetLogger(kitexlogrus.NewLogger())
    klog.SetLevel(klog.LevelDebug)

}
```

### 日志与 Context 关联

```go
// Echo implements the Echo interface.
func (s *EchoImpl) Echo(ctx context.Context, req *api.Request) (resp *api.Response, err error) {
	klog.CtxDebugf(ctx, "echo called: %s", req.GetMessage())
	return &api.Response{Message: req.Message}, nil
}
```

### 查看日志

```log
{"level":"debug","msg":"echo called: my request 1","span_id":"056e0cf9a8b2cec3","time":"2022-03-09T02:47:28+08:00","trace_flags":"01","trace_id":"33bdd3c81c9eb6cbc0fbb59c57ce088b"}
```

## 通过 Jaeger 进行工作

> [在 Jaeger 中引入对 OpenTelemetry 的原生支持](https://medium.com/jaegertracing/introducing-native-support-for-opentelemetry-in-jaeger-eb661be8183c)

Jaeger原生支持OTLP协议，我们可以直接向Jaeger发送数据，无需 OpenTelemetry Collector

### Jaeger 架构

> 图片来自 [jaeger](https://github.com/jaegertracing/jaeger)

![img.png](jaeger-arch/img.png)

### Demo

#### 使用 COLLECTOR_OTLP_ENABLED 运行 Jaeger

```yaml
version: "3.7"
services:
  # Jaeger
  jaeger-all-in-one:
    image: jaegertracing/all-in-one:latest
    environment:
      - COLLECTOR_OTLP_ENABLED=true
    ports:
      - "4317:4317" # OTLP gRPC receiver
```

#### 带有环境的配置导出器

```yaml
export OTEL_EXPORTER_OTLP_ENDPOINT=http://host.docker.internal:4317
export OTEL_EXPORTER_OTLP_PROTOCOL=grpc
```

#### 运行示例应用程序并查看 Jaeger

![img.png](static/jaeger-otlp.png)
