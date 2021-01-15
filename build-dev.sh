#!/bin/bash
docker-compose -f docker-compose.dev.yml up -d --build
docker tag tfapi_go qweliant/tfapi_go:latest
docker push qweliant/tfapi_go
kubectl delete service,deployment tfapi
kubectl apply -f kubes.yaml