# ArgoCD 

## Operator
```zsh
curl -sL https://github.com/operator-framework/operator-lifecycle-manager/releases/download/v0.20.0/install.sh | bash -s v0.20.0
kubectl create -f https://operatorhub.io/install/argocd-operator.yaml
```

## ArgoCD server instance

```zsh
❯ k create ns argocd
❯ k apply -f server.yaml
```

## Bootstrap the cluster