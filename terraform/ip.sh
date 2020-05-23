#!/bin/bash
ip=$(kubectl get ingress $service -o=jsonpath='{.status.loadBalancer.ingress[].ip}')
if [ -n "$ip" ]; then
  echo ip = \"$ip\" > ip.auto.tfvars
fi
