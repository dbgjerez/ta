# ArgoCD 

## Operator
```zsh
❯ curl -sL https://github.com/operator-framework/operator-lifecycle-manager/releases/download/v0.20.0/install.sh | bash -s v0.20.0
❯ k create -f https://operatorhub.io/install/argocd-operator.yaml
```

## ArgoCD server instance

```zsh
❯ k create ns argocd
❯ k apply -f argocd/server.yaml
```

## Bootstrap the cluster

### rootless

```zsh
minikube config set rootless true
```

### Create the namespaces
> **NOTE**: I have to add the namespace to the bootstrap. 

```zsh
❯ k apply -f argocd/ta-prod-namespaces.yaml
```

### Create the application
This application references a repositories, app projects and applications.

```zsh
❯ k apply -f argocd/ta-app-bootstrap.yaml
```
