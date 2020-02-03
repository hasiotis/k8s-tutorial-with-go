## Lesson 06: Service discovery

Create and push a docker images of our application:
```
CGO_ENABLED=0 go build tutorial.go
docker build -t ${DOCKER_REGISTRY}/tutorial:v0.0.6 .
docker push ${DOCKER_REGISTRY}/tutorial:v0.0.6
```

Redeploy to our kubernetes cluster:
```
kubectl apply -f k8s/configmap.yaml
kubectl apply -f k8s/deployment.yaml
kubectl rollout status deployment tutorial-deployment
http EXTERNAL_IP
```
