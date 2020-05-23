#/bin/bash
set -e

jinja2 frontend.yml -D service=$service -D tag=$tag |
  kubectl apply -f -

jinja2 secret.yml -D crt=$crt -D key=$key |
  kubectl apply -f -

jinja2 service.yml -D service=$service |
  kubectl apply -f -

jinja2 ingress.yml -D service=$service -D app=$app -D domain=$domain |
  kubectl apply -f -
