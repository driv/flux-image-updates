# FluxCD bootstrap
```
kubectl apply -k clusters/my-cluster/flux-system
```
## Automated cluster bootstrap

https://fluxcd.io/flux/installation/bootstrap/github/

## Cluster manual bootstrap

```
flux create secret git flux-system \
    --url=ssh://git@github.com/driv/flux-image-updates \
    --export > gotk-secret.yaml
```


Apply the secret to the cluster and use the `identity.pub` public key from the secret as a deployment key with write access for your repo. 
