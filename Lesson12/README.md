## Lesson 12: Batch Jobs

Train your model
```
$ kubectl apply -f k8s/train.yaml
```

Cleanup:
```
$ kubectl delete jobs.batch train-model
```
