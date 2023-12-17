# Kitex Examples

[English](README.md) | 中文

## 如何运行

您可以进入相关示例以获取有关“如何运行”的信息

## Bizdemo
- [bizdemo/easy_note](bizdemo/easy_note) 使用 kitex 作为 rpc 框架实现多中间键简易笔记服务的示例
- [Bookinfo](https://github.com/cloudwego/biz-demo/tree/main/bookinfo) 使用 kitex 作为 rpc 框架实现包含 istio 服务网格，opentelemetry 监控等多功能书籍服务程序的示例
- [Open Payment Platform](https://github.com/cloudwego/biz-demo/tree/main/open-payment-platform) 使用 kitex 作为 rpc 框架演示了 kitex 泛化调用的用法，构建了一体化支付平台
- [Book Shop](https://github.com/cloudwego/biz-demo/tree/main/book-shop) 使用 kitex 作为 rpc 框架实现包含 ElasticSearch 搜索引擎的电子商务系统的示例
- [FreeCar](https://github.com/CyanAsterisk/FreeCar) 使用 kitex 作为 rpc 框架实现云原生分时租车系统套件服务的示例

## 基本特性
- [basic](basic) kitex 的基础示例
- [async_call](async_call) 在 kitex server 中使用异步调用的示例
- [codec](codec) kitex 服务端和客户端使用自定义编解码器的示例
  - [codec-dubbo](https://github.com/kitex-contrib/codec-dubbo/tree/main/samples/helloworld) Kitex 为了支持 kitex <-> dubbo 互通 推出的 dubbo 协议编解码器的示例
- [long_connection](longconnection) kitex 服务端和客户端使用长连接的示例
- [streaming](streaming) kitex 服务端和客户端使用流式调用的示例
- [middleware](middleware) kitex 服务端和客户端使用中间件的示例

## 治理特性
- kitex 服务端使用配置中心对接治理特性的示例
  - [etcd](https://github.com/kitex-contrib/config-etcd/tree/main/example) kitex 服务端使用 etcd 作为配置中心对接治理特性的示例
  - [nacos](https://github.com/kitex-contrib/config-nacos/tree/main/example) kitex 服务端使用 nacos 作为配置中心对接治理特性的示例
  - [apollo](https://github.com/kitex-contrib/config-apollo/tree/main/example) kitex 服务端使用 apollo 作为配置中心对接治理特性的示例
- [discovery](discovery) kitex 服务端和客户端使用服务注册与发现的示例
  - [etcd](https://github.com/kitex-contrib/registry-etcd/tree/main/example) kitex 服务端和客户端使用 etcd 作为服务注册中心的示例
  - [nacos](https://github.com/kitex-contrib/registry-nacos/tree/main/example) kitex 服务端和客户端使用 nacos 作为服务注册中心的示例
  - [polaris](https://github.com/kitex-contrib/registry-polaris/tree/main/example) kitex 服务端和客户端使用 polaris 作为服务注册中心的示例
  - [zookeeper](https://github.com/kitex-contrib/registry-zookeeper) kitex 服务端和客户端使用 zookeeper 作为服务注册中心的示例
  - [consul](https://github.com/kitex-contrib/registry-consul/tree/main/example) kitex 服务端和客户端使用 consul 作为服务注册中心的示例
  - [servicecomb](https://github.com/kitex-contrib/registry-servicecomb/tree/main/example) kitex 服务端和客户端使用 servicecomb 作为服务注册中心的示例
  - [eureka](https://github.com/kitex-contrib/registry-eureka/tree/main/example) kitex 服务端和客户端使用 eureka 作为服务注册中心的示例
  - [dns](https://github.com/kitex-contrib/resolver-dns) kitex 服务端和客户端使用 dns 进行服务发现的示例
  - [resolver_rule_based](https://github.com/kitex-contrib/resolver-rule-based/tree/main/demo) 为 kitex 提供了一个基于规则的解析器。它允许用户在服务发现中配置规则来过滤服务实例，实现流量切分的功能。
- [timeout](governance/timeout) kitex 服务端和客户端使用超时控制的示例
- [limit](governance/limit) kitex 服务端使用限流的示例
- [circuit_breaker](governance/circuitbreak) kitex 客户端使用熔断的示例
- [retry](governance/retry) kitex 客户端使用重试的示例
- [load_balance](loadbalancer) kitex 服务端和客户使用负载均衡的示例

## 可观测性
- [opentelemetry](opentelemetry) kitex 服务端和客户端使用 opentelemetry 的示例
- [prometheus](prometheus) kitex 服务端和客户端使用 prometheus 的示例
- [tracer](tracer) kitex 服务端和客户端使用 tracer 的示例
- [klog](klog) kitex 服务端使用 klog 日志的示例

## 高级特性
- [frugal](frugal) kitex 服务端和客户端使用 frugal 的示例
- [grpc_proxy](grpcproxy) kitex 服务端和客户端使用 grpc_proxy 的示例
- [generic](generic) kitex 服务端和客户端使用泛化调用的示例
- [meta_info](metainfo) kitex 服务端和客户端使用元信息的示例
- [profiler](profiler) kitex 服务端和客户端使用请求成本度量进行性能分析的示例
- [proxyless](proxyless) 让 Kitex 服务以 Proxyless 的模式运行，被服务网格统一纳管的示例

## Kitex 生成代码
- [protobuf](kitex/protobuf) 使用 kitex 与 protobuf 生成服务端代码的示例
- [template](kitex/template) 使用 kitex 自定义模版生成服务端代码的示例

## Note

执行示例的所有命令都应在 kitex-examples 下执行。