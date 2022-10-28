# Kitex-Proxyless-Example

## About
This example shows the usage of the xDS enabled Kitex client. 
It includes the complete procedures from Environment Setup, Deployment and Feature demonstration.
Please follow this example step by step to have a try on 

The client-side usage is in `service/src/client.go`.
```
|-- service: the source code of the example services.
|-- yaml: the yaml files of the example services, which will be used for deployment.
```

## Run this example
### 1. Setup Minikube
Set up your local Kubernetes using Minikube, where we can deploy the control plan and our applications. 

ref: https://minikube.sigs.k8s.io/docs/start/ 

### 2. Install Istio
Download and install Istio in Minikube.

ref: https://istio.io/latest/docs/setup/getting-started/#download

#### Disable sidecar Injection
We will deploy our service in the `proxyless` namespace. 
So, we should disable the automatic sidecar injection in this namespace. 
```
kubectl label namespace proxyless istio-injection-

# check if the sidecar injection is disabled
kubectl get namespace -L istio-injection
```

### 3. Deploy Kitex Applications
#### Build the image
Server and client use the same image, which reads the environment variable to determine the role.

```
cd service
sh ./build_image.sh
```

#### Deploy with script
```
# This script execs "kubectl" command to deploy Server and Client. Do not deploy the test-controller.
# May replace the script with code using k8s client to control the whole deployment in the future. 

# execute at the root direction of this project
sh ./deploy.sh
```

#### Deploy manually
* create namespace 
```
kubectl create namespace proxyless
```
* server
```
kubectl apply -f "./yaml/server/kitex_server.yaml" --namespace=proxyless
```

* client
```
kubectl apply -f "./yaml/client/kitex_client.yaml" --namespace=proxyless
```
Since Kitex has not support mTLS, we disable tls in trafficPolicy for now.
```
trafficPolicy:
tls:
  mode: DISABLE
```

* test-controller (Optional)
> The test controller is used to test the Proxyless client.
>> 1. Delete the pods of Server randomly and check if the client can connect to the new pod of Server. 
```
kubectl apply -f "./yaml/testutil/controller.yaml" --namespace=proxyless
```

#### Observe the logs
* check the logs of kitex-client using kubectl logs
```
# get the podname of client
kubectl get pods --namespace=proxyless

# check the logs
kubectl logs ${pod_name} --namespace=proxyless -f
```

### 4. Apply Traffic Routing Policies
Before apply the traffic routing policies, 
the requests will be sent to both server-v1 and server-v2.

#### VirtualService

* traffic split: 90% to server-v1, 10% to server-v2
```
kubectl apply -f ./yaml/server/virtualService_traffic_split.yaml --namespace=proxyless
```

* traffic route based on path matching: 
  * In this case, all traffic will be routed to server-v2
```
kubectl apply -f ./yaml/server/virtualService_match_path.yaml --namespace=proxyless
```

* traffic route based on header matching: 
  * In this case, all traffic will be routed to server-v1
```
kubectl apply -f ./yaml/server/virtualService_match_tag.yaml --namespace=proxyless
```

#### Thrift Proxy
* patch the thrift-proxy with inline route to outbound listener
* The name of the listener should be replaced by `${clusterIP}_${port}` of the server service.

```
kubectl apply -f ./yaml/server/thrift_proxy.yaml --namespace=proxyless
```