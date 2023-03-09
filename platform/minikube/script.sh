PWD=$(pwd)
FILE_ARGOCD_OPERATOR=/components/base/argocd/argocd-operator.yaml
FILE_ARGOCD_SERVER=/components/base/argocd/argocd-server.yaml
FILE_ARGOCD_BOOTSTRAP=/platform/minikube/bootstrap-components.yaml
SLEEP=5

if [[ ! $PWD$FILE_ARGOCD_SERVER ]] ; then
	echo "❗You must execute the script from root git folder"
	exit
else
	echo "👍 $FILE_ARGOCD_SERVER found"
fi

if [[ ! $PWD$FILE_ARGOCD_OPERATOR ]] ; then
	echo "❗You must execute the script from root git folder"
	exit
else
	echo "👍 $FILE_ARGOCD_OPERATOR found"
fi

echo "👍 [All checks ok]"
echo "-------"

minikube start --cpus=6 --memory='20g' --vm-driver=kvm2

# minikube addons enable ingress

# kubectl --namespace ingress-nginx wait \
#     --for=condition=ready pod \
#     --selector=app.kubernetes.io/component=controller \
#     --timeout=120s

curl -sL https://github.com/operator-framework/operator-lifecycle-manager/releases/download/v0.22.0/install.sh | bash -s v0.22.0

kubectl create -f $PWD$FILE_ARGOCD_OPERATOR

while  
	! kubectl -n operators wait \
		--for condition=established \
		--timeout=60s \
		crd/argocds.argoproj.io
do 
	echo "⌛ Waiting for CRD creations"
	sleep $SLEEP 
done

while  
	! kubectl --namespace operators wait \
    	--for=condition=ready pod \
    	--selector=control-plane=controller-manager \
    	--timeout=60s
do 
	echo "⌛ Waiting for ArgoCD controller"
	sleep $SLEEP 
done

kubectl create ns argocd
kubectl apply -f $PWD$FILE_ARGOCD_SERVER