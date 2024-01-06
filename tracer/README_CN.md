# tracer 示例
[English](./README.md) | 中文
## opentracing
### 如何运行
1. 安装 docker
2. 运行 jaeger all-in-one  
`sh jaeger_run.sh`
3. 运行 Kitex 服务端  
`sh server/opentracing/run.sh`
4. 运行 Kitex 客户端
打开另一个终端并运行  
`sh client/opentracing/run.sh`
### 监控
然后，你前往 http://localhost:16686 来访问 Jaeger UI。（详情可访问 [Monitor Jaeger](https://www.jaegertracing.io/docs/1.24/monitoring/) ）