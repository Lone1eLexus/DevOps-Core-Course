# Lab 15 — StatefulSets & Persistent Storage

## Task 1 — StatefulSet Concepts

### StatefulSet Guarantees
- Stable, unique network identifiers → pod names with ordinal index
- Stable, persistent storage → per‑pod PVCs
- Ordered, graceful deployment → sequential create/delete

### Deployment vs StatefulSet

| Feature | Deployment | StatefulSet |
|---------|------------|-------------|
| Pod naming | Random suffix (e.g., myapp‑xyzab) | Predictable (myapp‑0, myapp‑1) |
| Storage | Shared PVC or ephemeral | Each pod its own PVC |
| Scaling order | Arbitrary | Sequential (0→1→2) |
| Network identity | Not stable (each restart may change) | Stable DNS name |
| Use cases | Stateless apps, APIs, web servers | Databases, message queues, distributed systems |

**Examples of stateful workloads**: PostgreSQL, MySQL, Kafka, Elasticsearch, etc.

### Headless Services
- A service with `clusterIP: None`
- Does **not** provide load‑balancing; instead returns DNS A records for each pod
- Enables direct pod‑to‑pod communication using stable hostnames
- Required for StatefulSets to allow pod discovery

**DNS naming pattern**:  
`<statefulset-name>-<ordinal>.<headless-service-name>.<namespace>.svc.cluster.local`

## Task 2 — Convert Deployment to StatefulSet

### Implementation

We added two new templates to the Helm chart:

- `templates/headless-service.yaml` – defines a Service with `clusterIP: None`
- `templates/statefulset.yaml` – defines the StatefulSet with `volumeClaimTemplates`

A `statefulset.enabled` flag in `values.yaml` controls whether to deploy the StatefulSet (instead of the Deployment).


```bash
$ helm upgrade --install app-stateful . -f values-statefulset.yaml
Release "app-stateful" has been upgraded. Happy Helming!
NAME: app-stateful
LAST DEPLOYED: Thu May  7 21:51:00 2026
NAMESPACE: default
STATUS: deployed
REVISION: 3
DESCRIPTION: Upgrade complete
TEST SUITE: None
$ kubectl get statefulset
NAME                               READY   AGE
app-stateful-devops-info-service   3/3     2m1s
vault                              1/1     28d
$ kubectl get pods -l app.kubernetes.io/instance=app-stateful
NAME                                               READY   STATUS    RESTARTS   AGE
app-stateful-devops-info-service-0                 1/1     Running   0          2m10s
app-stateful-devops-info-service-1                 1/1     Running   0          2m2s
app-stateful-devops-info-service-2                 1/1     Running   0          115s
app-stateful-devops-info-service-5cc78d9d6-b79ds   1/1     Running   0          2m10s
app-stateful-devops-info-service-5cc78d9d6-gqkcq   1/1     Running   0          2m5s
app-stateful-devops-info-service-5cc78d9d6-lpcvz   1/1     Running   0          2m2s
$ kubectl get pvc
NAME                                      STATUS   VOLUME                                     CAPACITY   ACCESS MODES   STORAGECLASS   VOLUMEATTRIBUTESCLASS   AGE
app-stateful-devops-info-service-data     Bound    pvc-acea9eaf-feab-4b58-b6fe-f7deb8df2be8   1Gi        RWO            standard       <unset>                 114s
data-app-stateful-devops-info-service-0   Bound    pvc-fb529325-ccc2-4980-bed9-e35f75d15a2d   1Gi        RWO            standard       <unset>                 6m34s
data-app-stateful-devops-info-service-1   Bound    pvc-0e8f5f9c-b669-4b06-8b9b-2fa07d25c706   1Gi        RWO            standard       <unset>                 6m24s
data-app-stateful-devops-info-service-2   Bound    pvc-1a456e03-3c77-4da6-889b-c404fbf89cce   1Gi        RWO            standard       <unset>                 6m14s
devops-app-devops-info-service-data       Bound    pvc-30352fe3-891b-4287-8ccf-54482f0a8e89   100Mi      RWO            standard       <unset>                 14d
```

## Task 3 — Headless Service & Pod Identity

```bash
$ kubectl run -it --rm debug --image=busybox --restart=Never -- /bin/sh
All commands and output from this session will be recorded in container logs, including credentials and sensitive information passed through the command prompt.
If you don't see a command prompt, try pressing enter.

/ # nslookup app-stateful-devops-info-service-1.headless
Server:         10.96.0.10
Address:        10.96.0.10:53

** server can't find app-stateful-devops-info-service-1.headless: NXDOMAIN

** server can't find app-stateful-devops-info-service-1.headless: NXDOMAIN

/ # 
```

