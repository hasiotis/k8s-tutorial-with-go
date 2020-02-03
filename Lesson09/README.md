## Lesson 09: Tips and tricks

You already now about logs:
```
$ kubectl logs <PODNAME>
$ stern -h
```

Apply, Get, describe. You can run get and describe on any kubernetes object.
But you can also try edit!
```
$ kubectl edit deployment tutorial-deployment
```

Try to point to a non existent docker image, and then:
```
$ kubectl describe pod <PODNAME>
```

Make a new docker image, this time starting from alpine:
```
CGO_ENABLED=0 go build tutorial.go
docker build -t ${DOCKER_REGISTRY}/tutorial:v0.0.9 .
docker push ${DOCKER_REGISTRY}/tutorial:v0.0.9
```

Redeploy to our kubernetes cluster:
```
kubectl apply -f k8s/deployment.yaml
kubectl rollout status deployment tutorial-deployment
http EXTERNAL_IP
```

Now we can get an interactive shell into the running container:
```
$ kubectl exec -it <PODNAME> -- sh
```
