#

## Task 1 — Local Kubernetes Setup

```bash
$ minikube start --driver=docker
😄  minikube v1.38.1 on Ubuntu 24.04
✨  Using the docker driver based on user configuration
❗  Starting v1.39.0, minikube will default to "containerd" container runtime. See #21973 for more info.
📌  Using Docker driver with root privileges
👍  Starting "minikube" primary control-plane node in "minikube" cluster
🚜  Pulling base image v0.0.50 ...
💾  Downloading Kubernetes v1.35.1 preload ...
    > preloaded-images-k8s-v18-v1...:  272.45 MiB / 272.45 MiB  100.00% 7.53 Mi
    > gcr.io/k8s-minikube/kicbase...:  519.58 MiB / 519.58 MiB  100.00% 6.18 Mi
🔥  Creating docker container (CPUs=2, Memory=3900MB) ...
🐳  Preparing Kubernetes v1.35.1 on Docker 29.2.1 ...
🔗  Configuring bridge CNI (Container Networking Interface) ...
🔎  Verifying Kubernetes components...
    ▪ Using image gcr.io/k8s-minikube/storage-provisioner:v5
🌟  Enabled addons: storage-provisioner, default-storageclass
🏄  Done! kubectl is now configured to use "minikube" cluster and "default" namespace by default
```

```bash
$ kubectl cluster-info
Kubernetes control plane is running at https://192.168.49.2:8443
CoreDNS is running at https://192.168.49.2:8443/api/v1/namespaces/kube-system/services/kube-dns:dns/proxy

To further debug and diagnose cluster problems, use 'kubectl cluster-info dump'.
```

```bash
$ kubectl get nodes
NAME       STATUS   ROLES           AGE   VERSION
minikube   Ready    control-plane   25m   v1.35.1
```

I chouse minikube because it sounds funny :)

## Task 2 — Application Deployment 

```bash
$ kubectl apply -f k8s/deployment.yml
deployment.apps/devops-info-service created
```

```bash
$ kubectl get deployments
NAME                  READY   UP-TO-DATE   AVAILABLE   AGE
devops-info-service   3/3     3            3           27s
```

```bash
$ kubectl get pods
NAME                                   READY   STATUS    RESTARTS   AGE
devops-info-service-78f44cdc7d-9ncw5   1/1     Running   0          33s
devops-info-service-78f44cdc7d-cxxpr   1/1     Running   0          33s
devops-info-service-78f44cdc7d-zwfgz   1/1     Running   0          33s
```

```bash
 kubectl describe deployment devops-info-service
Name:                   devops-info-service
Namespace:              default
CreationTimestamp:      Thu, 26 Mar 2026 16:47:06 +0300
Labels:                 app=devops-info-service
                        version=v1
Annotations:            deployment.kubernetes.io/revision: 1
Selector:               app=devops-info-service
Replicas:               3 desired | 3 updated | 3 total | 3 available | 0 unavailable
StrategyType:           RollingUpdate
MinReadySeconds:        0
RollingUpdateStrategy:  0 max unavailable, 1 max surge
Pod Template:
  Labels:  app=devops-info-service
           version=v1
  Containers:
   app:
    Image:      lehus1/devops-info-service:latest
    Port:       8000/TCP
    Host Port:  0/TCP
    Limits:
      cpu:     200m
      memory:  256Mi
    Requests:
      cpu:      100m
      memory:   128Mi
    Liveness:   http-get http://:8000/health delay=10s timeout=1s period=5s #success=1 #failure=3
    Readiness:  http-get http://:8000/health delay=5s timeout=1s period=3s #success=1 #failure=3
    Environment:
      PORT:        8000
      HOST:        0.0.0.0
    Mounts:        <none>
  Volumes:         <none>
  Node-Selectors:  <none>
  Tolerations:     <none>
Conditions:
  Type           Status  Reason
  ----           ------  ------
  Available      True    MinimumReplicasAvailable
  Progressing    True    NewReplicaSetAvailable
OldReplicaSets:  <none>
NewReplicaSet:   devops-info-service-78f44cdc7d (3/3 replicas created)
Events:
  Type    Reason             Age   From                   Message
  ----    ------             ----  ----                   -------
  Normal  ScalingReplicaSet  59s   deployment-controller  Scaled up replica set devops-info-service-78f44cdc7d from 0 to 3
```

