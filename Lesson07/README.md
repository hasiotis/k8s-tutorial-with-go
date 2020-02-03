## Lesson 07: Secrets

Create and push a docker images of our application:
```
CGO_ENABLED=0 go build tutorial.go
docker build -t ${DOCKER_REGISTRY}/tutorial:v0.0.7 .
docker push ${DOCKER_REGISTRY}/tutorial:v0.0.7
```

Redeploy to our kubernetes cluster:
```
kubectl apply -f k8s/secret.yaml
kubectl apply -f k8s/deployment.yaml
kubectl rollout status deployment tutorial-deployment
http EXTERNAL_IP
```
