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
* Open url=http://<node_ip>:32000/ in web browser


### 3. Working with [Ingress](https://kubernetes.io/docs/concepts/services-networking/ingress/)

Deploy [NGINX ingress controller](https://kubernetes.github.io/ingress-nginx/deploy/)
```
$kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.2.0/deploy/static/provider/cloud/deploy.yaml
```

Create ingress
```
$kubectl apply -f my_ingress.yml
$kubectl get ingress
$kubectl describe ingress

$kubectl get service ingress-nginx-controller --namespace=ingress-nginx

NAME                       TYPE           CLUSTER-IP       EXTERNAL-IP   PORT(S)                      AGE
ingress-nginx-controller   LoadBalancer   10.105.160.129   <pending>     80:32230/TCP,443:32112/TCP   41

$curl -kL http://178.128.122.79:32230/api/account/1
```

### Delete all resources
```
$kubectl delete -f k8s/
```
