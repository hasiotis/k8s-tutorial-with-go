## Lesson 04: Enable logging

Create and push a docker images of our application:
```
CGO_ENABLED=0 go build tutorial.go
docker build -t ${DOCKER_REGISTRY}/tutorial:v0.0.4 .
docker push ${DOCKER_REGISTRY}/tutorial:v0.0.4
```

Deploy to our kubernetes cluster:
```
kubectl apply -f k8s/deployment.yaml
kubectl rollout status deployment tutorial-deployment
http EXTERNAL_IP
```

Now check the logs:
```
$ kubectl get pod
$ kubectl logs -f <POD_NAME>
$ stern tutorial
```
