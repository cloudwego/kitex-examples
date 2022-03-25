// Copyright 2021 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package main

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/genericclient"
	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/cloudwego/kitex/pkg/klog"
)

func doBizMethod1(cli genericclient.Client) error {
	body := map[string]interface{}{
		"text": "my test",
	}
	data, err := json.Marshal(body)
	if err != nil {
		klog.Fatalf("body marshal failed: %v", err)
	}
	url := "http://example.com/life/client/11?vint64=1&items=item0,item1,itme2"
	req, err := http.NewRequest(http.MethodGet, url, bytes.NewBuffer(data))
	if err != nil {
		klog.Fatalf("new http request failed: %v", err)
	}
	req.Header.Set("token", "1")
	customReq, err := generic.FromHTTPRequest(req)
	if err != nil {
		klog.Fatalf("convert request failed: %v", err)
	}
	resp, err := cli.GenericCall(context.Background(), "", customReq)
	if err != nil {
		klog.Fatalf("generic call failed: %v", err)
	}
	realResp := resp.(*generic.HTTPResponse)
	klog.Infof("method1 response, status code: %v, headers: %v, body: %v\n", realResp.StatusCode, realResp.Header, realResp.Body)
	return nil
}

func doBizMethod2(cli genericclient.Client) error {
	body := map[string]interface{}{
		"text": "my test",
	}
	data, err := json.Marshal(body)
	if err != nil {
		klog.Fatalf("body marshal failed: %v", err)
	}
	url := "http://example.com/life/client/22?vint64=2&items=item0,item1,itme2"
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(data))
	if err != nil {
		klog.Fatalf("new http request failed: %v", err)
	}
	req.Header.Set("token", "2")
	customReq, err := generic.FromHTTPRequest(req)
	if err != nil {
		klog.Fatalf("convert request failed: %v", err)
	}
	resp, err := cli.GenericCall(context.Background(), "", customReq)
	if err != nil {
		klog.Fatalf("generic call failed: %v", err)
	}
	realResp := resp.(*generic.HTTPResponse)
	klog.Infof("method2 response, status code: %v, headers: %v, body: %v\n", realResp.StatusCode, realResp.Header, realResp.Body)
	return nil
}

func doBizMethod3(cli genericclient.Client) error {
	body := map[string]interface{}{
		"text": "my test",
	}
	data, err := json.Marshal(body)
	if err != nil {
		klog.Fatalf("body marshal failed: %v", err)
	}
	url := "http://example.com/life/client/33/other?vint64=3&items=item0,item1,itme2"
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(data))
	if err != nil {
		klog.Fatalf("new http request failed: %v", err)
	}
	req.Header.Set("token", "3")
	customReq, err := generic.FromHTTPRequest(req)
	if err != nil {
		klog.Fatalf("convert request failed: %v", err)
	}
	resp, err := cli.GenericCall(context.Background(), "", customReq)
	if err != nil {
		klog.Fatalf("generic call failed: %v", err)
	}
	realResp := resp.(*generic.HTTPResponse)
	klog.Infof("method3 response, status code: %v, headers: %v, body: %v\n", realResp.StatusCode, realResp.Header, realResp.Body)
	return nil
}

func main() {
	path := "./http.thrift" // depends on current directory
	p, err := generic.NewThriftFileProvider(path)
	if err != nil {
		klog.Fatalf("new thrift file provider failed: %v", err)
	}
	g, err := generic.HTTPThriftGeneric(p)
	if err != nil {
		klog.Fatalf("new http thrift generic failed: %v", err)
	}
	cli, err := genericclient.NewClient("echo", g, client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		klog.Fatalf("new http generic client failed: %v", err)
	}
	for {
		doBizMethod1(cli)
		time.Sleep(time.Second)
		doBizMethod2(cli)
		time.Sleep(time.Second)
		doBizMethod3(cli)
		time.Sleep(time.Second)
	}
}
