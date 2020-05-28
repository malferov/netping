#!/bin/bash
web=$(kubectl get ingress frontend -o=jsonpath='{.status.loadBalancer.ingress[].ip}')
api=$(kubectl get ingress backend  -o=jsonpath='{.status.loadBalancer.ingress[].ip}')

if [ -n "$web" ]; then
  echo web = \"$web\" > web.auto.tfvars
fi

if [ -n "$api" ]; then
  echo api = \"$api\" > api.auto.tfvars
fi
