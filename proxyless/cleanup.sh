#!/bin/bash

kubectl delete -f "./yaml/server/kitex_server.yaml" --namespace=proxyless
kubectl delete -f "./yaml/client/kitex_client.yaml" --namespace=proxyless
kubectl delete -f "./yaml/testutil/controller.yaml" --namespace=proxyless
kubectl delete namespace proxyless