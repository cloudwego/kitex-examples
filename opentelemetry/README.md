# Opentelemetry Example

English | [中文](./README_CN.md)

## How To Run

### Docker

Ensure Docker has been installed.

### Run opentelemetry-collector, jaeger, victoriametrics, grafana

```
docker compose up -d
```

### Run Kitex server

```
go run server/main.go
```

### Run Kitex client

```
go run client/main.go
```

## Monitoring

### View Trace

You can then navigate to http://localhost:16686 to access the Jaeger UI. (You can visit Monitor Jaeger for details)
![img.png](static/jaeger.png)

### View Metrics

You can then navigate to http://localhost:3000 to access the Grafana UI. (You can visit Monitor Grafana for metrics)

#### Add Datasource

Http URL:

```
http://victoriametrics:8428/
```

![img_1.png](static/grafana.png)

#### Add A Dashboard And A Panel

![img.png](static/panel.png)

#### Support Metrics

- RPC Metrics
- Runtime Metrics

## Tracing Associated Logs

### Set Logger Impl

```go
import (
    kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
)

func init()  {
    klog.SetLogger(kitexlogrus.NewLogger())
    klog.SetLevel(klog.LevelDebug)

}
```

### Log With Context

```go
// Echo implements the Echo interface.
func (s *EchoImpl) Echo(ctx context.Context, req *api.Request) (resp *api.Response, err error) {
	klog.CtxDebugf(ctx, "echo called: %s", req.GetMessage())
	return &api.Response{Message: req.Message}, nil
}
```

### view log

```log
{"level":"debug","msg":"echo called: my request 1","span_id":"056e0cf9a8b2cec3","time":"2022-03-09T02:47:28+08:00","trace_flags":"01","trace_id":"33bdd3c81c9eb6cbc0fbb59c57ce088b"}
```

## Work with Jaeger

> [Introducing native support for OpenTelemetry in Jaeger](https://medium.com/jaegertracing/introducing-native-support-for-opentelemetry-in-jaeger-eb661be8183c)

Jaeger natively supports OTLP protocol, and we can send data directly to Jaeger without OpenTelemetry Collector

### Jaeger Architecture

> Image from [jaeger](https://github.com/jaegertracing/jaeger)

![img.png](jaeger-arch/img.png)

### Demo

#### Run Jaeger with COLLECTOR_OTLP_ENABLED

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

#### Config Exporter with Environment

```yaml
export OTEL_EXPORTER_OTLP_ENDPOINT=http://host.docker.internal:4317
export OTEL_EXPORTER_OTLP_PROTOCOL=grpc
```

#### Run Exeample App and View Jaeger

![img.png](static/jaeger-otlp.png)
