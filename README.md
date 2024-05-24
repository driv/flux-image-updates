Since FluxCD pushes updates to the repo you need to fork this repo and mofify a few things before getting started.

```
# Fix users
export GIT_USER=<git_username>
export DOCKER_USER=<dockerhub_username>

grep -l -r '/driv/' | xargs -L1 sed -i "s,/driv/,/$GIT_USER/,g"
grep -l -r ' driv/' | xargs -L1 sed -i "s, driv/, $DOCKER_USER/,g"
```


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
kubectl apply -f gotk-secret.yaml
```

Use the `identity.pub` public key from the secret as a deployment key with write access for your repo. Do not push the secret to git.

# Kpack registry secret
Create a token with read/write access. https://hub.docker.com/settings/security

Replace `<username>` with your DockerHub username and <token> with your generated token, and apply the secret. Do not push the secret to git.
```
apiVersion: v1
kind: Secret
metadata:
  name: registry-credentials
  namespace: kpack-builders
  annotations:
    kpack.io/docker: https://index.docker.io/v1/
type: kubernetes.io/basic-auth
stringData:
  username: <username>
  password: <token>
```
