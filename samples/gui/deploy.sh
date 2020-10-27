#!/bin/bash
# Copyright 2020 IBM Corp.
# SPDX-License-Identifier: Apache-2.0

export VAULT_TOKEN=$(kubectl get secrets vault-unseal-keys -n m4d-system -o jsonpath={.data.vault-root} | base64 --decode)
echo -n $VAULT_TOKEN > ./token.txt
echo $VAULT_TOKEN
kubectl apply -f ../../manager/config/prod/deployment_configmap.yaml
kubectl delete secret vault-unseal-keys -n default
kubectl create secret generic vault-unseal-keys --from-file=vault-root=./token.txt -n default || true
kubectl delete service datauserserver || true
kubectl delete service datauserclient || true
kubectl delete deployment gui -n default
kubectl apply -f Deployment.yaml -n default
rm ./token.txt


