#/bin/bash
#pip install jinja2-cli
jinja2 frontend.yml \
  -D registry=$registry \
  -D actor=$actor \
  -D app=$app \
  -D service=$service \
  -D ver=$ver |
  kubectl apply -f -
#kubectl apply -f service.yml
#kubectl apply -f secret.yml
#kubectl apply -f ingress.yml
