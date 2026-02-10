# Kitex 的 Prometheus 监控

[English](./README.md) | 中文

## 使用示例

### 简单模式

#### 服务端

请参阅 [simple/server](./simple/server)

#### 客户端

请参阅 [simple/client](./simple/client)

### 自定义模式

#### 服务端

请参阅 [custom/server](./custom/server)

#### 客户端

请参阅 [custom/client](./custom/client)

## 如何运行

1.  安装 docker 并启动 docker
2.  在 prometheus.yml 中将`{{ INET_IP }}`更改为本地 ip
3.  运行 Prometheus 和 Grafana
    `docker compose up`
4.  运行 Kitex 服务端
    `go run simple/server/main.go`
5.  运行 Kitex 客户端
    `go run simple/client/main.go`
6.  访问`http://localhost:3000`，账号密码默认为`admin`
7.  配置 Prometheus 数据源
    1.  `Connection`
    2.  `Add new connection`
    3.  选择 `Prometheus` 并填写 URL 为 `http://prometheus:9090`
    4.  配置完成后点击 `Save & Test` 测试是否有效

8.  添加仪表板 `Create` -> `dashboard`，根据需要添加吞吐量、pct99 等监控指标。

    例如：
    - 成功请求的服务器吞吐量

    `sum(rate(kitex_server_throughput{statusCode="200"}[1m])) by (method)`
    - 成功请求的服务器延迟 pct99

    `histogram_quantile(0.9,sum(rate(kitex_server_latency_us_bucket{status="succeed"}[1m]))by(le)) / 1000`

有关 kitex 监控的更多信息，
请点击 [monitoring](https://www.cloudwego.io/zh/docs/kitex/tutorials/observability/monitoring/)
