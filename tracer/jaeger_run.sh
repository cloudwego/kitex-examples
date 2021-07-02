#!/bin/sh
docker run -d -p 5775:5775/udp -p 16686:16686 jaegertracing/all-in-one:latest