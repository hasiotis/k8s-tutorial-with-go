## Lesson 11: Cron Jobs

Let's enable our reset functionality
```
CGO_ENABLED=0 go build tutorial.go
docker build -t ${DOCKER_REGISTRY}/tutorial:v0.0.11 .
docker push ${DOCKER_REGISTRY}/tutorial:v0.0.11
```

Enable the cronjob
```
kubectl apply -f k8s/cronjob.yaml
```

Check what happened:
```
kubectl get cronjobs.batch reset-cache
kubectl get pod
```
