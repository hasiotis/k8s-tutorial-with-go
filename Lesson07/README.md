## Lesson 07: Secrets

Create and push a docker images of our application:
```
CGO_ENABLED=0 go build tutorial.go
docker build -t ${DOCKER_REGISTRY}/tutorial:v0.0.7 .
docker push ${DOCKER_REGISTRY}/tutorial:v0.0.7
```

Encode your secret using base64:
```
$ echo -n 'dupersecret' | base64
```

Inform kubernetes about it:
```
$ kubectl apply -f k8s/secret.yaml
```

Redeploy to our kubernetes cluster:
```
kubectl apply -f k8s/deployment.yaml
kubectl rollout status deployment tutorial-deployment
http EXTERNAL_IP
```

