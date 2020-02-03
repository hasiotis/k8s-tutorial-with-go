## Lesson 02: Start some workload

First let's isolate ourselfs:
```
kubectl create namespace tutorial
kubens tutorial
```

Now start an nginx server:
```
kubectl create deployment nginx --image=nginx
```

What happened?
```
kubectl get pod
kubectl get pod -o wide
```

What if we need two of them?
```
kubectl scale deployment --replicas=2 nginx
kubectl get pod -o wide
```

Now you have to create a service:
```
kubectl expose deployment nginx --port=80 --target-port=80 --type=LoadBalancer
kubectl get svc nginx
kubectl describe svc nginx
http EXTERNAL_IP
```

Now both pods serve traffic. Prove it with:
```
stern nginx
```

Now let's cleanup
```
kubectl delete deployment nginx
kubectl delete svc nginx
```

Move to lesson 03.