## Task 3 — Service Configuration

```bash
$ kubectl apply -f k8s/service.yml
service/devops-info-service created
```

```bash
$ kubectl get services
NAME                  TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)        AGE
devops-info-service   NodePort    10.109.62.222   <none>        80:30080/TCP   23s
kubernetes            ClusterIP   10.96.0.1       <none>        443/TCP        81m
```

```bash
$ minikube service devops-info-service --url
http://192.168.49.2:30080
```

```bash
$ curl $(minikube service devops-info-service --url)/health
{"status":"healthy","timestamp":"2026-03-26T14:28:37.901535+00:00","uptime_seconds":2473}
```

```bash
$ kubectl get endpoints
Warning: v1 Endpoints is deprecated in v1.33+; use discovery.k8s.io/v1 EndpointSlice
NAME                  ENDPOINTS                                         AGE
devops-info-service   10.244.0.3:8000,10.244.0.4:8000,10.244.0.5:8000   71s
kubernetes            192.168.49.2:8443                                 82m
```

```bash
$ kubectl describe service devops-info-service
Name:                     devops-info-service
Namespace:                default
Labels:                   app=devops-info-service
Annotations:              <none>
Selector:                 app=devops-info-service
Type:                     NodePort
IP Family Policy:         SingleStack
IP Families:              IPv4
IP:                       10.109.62.222
IPs:                      10.109.62.222
Port:                     <unset>  80/TCP
TargetPort:               8000/TCP
NodePort:                 <unset>  30080/TCP
Endpoints:                10.244.0.5:8000,10.244.0.3:8000,10.244.0.4:8000
Session Affinity:         None
External Traffic Policy:  Cluster
Internal Traffic Policy:  Cluster
Events:                   <none>
```


## Task 4 — Scaling and Updates

### Increase number of Replicas to five
```bash
$ kubectl get pods -w
NAME                                   READY   STATUS    RESTARTS   AGE
devops-info-service-78f44cdc7d-9ncw5   1/1     Running   0          44m
devops-info-service-78f44cdc7d-cxxpr   1/1     Running   0          44m
devops-info-service-78f44cdc7d-zwfgz   1/1     Running   0          44m
```

```bash
$ kubectl apply -f k8s/deployment.yml
deployment.apps/devops-info-service configured
```

```bash
$ kubectl get pods -w
NAME                                   READY   STATUS    RESTARTS   AGE
devops-info-service-78f44cdc7d-9ncw5   1/1     Running   0          58m
devops-info-service-78f44cdc7d-cxxpr   1/1     Running   0          58m
devops-info-service-78f44cdc7d-qqsl6   1/1     Running   0          24s
devops-info-service-78f44cdc7d-vz67r   1/1     Running   0          24s
devops-info-service-78f44cdc7d-zwfgz   1/1     Running   0          58m
```

### Change tag of image

```bash
$ kubectl apply -f k8s/deployment.yml
deployment.apps/devops-info-service configured
```

```bash
$ kubectl rollout status deployment/devops-info-service
Waiting for deployment "devops-info-service" rollout to finish: 1 out of 5 new replicas have been updated...
Waiting for deployment "devops-info-service" rollout to finish: 1 out of 5 new replicas have been updated...
Waiting for deployment "devops-info-service" rollout to finish: 1 out of 5 new replicas have been updated...
Waiting for deployment "devops-info-service" rollout to finish: 2 out of 5 new replicas have been updated...
Waiting for deployment "devops-info-service" rollout to finish: 2 out of 5 new replicas have been updated...
Waiting for deployment "devops-info-service" rollout to finish: 2 out of 5 new replicas have been updated...
Waiting for deployment "devops-info-service" rollout to finish: 2 out of 5 new replicas have been updated...
Waiting for deployment "devops-info-service" rollout to finish: 3 out of 5 new replicas have been updated...
Waiting for deployment "devops-info-service" rollout to finish: 3 out of 5 new replicas have been updated...
Waiting for deployment "devops-info-service" rollout to finish: 3 out of 5 new replicas have been updated...
Waiting for deployment "devops-info-service" rollout to finish: 3 out of 5 new replicas have been updated...
Waiting for deployment "devops-info-service" rollout to finish: 4 out of 5 new replicas have been updated...
Waiting for deployment "devops-info-service" rollout to finish: 4 out of 5 new replicas have been updated...
Waiting for deployment "devops-info-service" rollout to finish: 4 out of 5 new replicas have been updated...
Waiting for deployment "devops-info-service" rollout to finish: 4 out of 5 new replicas have been updated...
Waiting for deployment "devops-info-service" rollout to finish: 1 old replicas are pending termination...
Waiting for deployment "devops-info-service" rollout to finish: 1 old replicas are pending termination...
Waiting for deployment "devops-info-service" rollout to finish: 1 old replicas are pending termination...
deployment "devops-info-service" successfully rolled out
```

