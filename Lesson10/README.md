## Lesson 10: Tips and tricks

Create and push a docker images of our application:
```
CGO_ENABLED=0 go build tutorial.go
docker build -t ${DOCKER_REGISTRY}/tutorial:v0.0.10 .
docker push ${DOCKER_REGISTRY}/tutorial:v0.0.10
```

Redeploy to our kubernetes cluster:
```
kubectl apply -f k8s/deployment.yaml
kubectl rollout status deployment tutorial-deployment
http EXTERNAL_IP
```

Redeploy to our kubernetes cluster:
```
kubectl apply -f k8s/deployment-volume.yaml
kubectl rollout status deployment tutorial-deployment
http EXTERNAL_IP
```
