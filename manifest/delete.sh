#/bin/bash
set -e

jinja2 ingress.yml -D app=$app -D domain=$domain |
  kubectl delete -f -

jinja2 secret.yml | kubectl delete -f -

for s in web ping; do
  jinja2 deployment.yml -D service=$s | kubectl delete -f -
  jinja2 service.yml -D service=$s | kubectl delete -f -
done
