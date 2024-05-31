Istio Itegration Demo
---

Introduce the journey to integrate Istio.  

- Client: switch the client protocol to GRPC.

    ```
	client, err := api.NewClient("hello",
		client.WithHostPorts(serviceName+":8888"),
		// should use grpc protocol, thrift is not well compatible in istio.
		client.WithTransportProtocol(transport.GRPC),
	)
    ```
- Server: specify the protocol kind in the Kubernetes `Service` [definition](https://istio.io/latest/docs/ops/configuration/traffic-management/protocol-selection/#explicit-protocol-selection) .

```
apiVersion: v1
kind: Service
metadata:
  name: hello
spec:
  selector:
    app: hello
  ports:
  - port: 8888
    # istio detect the protocol from the name, this can be configured in two ways: 
    # 1. By the name of the port name: <protocol>[-<suffix>].
    # 2. In Kubernetes 1.18+, by the appProtocol field: appProtocol: <protocol>.
    name: grpc
    protocol: TCP
  type: ClusterIP
```
