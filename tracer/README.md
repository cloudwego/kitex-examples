# tracer example
## opentracing
### HOW-TO-RUN
1. install docker
2. run jaeger all-in-one   
`sh jaeger_run.sh`
3. run Kitex server   
`sh server/opentracing/run.sh`
4. run Kitex client   
open another terminal and run `sh client/opentracing/run.sh`
### MONITORING
You can then navigate to http://localhost:16686 to access the Jaeger UI. (You can visit [Monitor Jaeger](https://www.jaegertracing.io/docs/1.24/monitoring/) for details)