### Rollback

```bash
$ kubectl rollout history deployment/devops-info-service
deployment.apps/devops-info-service 
REVISION  CHANGE-CAUSE
2         <none>
3         <none>
```

```bash
$ kubectl rollout undo deployment/devops-info-service
deployment.apps/devops-info-service rolled back
```

```bash
$ kubectl rollout status deployment/devops-info-service
Waiting for deployment "devops-info-service" rollout to finish: 2 out of 5 new replicas have been updated...
Waiting for deployment "devops-info-service" rollout to finish: 2 out of 5 new replicas have been updated...
Waiting for deployment "devops-info-service" rollout to finish: 2 out of 5 new replicas have been updated...
Waiting for deployment "devops-info-service" rollout to finish: 2 out of 5 new replicas have been updated...
Waiting for deployment "devops-info-service" rollout to finish: 3 out of 5 new replicas have been updated...
Waiting for deployment "devops-info-service" rollout to finish: 3 out of 5 new replicas have been updated...
Waiting for deployment "devops-info-service" rollout to finish: 3 out of 5 new replicas have been updated...
Waiting for deployment "devops-info-service" rollout to finish: 3 out of 5 new replicas have been updated...
Waiting for deployment "devops-info-service" rollout to finish: 4 out of 5 new replicas have been updated...
Waiting for deployment "devops-info-service" rollout to finish: 4 out of 5 new replicas have been updated...
Waiting for deployment "devops-info-service" rollout to finish: 4 out of 5 new replicas have been updated...
Waiting for deployment "devops-info-service" rollout to finish: 4 out of 5 new replicas have been updated...
Waiting for deployment "devops-info-service" rollout to finish: 1 old replicas are pending termination...
Waiting for deployment "devops-info-service" rollout to finish: 1 old replicas are pending termination...
Waiting for deployment "devops-info-service" rollout to finish: 1 old replicas are pending termination...
deployment "devops-info-service" successfully rolled out
```

### Verify zero downtime

