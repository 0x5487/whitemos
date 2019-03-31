helm install --dry-run --debug ./whitemos

#### install or upgrade depending on if the release exists
helm upgrade --install whotemos ./whitemos --wait 
helm upgrade --install whotemos jasonsoft/whitemos --wait

helm delete --purge whitemos

helm lint ./whitemos

helm package ./whitemos