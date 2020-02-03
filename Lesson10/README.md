## Lesson 10: Managing state

So far we had completely stateless apps. Let's try something else:
```
CGO_ENABLED=0 go build tutorial.go
docker build -t ${DOCKER_REGISTRY}/tutorial:v0.0.10 .
docker push ${DOCKER_REGISTRY}/tutorial:v0.0.10
kubectl apply -f k8s/deployment.yaml
```

Get URL: http://_EXTERNALIP_/get
Set URL: http://_EXTERNALIP_/set/myvalue

Here is how we can deal with this problem:
```
kubectl apply -f k8s/pv.yaml
kubectl apply -f k8s/deployment-volume.yaml
```

Check that you have a disk
```
$ kubectl exec -it <PODNAME> -- sh
$ df -h /data
$ ls -l /data
```
