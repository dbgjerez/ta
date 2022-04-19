minikube start --cpus=4 --memory='8g'
minikube addons enable ingress

kubectl --namespace ingress-nginx wait \
    --for=condition=ready pod \
    --selector=app.kubernetes.io/component=controller \
    --timeout=120s

INGRESS_HOST=$(minikube ip)
HOST=argocd.$INGRESS_HOST.nip.io
NS=argocd

helm repo add argo \
    https://argoproj.github.io/argo-helm

helm repo update

helm upgrade --install \
    argocd argo/argo-cd \
    --namespace argocd \
    --create-namespace \
    --version 4.5.3 \
    --set server.ingress.hosts="{$HOST}" \
    --values argocd-values.yaml \
    --wait

## user=admin
## password
# kubectl -n argocd get secret argocd-initial-admin-secret -o json | jq -r ".data.password" | base64 -d

printf "URL: $HOST\n"

