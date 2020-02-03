## Lesson 13: Shared filesystems

Use a ReadWriteMany volume:
```
$ kubectl apply -f k8s/pv.yaml
$ kubectl apply -f k8s/deployment.yaml
```

Train your model
```
$ kubectl apply -f k8s/train.yaml
```
