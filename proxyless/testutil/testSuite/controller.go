/*
 * Copyright 2022 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package testSuite

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/rest"

	"github.com/cloudwego/kitex-examples/proxyless/config"
)

type TestController struct {
	client *kubernetes.Clientset
}

func NewController() *TestController {
	// creates the in-cluster config
	cfg, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientSet, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		panic(err.Error())
	}
	return &TestController{
		client: clientSet,
	}
}

func listPods(prefix string, pi corev1.PodInterface) {
	pods, err := pi.List(context.Background(), metav1.ListOptions{})
	if err != nil {
		klog.Error("ListPod failed: %s", err.Error())
		return
	}
	for _, p := range pods.Items {
		if strings.HasPrefix(p.Name, prefix) {
			fmt.Printf("Pod: %s, status: %s\n", p.Name, p.Status.Phase)
		}
	}
	fmt.Println()
}

func (c *TestController) deletePod(namespace, service string) {
	fmt.Printf("Randomly Delete [%s]'s Pods\n", service)

	podInterface := c.client.CoreV1().Pods(namespace)
	pods, err := podInterface.List(context.Background(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	// select the pods of the specified service
	ps := make([]string, 0)
	for _, p := range pods.Items {
		if strings.HasPrefix(p.Name, service) {
			ps = append(ps, p.Name)
		}
	}

	fmt.Printf("----------Pods Before Delete----------\n")
	listPods(service, podInterface)

	// randomly delete one pod
	idx := rand.Intn(len(ps))
	err = podInterface.Delete(context.Background(), ps[idx], metav1.DeleteOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("DELETE POD: %s\n\n", ps[idx])
	time.Sleep(time.Second)

	fmt.Printf("----------Pods After Delete----------\n")
	listPods(service, podInterface)
}

func (c *TestController) Run() {
	// use the current namespace
	namespace, ok := os.LookupEnv(config.POD_NAMESPACE_KEY)
	if !ok {
		panic("Please specify the namespace")
	}

	for {
		time.Sleep(30 * time.Second)
		// TODO: test controller can be deployed in the namespace different from the services
		c.deletePod(namespace, config.TestServerSvc)
	}
}