```bash
while true; do curl -s $(minikube service devops-info-service --url)/health && echo " OK"; sleep 0.5; done; done
{"status":"healthy","timestamp":"2026-03-26T14:53:30.774033+00:00","uptime_seconds":168} OK
{"status":"healthy","timestamp":"2026-03-26T14:53:31.544985+00:00","uptime_seconds":169} OK
{"status":"healthy","timestamp":"2026-03-26T14:53:32.297618+00:00","uptime_seconds":163} OK
{"status":"healthy","timestamp":"2026-03-26T14:53:33.060646+00:00","uptime_seconds":171} OK
{"status":"healthy","timestamp":"2026-03-26T14:53:33.807437+00:00","uptime_seconds":171} OK
{"status":"healthy","timestamp":"2026-03-26T14:53:34.598875+00:00","uptime_seconds":186} OK
{"status":"healthy","timestamp":"2026-03-26T14:53:35.400136+00:00","uptime_seconds":173} OK
{"status":"healthy","timestamp":"2026-03-26T14:53:36.149010+00:00","uptime_seconds":167} OK
{"status":"healthy","timestamp":"2026-03-26T14:53:36.886982+00:00","uptime_seconds":160} OK
{"status":"healthy","timestamp":"2026-03-26T14:53:37.640144+00:00","uptime_seconds":189} OK
{"status":"healthy","timestamp":"2026-03-26T14:53:38.381779+00:00","uptime_seconds":169} OK
{"status":"healthy","timestamp":"2026-03-26T14:53:39.127880+00:00","uptime_seconds":190} OK
{"status":"healthy","timestamp":"2026-03-26T14:53:39.872070+00:00","uptime_seconds":191} OK
{"status":"healthy","timestamp":"2026-03-26T14:53:40.583139+00:00","uptime_seconds":185} OK
{"status":"healthy","timestamp":"2026-03-26T14:53:41.388603+00:00","uptime_seconds":164} OK
{"status":"healthy","timestamp":"2026-03-26T14:53:42.129699+00:00","uptime_seconds":165} OK
{"status":"healthy","timestamp":"2026-03-26T14:53:42.871593+00:00","uptime_seconds":166} OK
{"status":"healthy","timestamp":"2026-03-26T14:53:43.608531+00:00","uptime_seconds":195} OK
{"status":"healthy","timestamp":"2026-03-26T14:53:44.356240+00:00","uptime_seconds":175} OK
{"status":"healthy","timestamp":"2026-03-26T14:53:45.084074+00:00","uptime_seconds":176} OK
{"status":"healthy","timestamp":"2026-03-26T14:53:45.816317+00:00","uptime_seconds":176} OK
{"status":"healthy","timestamp":"2026-03-26T14:53:46.560764+00:00","uptime_seconds":177} OK
{"status":"healthy","timestamp":"2026-03-26T14:53:47.352954+00:00","uptime_seconds":192} OK
{"status":"healthy","timestamp":"2026-03-26T14:53:48.076184+00:00","uptime_seconds":4} OK
{"status":"healthy","timestamp":"2026-03-26T14:53:48.809865+00:00","uptime_seconds":11} OK
{"status":"healthy","timestamp":"2026-03-26T14:53:49.529014+00:00","uptime_seconds":194} OK
{"status":"healthy","timestamp":"2026-03-26T14:53:50.250050+00:00","uptime_seconds":181} OK
{"status":"healthy","timestamp":"2026-03-26T14:53:50.987109+00:00","uptime_seconds":196} OK
{"status":"healthy","timestamp":"2026-03-26T14:53:51.724458+00:00","uptime_seconds":14} OK
{"status":"healthy","timestamp":"2026-03-26T14:53:52.466483+00:00","uptime_seconds":15} OK
{"status":"healthy","timestamp":"2026-03-26T14:53:53.203042+00:00","uptime_seconds":198} OK
{"status":"healthy","timestamp":"2026-03-26T14:53:53.935039+00:00","uptime_seconds":199} OK
{"status":"healthy","timestamp":"2026-03-26T14:53:54.662029+00:00","uptime_seconds":199} OK
{"status":"healthy","timestamp":"2026-03-26T14:53:55.380279+00:00","uptime_seconds":18} OK
{"status":"healthy","timestamp":"2026-03-26T14:53:56.096142+00:00","uptime_seconds":5} OK
{"status":"healthy","timestamp":"2026-03-26T14:53:56.835480+00:00","uptime_seconds":187} OK
{"status":"healthy","timestamp":"2026-03-26T14:53:57.571255+00:00","uptime_seconds":20} OK
{"status":"healthy","timestamp":"2026-03-26T14:53:58.317528+00:00","uptime_seconds":7} OK
{"status":"healthy","timestamp":"2026-03-26T14:53:59.057614+00:00","uptime_seconds":190} OK
{"status":"healthy","timestamp":"2026-03-26T14:53:59.788573+00:00","uptime_seconds":205} OK
{"status":"healthy","timestamp":"2026-03-26T14:54:00.537531+00:00","uptime_seconds":205} OK
{"status":"healthy","timestamp":"2026-03-26T14:54:01.284539+00:00","uptime_seconds":206} OK
{"status":"healthy","timestamp":"2026-03-26T14:54:02.015400+00:00","uptime_seconds":18} OK
{"status":"healthy","timestamp":"2026-03-26T14:54:02.749065+00:00","uptime_seconds":193} OK
{"status":"healthy","timestamp":"2026-03-26T14:54:03.486626+00:00","uptime_seconds":12} OK
{"status":"healthy","timestamp":"2026-03-26T14:54:04.204010+00:00","uptime_seconds":6} OK
{"status":"healthy","timestamp":"2026-03-26T14:54:04.947298+00:00","uptime_seconds":7} OK
{"status":"healthy","timestamp":"2026-03-26T14:54:05.691922+00:00","uptime_seconds":196} OK
{"status":"healthy","timestamp":"2026-03-26T14:54:06.422558+00:00","uptime_seconds":15} OK
{"status":"healthy","timestamp":"2026-03-26T14:54:07.158453+00:00","uptime_seconds":198} OK
{"status":"healthy","timestamp":"2026-03-26T14:54:07.893381+00:00","uptime_seconds":198} OK
{"status":"healthy","timestamp":"2026-03-26T14:54:09.357422+00:00","uptime_seconds":4} OK
{"status":"healthy","timestamp":"2026-03-26T14:54:10.102891+00:00","uptime_seconds":19} OK
```

