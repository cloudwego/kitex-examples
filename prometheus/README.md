# Prometheus monitoring for Kitex

## Usage Example

### Server

See [server](./server)

### Client

See [client](./client)

## HOW-TO-RUN

1. install docker and start docker
2. change `{{ INET_IP }}` to local ip in prometheus.yml
3. run Prometheus and Grafana  
   `docker-compose up`
4. run Kitex server   
   `go run server/main.go`
5. run Kitex client  
   `go run client/main.go`
6. visit `http://localhost:3000`, the account password is `admin` by default
7. configure Prometheus data sources
    1. `Configuration`
    2. `Data Source`
    3. `Add data source`
    4. Select `Prometheus` and fill the URL with `http://prometheus:9090`
    5. click `Save & Test` after configuration to test if it works
8. add dashboard `Create` -> `dashboard`, add monitoring metrics such as throughput and pct99 according to your needs,
   for example:

    - server throughput of succeed requests

   `sum(rate(kitex_server_throughput{statusCode="200"}[1m])) by (method)`

    - server latency pct99 of succeed requests

   `histogram_quantile(0.9,sum(rate(kitex_server_latency_us_bucket{status="succeed"}[1m]))by(le)) / 1000`

For more information about kitex monitoring, please
click [monitoring](https://www.cloudwego.io/docs/kitex/tutorials/service-governance/monitoring/)