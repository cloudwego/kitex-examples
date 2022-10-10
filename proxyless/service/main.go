package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"

	"github.com/cloudwego/kitex-proxyless-test/config"
	"github.com/cloudwego/kitex-proxyless-test/service/src"
	"github.com/cloudwego/kitex-proxyless-test/service/src/benchmark"
)

const (
	serviceNameKey = "MY_SERVICE_NAME"
	pprofAddr      = ":6789"
)

func pprof() {
	_ = http.ListenAndServe(pprofAddr, nil)
}

func main() {
	go pprof()
	serviceName, ok := os.LookupEnv(serviceNameKey)
	if !ok {
		panic("Please specify the service name")
	}

	// use the current namespace
	namespace, ok := os.LookupEnv(config.POD_NAMESPACE_KEY)
	if !ok {
		panic("Please specify the namespace")
	}

	var svc src.TestService
	switch serviceName {
	case config.TestClientSvc:
		svc = src.NewProxylessClient(fmt.Sprintf("%s.%s.%s:%s", config.TestServerSvc, namespace, config.Suffix, config.ServerServicePort))
	case config.TestServerSvc:
		svc = src.NewProxylessServer()
	case config.BenchmarkClient:
		svc = benchmark.NewBenchmarkRunner(fmt.Sprintf("%s.%s.%s:%s", config.TestServerSvc, namespace, config.Suffix, config.ServerServicePort))
	}
	fmt.Println("TEST SERVICE START")
	err := svc.Run()
	if err != nil {
		errMsg := fmt.Errorf("test failed with error: %s", err.Error())
		fmt.Println(errMsg)
	}
}
