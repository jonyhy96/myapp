# Kubectl 命令测试


### 1.测试命令 kubectl get all

```
[root@node1 ~]# kubectl get all

```
### 1.运行结果
```
NAME                                       READY     STATUS    RESTARTS   AGE
po/default-http-backend-3138300093-k20m8   1/1       Running   0          6d
po/lb-myapp-292358823-7cs4k                1/1       Running   0          6d
po/lb-mysql-745212643-nlph1                1/1       Running   0          4d
po/myapp-2672690291-5bxsz                  1/1       Running   0          2d
po/myapp-2672690291-m7f86                  1/1       Running   0          2d
po/myapp-2672690291-trnz5                  1/1       Running   0          2d
po/mysql-0                                 1/1       Running   0          4d

NAME                       CLUSTER-IP      EXTERNAL-IP   PORT(S)                      AGE
svc/default-http-backend   10.233.56.33    <none>        80/TCP                       6d
svc/kubernetes             10.233.0.1      <none>        443/TCP                      6d
svc/myapp                  10.233.62.142   <none>        8088/TCP,4959/TCP,4950/TCP   6d
svc/mysql                  None            <none>        3306/TCP                     4d

NAME                 DESIRED   CURRENT   AGE
statefulsets/mysql   1         1         4d

NAME                          DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE
deploy/default-http-backend   1         1         1            1           6d
deploy/lb-myapp               1         1         1            1           6d
deploy/lb-mysql               1         1         1            1           4d
deploy/myapp                  3         3         3            3           6d

NAME                                 DESIRED   CURRENT   READY     AGE
rs/default-http-backend-3138300093   1         1         1         6d
rs/lb-myapp-292358823                1         1         1         6d
rs/lb-mysql-745212643                1         1         1         4d
rs/myapp-1205861895                  0         0         0         2d
rs/myapp-1362230792                  0         0         0         2d
rs/myapp-1654652426                  0         0         0         2d
rs/myapp-2672690291                  3         3         3         2d
rs/myapp-913440261                   0         0         0         6d

```
### 2.测试命令 kubectl logs

```
[root@node1 ~]# kubectl logs myapp-2672690291-trnz5
```
### 2.运行结果
```
I1201 10:01:09.606449       1 main.go:70] start echo server
I1201 10:01:09.606700       1 main.go:72] start time server
2017/12/01 10:01:09 [I] [asm_amd64.s:2197] http server Running on http://:8088
I1201 10:02:31.348058       1 main.go:40] Send time:Dec 1, 2017 at 10:02am (UTC)
I1201 10:02:32.587322       1 main.go:29] client said:123

```
### 3.测试命令 kubectl exec
```
[root@node1 ~]# kubectl exec myapp-2672690291-5bxsz echo "hello world"
```
### 3.运行结果
```
hello world
```
### 4.测试命令 kubectl describe
```
[root@node1 ~]# kubectl describe pod myapp-2672690291-5bxsz
```
### 4.运行结果
```
Name:           myapp-2672690291-5bxsz
Namespace:      default
Node:           node2/192.168.34.37
Start Time:     Fri, 01 Dec 2017 18:01:01 +0800
Labels:         ekos-application=myapp
                ekos-service=myapp
                log-ignore=false
                pod-template-hash=2672690291
                type=service
                version=5
Annotations:    kubernetes.io/created-by={"kind":"SerializedReference","apiVersion":"v1","reference":{"kind":"ReplicaSet","namespace":"default","name":"myapp-2672690291","uid":"895995a4-d67e-11e7-89df-001a4a1800d9",""
Status:         Running
IP:             10.233.75.33
Controllers:    ReplicaSet/myapp-2672690291
Containers:
  myapp:
    Container ID:       docker://4bc73837cfcbd96b891fbe2e5ed40f1f8fac28acd199db40cc35709de10a0a79
    Image:              registry.ekos.local/default/myapp:3.0
    Image ID:           docker-pullable://registry.ekos.local/default/myapp@sha256:eaa411bdfd4ef00793843d156fb2e3af6184f248a321916d7aff0069c4cd124a
    Port:
    Command:
      ./myapp
    Args:
      -logtostderr
      =
      true;
    State:              Running
      Started:          Fri, 01 Dec 2017 18:01:06 +0800
    Ready:              True
    Restart Count:      0
    Limits:
      cpu:      125m
      memory:   64M
    Requests:
      cpu:              62m
      memory:           32M
    Environment:        <none>
    Mounts:
      /var/run/secrets/kubernetes.io/serviceaccount from default-token-7nm62 (ro)
Conditions:
  Type          Status
  Initialized   True
  Ready         True
  PodScheduled  True
Volumes:
  default-token-7nm62:
    Type:       Secret (a volume populated by a Secret)
    SecretName: default-token-7nm62
    Optional:   false
QoS Class:      Burstable
Node-Selectors: <none>
Tolerations:    <none>
Events:         <none>
```
### 5. 测试命令 kubectl apply
```
[root@node1 ~]# kubectl apply -f pod.yaml
```
### 5.运行结果
```
pod "frontend" created
[root@node1 ~]# kubectl get pod
NAME                                    READY     STATUS              RESTARTS
default-http-backend-3138300093-k20m8   1/1       Running             0
frontend                                0/2       ContainerCreating   0
```
### 6.测试命令 kubectl create
```
[root@node1 ~]# kubectl create -f pod.yaml
```
### 6.运行结果
```
pod "frontend" created
[root@node1 ~]# kubectl get pods
NAME                                    READY     STATUS              RESTARTS   AGE
default-http-backend-3138300093-k20m8   1/1       Running             0          6d
frontend                                0/2       ContainerCreating   0   
```
### 7.测试命令 kubectl delete
```
[root@node1 ~]# kubectl delete pod frontend
```
### 7.运行结果
```
pod "frontend" deleted
```