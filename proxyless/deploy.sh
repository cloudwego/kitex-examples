#!/bin/bash

kubectl create namespace proxyless
kubectl apply -f "./yaml/server/kitex_server.yaml" --namespace=proxyless
kubectl apply -f "./yaml/client/kitex_client.yaml" --namespace=proxyless
kubectl apply -f "./yaml/testutil/controller.yaml" --namespace=proxyless