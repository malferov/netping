#/bin/bash
set -e

jinja2 secret.yml -D crt=$crt -D key=$key |
  kubectl apply -f -

jinja2 ingress.yml -D app=$app -D domain=$domain |
  kubectl apply -f -
