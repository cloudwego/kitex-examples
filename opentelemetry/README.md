# Opentelemetry Example

## HOW-TO-RUN
1. install docker
2. run jaeger + vm + grafana
   `docker-compose up -d`
3. run Kitex server
    `go run server/main.go`
4. run Kitex client
   `go run client/main.go`

## MONITORING

## View Trace
You can then navigate to http://localhost:16686 to access the Jaeger UI. (You can visit Monitor Jaeger for details)
![img.png](static/jaeger.png)

## View Metrics
You can then navigate to http://localhost:3000 to access the Grafana UI. (You can visit Monitor Grafana for metrics)

### add datasource

URL: `http://host.docker.internal:8428/`

![img_1.png](static/grafana.png)
### add a dashboard and a panel
![img.png](static/panel.png)
### support metrics 
- RPC Metrics
- Runtime Metrics


