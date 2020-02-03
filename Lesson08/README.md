## Lesson 08: Liveness and Readiness

Create and push a docker images of our application:
```
CGO_ENABLED=0 go build tutorial.go
docker build -t ${DOCKER_REGISTRY}/tutorial:v0.0.8 .
docker push ${DOCKER_REGISTRY}/tutorial:v0.0.8
```

Redeploy to our kubernetes cluster:
```
kubectl apply -f k8s/deployment.yaml
kubectl rollout status deployment tutorial-deployment
http EXTERNAL_IP
```
