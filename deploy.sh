docker build -t cx-config-center .
docker tag cx-config-center registry.local.com/cx-config-center
docker push registry.local.com/cx-config-center
kubectl delete deployment cx-config-center-deployment -n cx-rpc-base
kubectl apply -f cx-config-center-deployment.yaml
docker rmi cx-config-center
docker rmi registry.local.com/cx-config-center
