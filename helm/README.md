helm install --dry-run --debug ./whitemos

helm install --name whitemos ./whitemos

helm delete --purge whitemos

helm lint ./whitemos

helm package ./whitemos