#/bin/bash
set -e

sudo pip install jinja2-cli
whereis jinja2

jinja2 frontend.yml \
  -D registry=$registry \
  -D actor=$actor \
  -D app=$app \
  -D service=$service \
  -D ver=$ver |
  kubectl apply -f -

jinja2 secret.yml -D crt=$crt -D key=$key | kubectl apply -f -
kubectl apply -f service.yml
kubectl apply -f ingress.yml
