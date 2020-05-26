#/bin/bash
set -e

port=5000
if [ "$1" = "web" ]; then
  $port=80
fi

jinja2 deployment.yml -D service=$service -D port=$port -D tag=$tag |
  kubectl apply -f -

jinja2 service.yml -D service=$service -D port=$port |
  kubectl apply -f -
