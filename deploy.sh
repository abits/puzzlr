#!/bin/sh

curl -LO https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl
chmod +x ./kubectl
sudo mv ./kubectl /usr/local/bin/kubectl

echo "Create Guestbook"
IP_ADDR=$(bx cs workers $CLUSTER_NAME | grep normal | awk '{ print $2 }')
if [ -z $IP_ADDR ]; then
  echo "$CLUSTER_NAME not created or workers not ready"
  exit 1
fi

echo -e "Configuring vars"
exp=$(bx cs cluster-config $CLUSTER_NAME | grep export)
if [ $? -ne 0 ]; then
  echo "Cluster $CLUSTER_NAME not created or not ready."
  exit 1
fi
eval "$exp"

echo -e "Downloading k8s-deploy.yaml"
curl --silent "https://raw.githubusercontent.com/abits/puzzlr/master/k8s-deploy.yaml" > k8s-deploy.yaml

#Find the line that has the comment about the load balancer and add the nodeport def after this
#let NU=$(awk '/^  # type: LoadBalancer/{ print NR; exit }' k8s-deploy.yaml)+3
#NU=$NU\i
#sed -i "$NU\ \ type: NodePort" k8s-deploy.yaml #For OSX: brew install gnu-sed; replace sed references with gsed

#echo -e "Deleting previous version of app if it exists"
#kubectl delete --ignore-not-found=true -f k8s-deploy.yaml

echo -e "Creating pods"
kubectl apply -f k8s-deploy.yaml

PORT=$(kubectl get services | grep frontend | sed 's/.*:\([0-9]*\).*/\1/g')

echo ""
echo "View the app at http://$IP_ADDR:$PORT"