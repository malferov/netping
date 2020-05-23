#!/bin/bash
set -e
terraform apply -target=acme_certificate.cert -auto-approve
export crt=$(terraform output -json crt | jq '.value')
export key=$(terraform output -json key | jq '.value')
