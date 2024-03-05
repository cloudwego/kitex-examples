# Kitex Examples

English | [中文](README_CN.md)

## How to run

You can go into the related examples for information on "how to run"

## Bizdemo
- [bizdemo/kitex_gorm](bizdemo/kitex_gorm) Example of using Kitex and gorm
- [bizdemo/kitex_gorm_gen](bizdemo/kitex_gorm_gen) Example of using Kitex and gorm_gen
- [bizdemo/kitex_zorm](bizdemo/kitex_zorm) Example of using Kitex and zorm
- [bizdemo/kitex_ent](bizdemo/kitex_ent) Example of using Kitex and ent
- [bizdemo/easy_note](bizdemo/easy_note) Example of using Kitex as an rpc framework to implement a simple note service with multiple middlewares
- [Bookinfo](https://github.com/cloudwego/biz-demo/tree/main/bookinfo) Example of using Kitex as the rpc framework to implement multi-functional book service programs including istio, OpenTelemetry, etc.
- [Open Payment Platform](https://github.com/cloudwego/biz-demo/tree/main/open-payment-platform) Example of using Kitex as the rpc framework to demonstrate the usage of Kitex generic call, and builds an integrated payment platform.
- [Book Shop](https://github.com/cloudwego/biz-demo/tree/main/book-shop) Example of using Kitex as an rpc framework to implement an e-commerce system including an ElasticSearch search engine
- [FreeCar](https://github.com/CyanAsterisk/FreeCar) Example of using Kitex as the rpc framework to implement cloud native time-sharing car rental system suite services

## Basic features
- [basic](basic) Basic example of Kitex
- [async_call](async_call) Example of using asynchronous call in Kitex server
- [codec](codec) Example of Kitex server and client using custom codecs
    - [codec-Dubbo](https://github.com/kitex-contrib/codec-dubbo/tree/main/samples/helloworld) Example of Dubbo protocol codec launched by Kitex to support Kitex <-> Dubbo interoperability
- [long_connection](longconnection) Kitex example of using long connections between server and client
- [streaming](streaming) Example of Kitex server and client using streams (based on GRPC/HTTP2, using Protobuf IDL)
- [thrift streaming](thrift_streaming) Example of Kitex server and client using thrift streaming (based on GRPC/HTTP2, using Thrift IDL)
- [business_exception](business_exception) Example of Kitex server and client using business exceptions
- [middleware](middleware) Kitex Example of using middleware between server and client

## Governance Features
- Example of kitex server using configuration center to connect to governance features
    - [etcd](https://github.com/kitex-contrib/config-etcd/tree/main/example) Kitex server uses etcd as an example of connecting the configuration center with governance features
    - [nacos](https://github.com/kitex-contrib/config-nacos/tree/main/example) Kitex server uses nacos as an example of connecting the configuration center with governance features
    - [apollo](https://github.com/kitex-contrib/config-apollo/tree/main/example) Kitex server uses apollo as an example of connecting the configuration center with governance features
- [discovery](discovery) Example of Kitex server and client using service registration and discovery
    - [etcd](https://github.com/kitex-contrib/registry-etcd/tree/main/example) Example of Kitex server and client using etcd as service registration center
    - [nacos](https://github.com/kitex-contrib/registry-nacos/tree/main/example) Kitex server and client use nacos as an example of service registration center
    - [polaris](https://github.com/kitex-contrib/registry-polaris/tree/main/example) Kitex server and client use polaris as an example of service registration center
    - [zookeeper](https://github.com/kitex-contrib/registry-zookeeper) Example of Kitex server and client using zookeeper as service registration center
    - [consul](https://github.com/kitex-contrib/registry-consul/tree/main/example) Example of Kitex server and client using consul as service registration center
    - [servicecomb](https://github.com/kitex-contrib/registry-servicecomb/tree/main/example) Kitex server and client use servicecomb as an example of service registration center
    - [eureka](https://github.com/kitex-contrib/registry-eureka/tree/main/example) Kitex server and client use eureka as an example of service registration center
    - [dns](https://github.com/kitex-contrib/resolver-dns) Example of Kitex server and client using dns for service discovery
    - [resolver_rule_based](https://github.com/kitex-contrib/resolver-rule-based/tree/main/demo) provides a rule-based resolver for Kitex. It allows users to configure rules in service discovery to filter service instances and implement traffic segmentation.
- [timeout](governance/timeout) Example of using timeout control between Kitex server and client
- [limit](governance/limit) Example of using current limit on Kitex server
- [circuit_breaker](governance/circuitbreak) Example of Kitex client using circuit breaker
- [retry](governance/retry) Example of Kitex client using retry
- [load_balance](loadbalancer) Example of Kitex server and client using load balancing

## Observability
- [opentelemetry](opentelemetry) Example of Kitex server and client using OpenTelemetry
- [prometheus](prometheus) Example of Kitex server and client using prometheus
- [tracer](tracer) Example of using tracer on Kitex server and client
- [klog](klog) Example of using klog logs on the Kitex server

## Advanced features
- [frugal](frugal) Kitex example of using frugal on the server and client
- [grpc_proxy](grpcproxy) Example of Kitex server and client using grpc_proxy
- [generic](generic) Kitex example of using generic calls between server and client
- [meta_info](metainfo) Example of Kitex server and client using meta information
- [server_hook](server_hook) Examples of customizing business logic before and after Kitex server startup/exit
- [server_sdk](server_sdk) Example of Kitex server-side SDKization
- [profiler](profiler) Example of Kitex server and client performance analysis using request cost metrics
- [proxyless](proxyless) Example of letting Kitex services run in proxyless mode and be managed uniformly by the service mesh
- [grpc_multi_service](grpc_multi_service) Example of Kitex server and client using grpc multiservice
- [thrift_multi_service](thrift_multi_service) Example of Kitex server and client using thrift (non-streaming) multiservice
- [protobuf_multi_service](protobuf_multi_service) Example of Kitex server and client using protobuf (non-streaming) multiservice
- [goroutine_local_storage](goroutine-local-storage) Example of Kitex server and client using goroutine_local_storage 


## Kitex generated code
- [protobuf](kitex/protobuf) Example of using Kitex and protobuf to generate server code
- [template](kitex/template) Example of using Kitex custom template to generate server code
- [thrift](kitex/thrift) Example of using Kitex and thrift to generate server code
- [protobuf](kitex/protobuf) Example of using Kitex and protobuf to generate server code

## Note

All commands to execute the example should be executed under "kitex-examples".