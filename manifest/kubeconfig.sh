#!/bin/bash
set -e
kubectl config set-cluster kube --server=https://$K8S_SERVER --insecure-skip-tls-verify=true
kubectl config set-credentials CI --username=$K8S_USER --password=$K8S_PASS
kubectl config set-context kube --cluster=kube --user=CI --namespace=$app
kubectl config use-context kube
kubectl version
