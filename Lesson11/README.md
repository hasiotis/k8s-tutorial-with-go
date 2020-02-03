## Lesson 11: Cron Jobs

Let's enable our reset functionality
```
CGO_ENABLED=0 go build tutorial.go
docker build -t ${DOCKER_REGISTRY}/tutorial:v0.0.11 .
docker push ${DOCKER_REGISTRY}/tutorial:v0.0.11
kubectl apply -f k8s/deployment.yaml
```

Enable the cronjob
```
kubectl apply -f k8s/cronjob.yaml
```

Check what happened:
```
kubectl get cronjobs.batch tutorial-reset-cache
kubectl get pod
```
