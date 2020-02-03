## Lesson 01: Create your cluster

First let's create the project that will host our cluster:
```
export PROJECT_ID="k8s-tutorial-${USER}"
gcloud projects create ${PROJECT_ID} --name "Kubernetes Tutorial"
gcloud config set project ${PROJECT_ID}
```

Enable billing:
```
gcloud alpha billing accounts list
gcloud alpha billing projects link ${PROJECT_ID} --billing-account 0X0X0X-0X0X0X-0X0X0X
```

And finally create our cluster:
```
gcloud container clusters create k8s-tutorial
```

Install some extra tooling:
```
mkdir -p ~/tmp
cd ~/tmp

# kubectl
curl -LO https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl
chmod +x ./kubectl
sudo mv ./kubectl /usr/local/bin/kubectl

# kubectx
curl -LO https://raw.githubusercontent.com/ahmetb/kubectx/master/kubectx
chmod +x ./kubectx
sudo mv ./kubectx /usr/local/bin/kubectx

# kubens
curl -LO https://raw.githubusercontent.com/ahmetb/kubectx/master/kubens
chmod +x ./kubens
sudo mv ./kubens /usr/local/bin/kubens

# stern
export URL=`http https://api.github.com/repos/wercker/stern/releases/latest | jq -r '.assets[] | select(.name | contains("linux")).browser_download_url'`
wget -q $URL
mv stern_linux_amd64 ~/.local/bin/stern
chmod +x ~/.local/bin/stern
```

Congrats. Move to lesson 02.
