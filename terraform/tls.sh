#!/bin/bash
set -e
terraform apply -auto-approve
export crt=$(terraform output crt)
export key=$(terraform output key)
