# Kitex Examples

[English](README.md) | 中文

## 如何运行

您可以进入相关示例以获取有关“如何运行”的信息

## Bizdemo
- [bizdemo/kitex_gorm](bizdemo/kitex_gorm) 使用 Kitex 和 gorm 的示例
- [bizdemo/kitex_gorm_gen](bizdemo/kitex_gorm_gen) 使用 Kitex 和 gorm_gen 的示例
- [bizdemo/kitex_zorm](bizdemo/kitex_zorm) 使用 Kitex 和 zorm 的示例
- [bizdemo/kitex_ent](bizdemo/kitex_ent) 使用 Kitex 和 ent 的示例
- [bizdemo/easy_note](bizdemo/easy_note) 使用 Kitex 作为 rpc 框架实现多中间件简易笔记服务的示例
- [Bookinfo](https://github.com/cloudwego/biz-demo/tree/main/bookinfo) 使用 Kitex 作为 rpc 框架实现包含 istio 服务网格，OpenTelemetry 监控等多功能书籍服务程序的示例
- [Open Payment Platform](https://github.com/cloudwego/biz-demo/tree/main/open-payment-platform) 使用 Kitex 作为 rpc 框架演示了 Kitex 泛化调用的用法，构建了一体化支付平台
- [Book Shop](https://github.com/cloudwego/biz-demo/tree/main/book-shop) 使用 Kitex 作为 rpc 框架实现包含 ElasticSearch 搜索引擎的电子商务系统的示例
- [FreeCar](https://github.com/CyanAsterisk/FreeCar) 使用 Kitex 作为 rpc 框架实现云原生分时租车系统套件服务的示例

## 基本特性
- [basic](basic) Kitex 的基础示例
- [async_call](async_call) 在 Kitex server 中使用异步调用的示例
- [codec](codec) Kitex 服务端和客户端使用自定义编解码器的示例
  - [codec-Dubbo](https://github.com/kitex-contrib/codec-dubbo/tree/main/samples/helloworld) Kitex 为了支持 Kitex <-> Dubbo 互通 推出的 Dubbo 协议编解码器的示例
- [long_connection](longconnection) Kitex 服务端和客户端使用长连接的示例
- [streaming](streaming) Kitex 服务端和客户端使用流式调用的示例（基于 GRPC/HTTP2，使用 Protobuf IDL）
- [thrift streaming](thrift_streaming) Kitex 服务端和客户端使用 Thrift 流式调用的示例 (基于 GRPC/HTTP2，使用 Thrift IDL)

- [business_exception](business_exception) Kitex 服务端和客户端使用业务异常的示例
- [middleware](middleware) Kitex 服务端和客户端使用中间件的示例

## 治理特性
- Kitex 服务端使用配置中心对接治理特性的示例
  - [etcd](https://github.com/kitex-contrib/config-etcd/tree/main/example) Kitex 服务端使用 etcd 作为配置中心对接治理特性的示例
  - [nacos](https://github.com/kitex-contrib/config-nacos/tree/main/example) Kitex 服务端使用 nacos 作为配置中心对接治理特性的示例
  - [apollo](https://github.com/kitex-contrib/config-apollo/tree/main/example) Kitex 服务端使用 apollo 作为配置中心对接治理特性的示例
- [discovery](discovery) Kitex 服务端和客户端使用服务注册与发现的示例
  - [etcd](https://github.com/kitex-contrib/registry-etcd/tree/main/example) Kitex 服务端和客户端使用 etcd 作为服务注册中心的示例
  - [nacos](https://github.com/kitex-contrib/registry-nacos/tree/main/example) Kitex 服务端和客户端使用 nacos 作为服务注册中心的示例
  - [polaris](https://github.com/kitex-contrib/registry-polaris/tree/main/example) Kitex 服务端和客户端使用 polaris 作为服务注册中心的示例
  - [zookeeper](https://github.com/kitex-contrib/registry-zookeeper) Kitex 服务端和客户端使用 zookeeper 作为服务注册中心的示例
  - [consul](https://github.com/kitex-contrib/registry-consul/tree/main/example) Kitex 服务端和客户端使用 consul 作为服务注册中心的示例
  - [servicecomb](https://github.com/kitex-contrib/registry-servicecomb/tree/main/example) Kitex 服务端和客户端使用 servicecomb 作为服务注册中心的示例
  - [eureka](https://github.com/kitex-contrib/registry-eureka/tree/main/example) Kitex 服务端和客户端使用 eureka 作为服务注册中心的示例
  - [dns](https://github.com/kitex-contrib/resolver-dns) Kitex 服务端和客户端使用 dns 进行服务发现的示例
  - [resolver_rule_based](https://github.com/kitex-contrib/resolver-rule-based/tree/main/demo) 为 Kitex 提供了一个基于规则的解析器。它允许用户在服务发现中配置规则来过滤服务实例，实现流量切分的功能。
- [timeout](governance/timeout) Kitex 服务端和客户端使用超时控制的示例
- [limit](governance/limit) Kitex 服务端使用限流的示例
- [circuit_breaker](governance/circuitbreak) Kitex 客户端使用熔断的示例
- [retry](governance/retry) Kitex 客户端使用重试的示例
- [load_balance](loadbalancer) Kitex 服务端和客户使用负载均衡的示例

## 可观测性
- [opentelemetry](opentelemetry) Kitex 服务端和客户端使用 OpenTelemetry 的示例
- [prometheus](prometheus) Kitex 服务端和客户端使用 prometheus 的示例
- [tracer](tracer) Kitex 服务端和客户端使用 tracer 的示例
- [klog](klog) Kitex 服务端使用 klog 日志的示例

## 高级特性
- [frugal](frugal) Kitex 服务端和客户端使用 frugal 的示例
- [grpc_proxy](grpcproxy) Kitex 服务端和客户端使用 grpc_proxy 的示例
- [generic](generic) Kitex 服务端和客户端使用泛化调用的示例
- [meta_info](metainfo) Kitex 服务端和客户端使用元信息的示例
- [server_hook](server_hook) Kitex 服务端 启动/退出 前后定制业务逻辑的示例
- [server_sdk](server_sdk) Kitex 服务端 SDK 化的示例
- [profiler](profiler) Kitex 服务端和客户端使用请求成本度量进行性能分析的示例
- [proxyless](proxyless) 让 Kitex 服务以 Proxyless 的模式运行，被服务网格统一纳管的示例
- [grpc_multi_service](grpc_multi_service) Kitex 服务端和客户端使用 grpc 多服务的示例
- [thrift_multi_service](thrift_multi_service) Kitex 服务端和客户端使用 thrift (non-streaming) 多服务的示例
- [protobuf_multi_service](protobuf_multi_service) Kitex 服务端和客户端使用 protobuf (non-streaming) 多服务的示例
- [goruntine_local_storage](goroutine-local-storage) Kitex 服务端和客户端使用 goruntine_local_storage 的示例

## Kitex 生成代码
- [protobuf](kitex/protobuf) 使用 Kitex 与 protobuf 生成服务端代码的示例
- [template](kitex/template) 使用 Kitex 自定义模版生成服务端代码的示例
- [thrift](kitex/thrift) 使用 Kitex 与 thrift 生成服务端代码的示例

## Note

执行示例的所有命令都应在 kitex-examples 下执行。
