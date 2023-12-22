# Kitex Examples

English | [中文](README_CN.md)

## How to run

You can go into the related examples for information on "how to run"

## Bizdemo
- [bizdemo/easy_note](bizdemo/easy_note) Example of using kitex as an rpc framework to implement a simple note service with multiple middlewares
- [Bookinfo](https://github.com/cloudwego/biz-demo/tree/main/bookinfo) Example of using kitex as the rpc framework to implement multi-functional book service programs including istio, opentelemetry, etc.
- [Open Payment Platform](https://github.com/cloudwego/biz-demo/tree/main/open-payment-platform) Example of using kitex as the rpc framework to demonstrate the usage of kitex generic call, and builds an integrated payment platform.
- [Book Shop](https://github.com/cloudwego/biz-demo/tree/main/book-shop) Example of using kitex as an rpc framework to implement an e-commerce system including an ElasticSearch search engine
- [FreeCar](https://github.com/CyanAsterisk/FreeCar) Example of using kitex as the rpc framework to implement cloud native time-sharing car rental system suite services

## Basic features
- [basic](basic) Basic example of kitex
- [async_call](async_call) Example of using asynchronous call in kitex server
- [codec](codec) Example of kitex server and client using custom codecs
    - [codec-dubbo](https://github.com/kitex-contrib/codec-dubbo/tree/main/samples/helloworld) Example of dubbo protocol codec launched by Kitex to support kitex <-> dubbo interoperability
- [long_connection](longconnection) Kitex example of using long connections between server and client
- [streaming](streaming) Example of kitex server and client using streams
- [middleware](middleware) Kitex Example of using middleware between server and client

## Governance Features
- Example of kitex server using configuration center to connect to governance features
    - [etcd](https://github.com/kitex-contrib/config-etcd/tree/main/example) Kitex server uses etcd as an example of connecting the configuration center with governance features
    - [nacos](https://github.com/kitex-contrib/config-nacos/tree/main/example) Kitex server uses nacos as an example of connecting the configuration center with governance features
    - [apollo](https://github.com/kitex-contrib/config-apollo/tree/main/example) Kitex server uses apollo as an example of connecting the configuration center with governance features
- [discovery](discovery) Example of kitex server and client using service registration and discovery
    - [etcd](https://github.com/kitex-contrib/registry-etcd/tree/main/example) Example of kitex server and client using etcd as service registration center
    - [nacos](https://github.com/kitex-contrib/registry-nacos/tree/main/example) Kitex server and client use nacos as an example of service registration center
    - [polaris](https://github.com/kitex-contrib/registry-polaris/tree/main/example) Kitex server and client use polaris as an example of service registration center
    - [zookeeper](https://github.com/kitex-contrib/registry-zookeeper) Example of kitex server and client using zookeeper as service registration center
    - [consul](https://github.com/kitex-contrib/registry-consul/tree/main/example) Example of kitex server and client using consul as service registration center
    - [servicecomb](https://github.com/kitex-contrib/registry-servicecomb/tree/main/example) Kitex server and client use servicecomb as an example of service registration center
    - [eureka](https://github.com/kitex-contrib/registry-eureka/tree/main/example) Kitex server and client use eureka as an example of service registration center
    - [dns](https://github.com/kitex-contrib/resolver-dns) Example of kitex server and client using dns for service discovery
    - [resolver_rule_based](https://github.com/kitex-contrib/resolver-rule-based/tree/main/demo) provides a rule-based resolver for kitex. It allows users to configure rules in service discovery to filter service instances and implement traffic segmentation.
- [timeout](governance/timeout) Example of using timeout control between kitex server and client
- [limit](governance/limit) Example of using current limit on kitex server
- [circuit_breaker](governance/circuitbreak) Example of kitex client using circuit breaker
- [retry](governance/retry) Example of kitex client using retry
- [load_balance](loadbalancer) Example of kitex server and client using load balancing

## Observability
- [opentelemetry](opentelemetry) Example of kitex server and client using opentelemetry
- [prometheus](prometheus) Example of kitex server and client using prometheus
- [tracer](tracer) Example of using tracer on kitex server and client
- [klog](klog) Example of using klog logs on the kitex server

## Advanced features
- [frugal](frugal) Kitex example of using frugal on the server and client
- [grpc_proxy](grpcproxy) Example of kitex server and client using grpc_proxy
- [generic](generic) kitex example of using generic calls between server and client
- [meta_info](metainfo) Example of kitex server and client using meta information
- [profiler](profiler) Example of kitex server and client performance analysis using request cost metrics
- [proxyless](proxyless) Example of letting Kitex services run in proxyless mode and be managed uniformly by the service mesh

## Kitex generated code
- [protobuf](kitex/protobuf) Example of using kitex and protobuf to generate server code
- [template](kitex/template) Example of using kitex custom template to generate server code
- [thrift](kitex/thrift) Example of using kitex and thrift to generate server code
- [protobuf](kitex/protobuf) Example of using kitex and protobuf to generate server code

## Note

All commands to execute the example should be executed under "kitex-example".