#/bin/bash
set -e

port=5000
if [ "$1" = "web" ]; then
  port=80
fi

jinja2 deployment.yml -D service=$1 -D port=$port -D tag=$tag -D email=$email bot_token=$bot_token |
  kubectl apply -f -

jinja2 service.yml -D service=$1 -D port=$port |
  kubectl apply -f -
