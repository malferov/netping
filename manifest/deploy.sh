#/bin/bash
set -e

jinja2 frontend.yml -D service=$service -D tag=$tag |
  kubectl apply -f -

jinja2 secret.yml -D crt=$crt -D key=$key |
  kubectl apply -f -

kubectl apply -f service.yml
kubectl apply -f ingress.yml
