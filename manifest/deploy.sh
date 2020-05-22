#/bin/bash
kubectl apply -f frontend.yml
kubectl apply -f service.yml
kubectl apply -f secret.yml
kubectl apply -f ingress.yml
