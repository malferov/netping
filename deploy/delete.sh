#/bin/bash
set -e

jinja2 ingress.yaml -D app=$app -D domain=$domain |
  kubectl delete -f -

jinja2 secret.yaml | kubectl delete -f -

for s in web ping; do
  jinja2 deployment.yaml -D service=$s | kubectl delete -f -
  jinja2 service.yaml -D service=$s | kubectl delete -f -
done
