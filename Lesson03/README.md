## Lesson 03: Let's get serious

Make sure we authenticate to docker registry:
```
gcloud auth configure-docker
```

Create and push a docker images of our application:
```
export DOCKER_REGISTRY="gcr.io/${PROJECT_ID}"
CGO_ENABLED=0 go build tutorial.go
docker build -t ${DOCKER_REGISTRY}/tutorial:v0.0.3 .
docker push ${DOCKER_REGISTRY}/tutorial:v0.0.3
gcloud container images list --repository=${DOCKER_REGISTRY}
```

Deploy to our kubernetes cluster:
```
envsubst < k8s/deployment.yaml | kubectl apply -f -
envsubst < k8s/service.yaml | kubectl apply -f -
kubectl get svc tutorial-service
http EXTERNAL_IP
```


Finally a trick to fix the docker registry:
```
cd ..
find . -name "*.yaml" -exec sed -i s/\${DOCKER_REGISTRY}/gcr.io\\/${PROJECT_ID}/ {} \;
git diff
```
