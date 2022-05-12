## Step to deploy in Kubernetes cluster

### 1. Generate [Secret](https://kubernetes.io/docs/tasks/configmap-secret/)
* Using base64 decode !!
```
$echo -n "user" | base64
$echo -n "password" | base64
$echo -n "data" | base64
```

Secret of database in file `db_secret.yml`

### 2. Deploy in Kubernetes cluster
```
$kubectl apply -f k8s/
$kubectl get all
```

Get all resources
```
$kubectl get pod
$kubectl get deployment
$kubectl get rs
$kubectl get svc
```

Scaling deployment
```
$kubectl scale deployment frontend --replicas=5
$kubectl scale deployment backend --replicas=5
```

Access to frontend
* Open url=http://68.183.235.193:32000/ in web browser


### 3. Working with [Ingress](https://kubernetes.io/docs/concepts/services-networking/ingress/)



### Delete all resources
```
$kubectl delete -f k8s/
```
