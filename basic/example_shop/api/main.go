package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"example_shop/kitex_gen/example/shop/item"
	"example_shop/kitex_gen/example/shop/item/itemservice"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func main() {
	http.HandleFunc("/api/item", Handler)
	http.ListenAndServe("localhost:8080", nil)
}

func Handler(rw http.ResponseWriter, r *http.Request) {
	resolver, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err)
	}

	c, err := itemservice.NewClient("example.shop.item", client.WithResolver(resolver), client.WithConnectTimeout(time.Hour), client.WithRPCTimeout(time.Hour))
	if err != nil {
		log.Fatal(err)
	}
	req := item.NewGetItemReq()
	req.Id = 1024
	resp, err := c.GetItem(context.Background(), req, callopt.WithRPCTimeout(3*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	rw.Write([]byte(resp.String()))
}
