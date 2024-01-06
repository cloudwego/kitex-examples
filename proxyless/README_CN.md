# Kitex-Proxyless 示例
[English](./README.md) | 中文
## 关于
此示例显示启用 xDS 的 Kitex 客户端的用法。
它包括从环境设置、部署到功能演示的完整过程。
请按照此示例逐步尝试

客户端的用法在 `service/src/client.go` 中。
````
|-- service：示例服务的源代码。
|-- yaml：示例服务的yaml文件，将用于部署。
````

## 运行这个例子
### 1. 设置 Minikube
使用 Minikube 设置本地 Kubernetes，我们可以在其中部署控制计划和应用程序。

参考：https://minikube.sigs.k8s.io/docs/start/

### 2.安装Istio
在 Minikube 中下载并安装 Istio。

参考：https://istio.io/latest/docs/setup/getting-started/#download

#### 禁用 sidecar 注入
我们将在 `proxyless` 命名空间中部署我们的服务。
因此，我们应该在此命名空间中禁用自动 sidecar 注入。
````
kubectl label namespace proxyless istio-injection-

# 检查 sidecar 注入是否被禁用
kubectl get namespace -L istio-injection
````

### 3. 部署 Kitex 应用程序
#### 构建镜像
服务器和客户端使用相同的映像，读取环境变量来确定角色。

```
cd service
sh ./build_image.sh
```

#### 使用脚本部署
````
# 该脚本执行 “kubectl” 命令来部署服务器和客户端。 不要部署测试控制器。
# 将来可能会将脚本替换为使用 k8s 客户端的代码来控制整个部署。

# 在本项目的根方向执行
sh ./deploy.sh
````

#### 手动部署
* 创建命名空间
````
kubectl create namespace proxyless
````
* 服务端
````
kubectl apply -f "./yaml/server/kitex_server.yaml" --namespace=proxyless
````

* 客户端
````
kubectl apply -f "./yaml/client/kitex_client.yaml" --namespace=proxyless
````
由于 Kitex 不支持 mTLS，我们暂时在 trafficPolicy 中禁用 tls。
```
trafficPolicy:
tls:
  mode: DISABLE
```

* 测试控制器（可选）
> 测试控制器用于测试Proxyless客户端。
>> 1. 随机删除Server的Pod，检查客户端是否可以连接到Server的新Pod。
````
kubectl apply -f "./yaml/testutil/controller.yaml" --namespace=proxyless
````

#### 观察日志
* 使用 `kubectl logs` 检查 kitex-client 的日志
````
# 获取客户端的 podname
kubectl get pods --namespace=proxyless

# 检查日志
kubectl logs ${pod_name} --namespace=proxyless -f
````

### 4. 应用流量路由策略
在应用流量路由策略之前，
请求将发送到 server-v1 和 server-v2。

#### 虚拟服务

* 流量分配：90% 分配给 server-v1，10% 分配给 server-v2
````
kubectl apply -f ./yaml/server/virtualService_traffic_split.yaml --namespace=proxyless
````

* 基于路径匹配的流量路由：
    * 在这种情况下，所有流量将被路由到 server-v2
````
kubectl apply -f ./yaml/server/virtualService_match_path.yaml --namespace=proxyless
````

* 基于标头匹配的流量路由：
    * 在这种情况下，所有流量将被路由到 server-v1
````
kubectl apply -f ./yaml/server/virtualService_match_tag.yaml --namespace=proxyless
````

#### Thrift 代理
* 使用到出站监听器的内联路由修补 thrift-proxy
* 监听器的名称应该替换为服务端服务的 `${clusterIP}_${port}`。

````
kubectl apply -f ./yaml/server/thrift_proxy.yaml --namespace=proxyless
````