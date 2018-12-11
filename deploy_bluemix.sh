#!/bin/bash
echo "Install Bluemix CLI"

echo "Download Bluemix CLI"
wget --quiet --output-document=/tmp/Bluemix_CLI_amd64.tar.gz  http://public.dhe.ibm.com/cloud/bluemix/cli/bluemix-cli/latest/Bluemix_CLI_amd64.tar.gz
tar -xf /tmp/Bluemix_CLI_amd64.tar.gz --directory=/tmp

# Create bx alias
echo "#!/bin/sh" >/tmp/Bluemix_CLI/bin/bx
echo "/tmp/Bluemix_CLI/bin/bluemix \"\$@\" " >>/tmp/Bluemix_CLI/bin/bx
chmod +x /tmp/Bluemix_CLI/bin/*

export PATH="/tmp/Bluemix_CLI/bin:$PATH"

# Configure bluemix
bx config --check-version false

# Install Armada CS plugin
echo "Install the Bluemix container-service plugin"
bx plugin install container-service -r Bluemix

echo "Setup Bluemix API endpoint"
bx api https://api.eu-de.bluemix.net

echo "Install kubectl"
wget --quiet --output-document=/tmp/Bluemix_CLI/bin/kubectl  https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl
chmod +x /tmp/Bluemix_CLI/bin/kubectl

if [ -n "$DEBUG" ]; then
  bx --version
  bx plugin list
fi

if [ $? -ne 0 ]; then
  echo "Failed to install Bluemix Container Service CLI prerequisites"
  exit 1
fi

#####################

echo "Login to Bluemix"

if [ -z $CF_ORG ]; then
  CF_ORG="$BLUEMIX_ORG"
fi
if [ -z $CF_SPACE ]; then
  CF_SPACE="$BLUEMIX_SPACE"
fi

if [ -z "$BLUEMIX_USER" ] || [ -z "$BLUEMIX_PASSWORD" ] || [ -z "$BLUEMIX_ACCOUNT" ]; then
  echo "Define all required environment variables and rerun the stage."
  exit 1
fi
echo "Deploy pods"

echo "bx login -a $CF_TARGET_URL"
#bx login -a "$CF_TARGET_URL" -u "$BLUEMIX_USER" -p "$BLUEMIX_PASSWORD" -c "$BLUEMIX_ACCOUNT" -o "$CF_ORG" -s "$CF_SPACE"
bx login -a "$CF_TARGET_URL" -u "$BLUEMIX_USER" -p "$BLUEMIX_PASSWORD" -c "$BLUEMIX_ACCOUNT"
if [ $? -ne 0 ]; then
  echo "Failed to authenticate to Bluemix"
  exit 1
fi

# set target resourcegroup
echo "bx target -g $BLUEMIX_RESOURCEGROUP"
bx target -g "$BLUEMIX_RESOURCEGROUP"
if [ $? -ne 0 ]; then
  echo "Failed to set Bluemix target resourcegroup"
  exit 1
fi

# Init container clusters
echo "bx cs init"
bx cs init
if [ $? -ne 0 ]; then
  echo "Failed to initialize to Bluemix Container Service"
  exit 1
fi

if [ $? -ne 0 ]; then
  echo "Failed to authenticate to Bluemix Container Service"
  exit 1
fi

#####################

echo "Deploy pods"

#!/bin/sh

curl -LO https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl
chmod +x ./kubectl
sudo mv ./kubectl /usr/local/bin/kubectl

echo "Create App"
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

PORT=$(kubectl get services | grep puzzlr-service | sed 's/.*:\([0-9]*\).*/\1/g')

echo ""
echo "View the app at http://$IP_ADDR:$PORT"

if [ $? -ne 0 ]; then
  echo "Failed to Deploy pods to Bluemix Container Service"
  exit 1
fi
