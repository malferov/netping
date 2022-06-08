#/bin/bash
set -e

cd ../infra
jinja2 ../deploy/secret.yaml \
  -D crt=$(terraform output crt) \
  -D key=$(terraform output key) \
  > ../deploy/secret.rendered

cd ../deploy
kubectl apply -f secret.rendered

jinja2 ingress.yaml \
  -D app=$app \
  -D domain=$domain \
  | kubectl apply -f -