```bash
 kubectl exec app-stateful-devops-info-service-0 -- sh -c 'echo "pod-0-data" > /data/pod-id.txt'
$ kubectl exec app-stateful-devops-info-service-1 -- sh -c 'echo "pod-1-data" > /data/pod-id.txt'
$ kubectl exec app-stateful-devops-info-service-2 -- sh -c 'echo "pod-2-data" > /data/pod-id.txt'
```

```bash
$ kubectl delete pod app-stateful-devops-info-service-0
pod "app-stateful-devops-info-service-0" deleted from default namespace
$ kubectl get pods -w
NAME                                                READY   STATUS    RESTARTS        AGE
app-stateful-devops-info-service-0                  0/1     Running   0               5s
app-stateful-devops-info-service-1                  1/1     Running   0               54m
app-stateful-devops-info-service-2                  1/1     Running   0               54m
app-stateful-devops-info-service-5cc78d9d6-b79ds    1/1     Running   0               54m
app-stateful-devops-info-service-5cc78d9d6-gqkcq    1/1     Running   0               54m
app-stateful-devops-info-service-5cc78d9d6-lpcvz    1/1     Running   0               54m
dev-app-bg-devops-info-service-post-install-2smjd   0/1     Error     0               6d23h
dev-app-bg-devops-info-service-post-install-8hnp7   0/1     Error     0               6d23h
dev-app-bg-devops-info-service-post-install-dcxvd   0/1     Error     0               6d22h
dev-app-bg-devops-info-service-post-install-fdv9q   0/1     Error     0               6d23h
dev-app-bg-devops-info-service-post-install-jhfl8   0/1     Error     0               6d23h
dev-app-bg-devops-info-service-post-install-s9zbl   0/1     Error     0               6d22h
dev-app-bg-devops-info-service-post-install-td5cc   0/1     Error     0               6d22h
devops-app-devops-info-service-8996948df-4khxm      1/1     Running   2 (6d22h ago)   14d
devops-app-devops-info-service-8996948df-g9fll      1/1     Running   2 (6d22h ago)   14d
devops-app-devops-info-service-8996948df-knrkn      1/1     Running   2 (6d22h ago)   14d
devops-info-service-78f44cdc7d-9lbts                1/1     Running   6 (6d22h ago)   42d
devops-info-service-78f44cdc7d-gwbd9                1/1     Running   6 (6d22h ago)   42d
devops-info-service-78f44cdc7d-kbfjx                1/1     Running   6 (6d22h ago)   42d
devops-info-service-78f44cdc7d-ksxq4                1/1     Running   6 (6d22h ago)   42d
devops-info-service-78f44cdc7d-rdx7w                1/1     Running   6 (6d22h ago)   42d
devops-info-service-go-686d4f5dcd-9h44m             1/1     Running   6 (6d22h ago)   42d
devops-info-service-go-686d4f5dcd-n2t2f             1/1     Running   6 (6d22h ago)   42d
devops-info-service-go-686d4f5dcd-nxgbh             1/1     Running   6 (6d22h ago)   42d
devops-info-service-go-686d4f5dcd-qd4nx             1/1     Running   6 (6d22h ago)   42d
devops-info-service-go-686d4f5dcd-zf92x             1/1     Running   6 (6d22h ago)   42d
vault-0                                             1/1     Running   4 (6d22h ago)   28d
vault-agent-injector-848dd747d7-znxsq               1/1     Running   4 (6d22h ago)   28d
app-stateful-devops-info-service-0                  1/1     Running   0               10s
$ kubectl exec app-stateful-devops-info-service-0 -- cat /data/pod-id.txt
pod-0-data
```

```bash
$ kubectl port-forward pod/app-stateful-devops-info-service-0 8080:8000 &
...
```

```bash
$ curl http://localhost:8080/visits
{"visits":0}
$ curl http://localhost:8080/
...
$ curl http://localhost:8080/
...
$ curl http://localhost:8080/
...
$ curl http://localhost:8080/visits
{"visits":3}
```

```bash
$ kubectl port-forward pod/app-stateful-devops-info-service-1 8081:8000 &
...
```

```bash
$ curl http://localhost:8080/visits
{"visits":0}
```




