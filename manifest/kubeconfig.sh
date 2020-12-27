#!/bin/bash
set -e
kubectl config set-cluster kube --server=https://$K8S_SERVER:6443
kubectl config set-cluster kube --embed-certs --certificate-authority <(echo $K8S_CA)
kubectl config set-context ci@kube --cluster=kube --user=ci --namespace=$app
kubectl config use-context ci@kube
kubectl config set-credentials ci --embed-certs --client-certificate <(echo $K8S_CRT)
kubectl config set-credentials ci --embed-certs --client-key <(echo $K8S_KEY)