```bash
$ kubectl get all
NAME                                       READY   STATUS    RESTARTS   AGE
pod/devops-info-service-78f44cdc7d-9lbts   1/1     Running   0          3h17m
pod/devops-info-service-78f44cdc7d-gwbd9   1/1     Running   0          3h17m
pod/devops-info-service-78f44cdc7d-kbfjx   1/1     Running   0          3h17m
pod/devops-info-service-78f44cdc7d-ksxq4   1/1     Running   0          3h17m
pod/devops-info-service-78f44cdc7d-rdx7w   1/1     Running   0          3h17m

NAME                          TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)        AGE
service/devops-info-service   NodePort    10.109.62.222   <none>        80:30080/TCP   3h43m
service/kubernetes            ClusterIP   10.96.0.1       <none>        443/TCP        5h4m

NAME                                  READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/devops-info-service   5/5     5            5           4h23m

NAME                                             DESIRED   CURRENT   READY   AGE
replicaset.apps/devops-info-service-78f44cdc7d   5         5         5       4h23m
replicaset.apps/devops-info-service-fb7cdf9      0         0         0       3h23m
```

## Task 5

**Architecture Overview**
The deployment consists of:

- **3 replicas** (later scaled to 5) of the application pod.
- Each pod runs a single container exposing port 8000.
- A **NodePort Service** (`devops-info-service`) exposes the application on port 30080 of the cluster node.
- The application is accessible externally via `minikube service` or `kubectl port-forward`.

**Resource allocation**: 128Mi memory request / 256Mi limit, 100m CPU request / 200m limit per pod.

**Health checks**: Liveness and readiness probes on `/health` endpoint ensure pods are healthy and ready to serve traffic.

**Operatations**: I performed operations using only declarative way (modifying .yml files and applying them)

## Bonus task

```bash
$ kubectl apply -f k8s/deployment-go.yml
deployment.apps/devops-info-service-go created
```

```bash
$ kubectl apply -f k8s/service-go.yml
service/devops-info-service-go created
```

```bash
$ kubectl get services
NAME                     TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)        AGE
devops-info-service      NodePort    10.109.62.222    <none>        80:30080/TCP   4h33m
devops-info-service-go   NodePort    10.111.217.117   <none>        80:30090/TCP   31s
kubernetes               ClusterIP   10.96.0.1        <none>        443/TCP        5h54m
```

```bash
$ minikube service devops-info-service-go --url
http://192.168.49.2:30090
```

```bash
$ curl $(minikube service devops-info-service-go --url)/health
{"status":"healthy","timestamp":"2026-03-26T19:03:31Z","uptime_seconds":172}
```

```bash
$ minikube addons enable ingress
💡  ingress is an addon maintained by Kubernetes. For any concerns contact minikube on GitHub.
You can view the list of minikube maintainers at: https://github.com/kubernetes/minikube/blob/master/OWNERS
    ▪ Using image registry.k8s.io/ingress-nginx/kube-webhook-certgen:v1.6.7
    ▪ Using image registry.k8s.io/ingress-nginx/controller:v1.14.3
    ▪ Using image registry.k8s.io/ingress-nginx/kube-webhook-certgen:v1.6.7
🔎  Verifying ingress addon...
🌟  The 'ingress' addon is enabled
```

```bash
$ kubectl apply -f k8s/ingress.yml
ingress.networking.k8s.io/apps-ingress created
```

```bash
$ curl http://local.example.com/app1/health
{"service":{"name":"devops-info-service","version":"1.0.0","description":"DevOps course info service","framework":"FastAPI"},"system":{"hostname":"devops-info-service-78f44cdc7d-9lbts","platform":"Linux","platform_version":"#14~24.04.1-Ubuntu SMP PREEMPT_DYNAMIC Thu Jan 15 15:52:10 UTC 2","architecture":"x86_64","cpu_count":16,"python_version":"3.13.12"},"runtime":{"uptime_seconds":15386,"uptime_human":"4 hours, 16 minutes","current_time":"2026-03-26T19:10:31.677875+00:00","timezone":"UTC"},"request":{"client_ip":"10.244.0.35","user_agent":"curl/8.5.0","method":"GET","path":"/"},"endpoints":[{"path":"/","method":"GET","description":"Service information"},{"path":"/health","method":"GET","description":"Health check"},{"path":"/docs","method":"GET","description":"Auto-generated API documentation"}]}
```

