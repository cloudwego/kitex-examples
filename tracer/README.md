# tracer example
## opentracing
### HOW-TO-RUN
1. install docker
2. run jaeger all-in-one   
`sh jaeger_run.sh`
3. run KiteX server   
`sh server/opentracing/run.sh`
4. run KiteX client   
open another terminal and run `sh client/opentracing/run.sh`