```bash
$ curl http://local.example.com/app2/health
{"service":{"name":"devops-info-service","version":"1.0.0","description":"DevOps course info service","framework":"Go net/http"},"system":{"hostname":"devops-info-service-go-686d4f5dcd-9h44m","platform":"linux","platform_version":"linux/amd64","architecture":"amd64","cpu_count":16,"go_version":"go1.21.13"},"runtime":{"uptime_seconds":636,"uptime_human":"0 hours, 10 minutes","current_time":"2026-03-26T19:11:20Z","timezone":"UTC"},"request":{"client_ip":"192.168.49.1","user_agent":"curl/8.5.0","method":"GET","path":"/"},"endpoints":[{"path":"/","method":"GET","description":"Service information"},{"path":"/health","method":"GET","description":"Health check"}]}
```

**Add TLS**

```bash
$ kubectl apply -f k8s/ingress.yml
ingress.networking.k8s.io/apps-ingress configured
```

```bash
$ curl http://local.example.com/app2/health
<html>
<head><title>308 Permanent Redirect</title></head>
<body>
<center><h1>308 Permanent Redirect</h1></center>
<hr><center>nginx</center>
</body>
</html>
```

```bash
$ curl https://local.example.com/app2/health
curl: (60) SSL certificate problem: self-signed certificate
More details here: https://curl.se/docs/sslcerts.html

curl failed to verify the legitimacy of the server and therefore could not
establish a secure connection to it. To learn more about this situation and
how to fix it, please visit the web page mentioned above.
```

```bash
$ curl -k https://local.example.com/app2/health
{"service":{"name":"devops-info-service","version":"1.0.0","description":"DevOps course info service","framework":"Go net/http"},"system":{"hostname":"devops-info-service-go-686d4f5dcd-nxgbh","platform":"linux","platform_version":"linux/amd64","architecture":"amd64","cpu_count":16,"go_version":"go1.21.13"},"runtime":{"uptime_seconds":1739,"uptime_human":"0 hours, 28 minutes","current_time":"2026-03-26T19:29:37Z","timezone":"UTC"},"request":{"client_ip":"192.168.49.1","user_agent":"curl/8.5.0","method":"GET","path":"/"},"endpoints":[{"path":"/","method":"GET","description":"Service information"},{"path":"/health","method":"GET","description":"Health check"}]}
```

```bash
$ curl -k https://local.example.com/app1/health
{"service":{"name":"devops-info-service","version":"1.0.0","description":"DevOps course info service","framework":"FastAPI"},"system":{"hostname":"devops-info-service-78f44cdc7d-ksxq4","platform":"Linux","platform_version":"#14~24.04.1-Ubuntu SMP PREEMPT_DYNAMIC Thu Jan 15 15:52:10 UTC 2","architecture":"x86_64","cpu_count":16,"python_version":"3.13.12"},"runtime":{"uptime_seconds":16545,"uptime_human":"4 hours, 35 minutes","current_time":"2026-03-26T19:29:43.023055+00:00","timezone":"UTC"},"request":{"client_ip":"10.244.0.35","user_agent":"curl/8.5.0","method":"GET","path":"/"},"endpoints":[{"path":"/","method":"GET","description":"Service information"},{"path":"/health","method":"GET","description":"Health check"},{"path":"/docs","method":"GET","description":"Auto-generated API documentation"}]}
```

### Docs:

For the bonus task, I extended the setup to route traffic to two different applications via a single Ingress controller, secured with a self‑signed TLS certificate.

**Production Considerations**:

- Self‑signed certificates are only for development; in production use a trusted CA (e.g., Let's Encrypt) and cert-manager for automatic renewal.

- The Ingress controller’s default redirect from HTTP to HTTPS is desirable for security.

- For production, also set appropriate timeouts, rate limits, and consider using the Gateway API as the successor to Ingress.


