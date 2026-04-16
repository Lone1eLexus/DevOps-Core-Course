# Lab 12 — ConfigMaps & Persistent Volumes

## Task 1 — Application Persistence Upgrade

```bash
$ docker build -t lehus/devops-info-service:latest .
...
$ docker compose up -d
...
```

```bash
$ curl http://localhost:8000/
...
$ curl http://localhost:8000/
...
$ curl http://localhost:8000/
...
```

```bash
curl http://localhost:8000/visits
{"visits":3}
$ cat data/visits
3
```

```bash
$ docker compose down
...
$ docker compose up -d
...
```

```bash
$ curl http://localhost:8000/
...
$ curl http://localhost:8000/visits
{"visits":4}
```

## Task 2 — ConfigMaps

```bash
helm install devops-app-helm . \
  --no-hooks \
  --set vault.enabled=false \
  --set environment=production \
  --set logLevel=warn \
  --set cacheTTL=7200 \
  --set featureDebug=false
```

```bash
$ POD=$(kubectl get pod -l app.kubernetes.io/instance=devops-app-helm -o jsonpath='{.items[0].metadata.name}')
$ kubectl exec $POD -- cat /config/config.json

{
  "app_name": "devops-info-service",
  "environment": "production",
  "features": {
    "debug_mode": false,
    "cache_enabled": true,
    "cache_ttl_seconds": 3600
  },
  "logging": {
    "level": "INFO",
    "format": "json"
  }
}

$ kubectl exec $POD -- printenv | grep -E "APP_ENV|LOG_LEVEL|CACHE_TTL|FEATURE_DEBUG""
LOG_LEVEL=warn
APP_ENV=production
CACHE_TTL=7200
FEATURE_DEBUG=false

$ kubectl get configmap -l app.kubernetes.io/instance=devops-app-helm
NAME                                              DATA   AGE
devops-app-helm-devops-info-service-config-file   1      2m46s
devops-app-helm-devops-info-service-env           4      2m46s
```

## Task 3 — Persistent Volumes 

```bash
$ kubectl port-forward service/devops-app-helm-devops-info-service 8080:80
...
```

```bash
$ curl http://localhost:8080/
...
$ curl http://localhost:8080/
...
$ curl http://localhost:8080/
...
$ curl http://localhost:8080/
...
$ curl http://localhost:8080/
...
$ curl http://localhost:8080/visits
{"visits":5}
```

```bash
$ kubectl get pvc
NAME                                       STATUS   VOLUME                                     CAPACITY   ACCESS MODES   STORAGECLASS   VOLUMEATTRIBUTESCLASS   AGE
devops-app-helm-devops-info-service-data   Bound    pvc-c0290cbb-2ea2-4a78-b339-bc0c79590125   100Mi      RWO            standard       <unset>                 7m20s
$ kubectl describe pvc devops-app-helm-devops-info-service-data
Name:          devops-app-helm-devops-info-service-data
Namespace:     default
StorageClass:  standard
Status:        Bound
Volume:        pvc-c0290cbb-2ea2-4a78-b339-bc0c79590125
Labels:        app.kubernetes.io/instance=devops-app-helm
               app.kubernetes.io/managed-by=Helm
               app.kubernetes.io/name=devops-info-service
               app.kubernetes.io/version=1.0.0
               helm.sh/chart=devops-info-service-0.1.0
Annotations:   meta.helm.sh/release-name: devops-app-helm
               meta.helm.sh/release-namespace: default
               pv.kubernetes.io/bind-completed: yes
               pv.kubernetes.io/bound-by-controller: yes
               volume.beta.kubernetes.io/storage-provisioner: k8s.io/minikube-hostpath
               volume.kubernetes.io/storage-provisioner: k8s.io/minikube-hostpath
Finalizers:    [kubernetes.io/pvc-protection]
Capacity:      100Mi
Access Modes:  RWO
VolumeMode:    Filesystem
Used By:       devops-app-helm-devops-info-service-7d75954975-4xdfv
               devops-app-helm-devops-info-service-7d75954975-gq6ql
               devops-app-helm-devops-info-service-7d75954975-qt9fr
Events:
  Type    Reason                 Age    From                                                                    Message
  ----    ------                 ----   ----                                                                    -------
  Normal  ExternalProvisioning   7m37s  persistentvolume-controller                                             Waiting for a volume to be created either by the external provisioner 'k8s.io/minikube-hostpath' or manually by the system administrator. If volume creation is delayed, please verify that the provisioner is running and correctly registered.
  Normal  Provisioning           7m37s  k8s.io/minikube-hostpath_minikube_f0f1d3cf-2b71-48bf-9adb-15b1181b4cd6  External provisioner is provisioning volume for claim "default/devops-app-helm-devops-info-service-data"
  Normal  ProvisioningSucceeded  7m37s  k8s.io/minikube-hostpath_minikube_f0f1d3cf-2b71-48bf-9adb-15b1181b4cd6  Successfully provisioned volume pvc-c0290cbb-2ea2-4a78-b339-bc0c79590125
$ kubectl exec devops-app-helm-devops-info-service-7d75954975-4xdfv -- cat /data/visits
5
```

## Task 4

| Feature | ConfigMap | Secret |
|---------|-----------|--------|
| **Purpose** | Non‑sensitive configuration | Sensitive data |
| **Encoding** | Plain text | Base64 (not encryption) |
| **Size limit** | 1MB | 1MB |
| **Use for** | URLs, feature flags, log levels | Passwords, tokens, keys |

> **Always use Secrets for credentials**, even if base64 is not real encryption – it adds a layer of accidental exposure prevention.

## Bonus task

```bash
$  kubectl exec devops-app-helm-devops-info-service-7d75954975-4xdfv-- cat /config/config.json | jq '.environment''
"production"
```

```bash
# Patch the ConfigMap to change the 'environment' field
kubectl patch configmap devops-app-helm-devops-info-service-config-file --patch '
data:
  config.json: |
    {
      "app_name": "devops-info-service",
      "environment": "staging",
      "features": {
        "debug_mode": false,
        "cache_enabled": true,
        "cache_ttl_seconds": 3600
      },
      "logging": {
        "level": "INFO",
        "format": "json"
      }
    }
'
```

```bash
$ kubectl exec devops-app-helm-devops-info-service-7d75954975-4xdfv -- cat /config/config.json | jq '.environment'
"production"
# After less than 40 sec:
$ kubectl exec devops-app-helm-devops-info-service-7d75954975-4xdfv -- cat /config/config.json | jq '.environment'
"staging"
```

```bash
$ kubectl delete configmap devops-app-helm-devops-info-service-config-file
configmap "devops-app-helm-devops-info-service-config-file" deleted from default namespace
$ helm upgrade devops-app-helm . --no-hooks --set vault.enabled=false
...
```

```bash
$ helm upgrade devops-app-helm . --no-hooks --set vault.enabled=false
Release "devops-app-helm" has been upgraded. Happy Helming!
NAME: devops-app-helm
LAST DEPLOYED: Thu Apr 16 21:21:56 2026
NAMESPACE: default
STATUS: deployed
REVISION: 5
DESCRIPTION: Upgrade complete
TEST SUITE: None
```

```bash
$ kubectl exec devops-app-helm-devops-info-service-b7d7dfc48-cprqc  -- cat /config/config.json | jq '.environment'
"production"
```

```bash
$ helm upgrade devops-app-helm . --set vault.enabled=false
Release "devops-app-helm" has been upgraded. Happy Helming!
NAME: devops-app-helm
LAST DEPLOYED: Thu Apr 16 21:44:07 2026
NAMESPACE: default
STATUS: deployed
REVISION: 7
DESCRIPTION: Upgrade complete
TEST SUITE: None
```

```bash
$ kubectl rollout status deployment/devops-app-helm-devops-info-service
Waiting for deployment "devops-app-helm-devops-info-service" rollout to finish: 1 out of 3 new replicas have been updated...
Waiting for deployment "devops-app-helm-devops-info-service" rollout to finish: 1 out of 3 new replicas have been updated...
Waiting for deployment "devops-app-helm-devops-info-service" rollout to finish: 1 out of 3 new replicas have been updated...
Waiting for deployment "devops-app-helm-devops-info-service" rollout to finish: 2 out of 3 new replicas have been updated...
Waiting for deployment "devops-app-helm-devops-info-service" rollout to finish: 2 out of 3 new replicas have been updated...
Waiting for deployment "devops-app-helm-devops-info-service" rollout to finish: 2 out of 3 new replicas have been updated...
Waiting for deployment "devops-app-helm-devops-info-service" rollout to finish: 2 out of 3 new replicas have been updated...
Waiting for deployment "devops-app-helm-devops-info-service" rollout to finish: 1 old replicas are pending termination...
Waiting for deployment "devops-app-helm-devops-info-service" rollout to finish: 1 old replicas are pending termination...
Waiting for deployment "devops-app-helm-devops-info-service" rollout to finish: 1 old replicas are pending termination...
deployment "devops-app-helm-devops-info-service" successfully rolled out
```

```bash
$ kubectl exec devops-app-helm-devops-info-service-745bbdbc7-cw94v  -- cat /config/config.json | jq '.environment
'
"staging"
```

### Update Behavior Analysis

| ConfigMap Mount Type | Auto-Update? | Delay | Notes |
|---------------------|--------------|-------|-------|
| Directory mount (`/config`) | ✅ Yes | ~60-120s | Via kubelet symlink refresh |
| `subPath` file mount | ❌ No | N/A | File is copied, not linked |
| Environment variables (`envFrom`) | ❌ No | N/A | Requires pod restart |

### subPath Limitation

When using `subPath`, Kubernetes performs a file copy at mount time instead of creating a symlink. This means:
- Changes to the ConfigMap are **not reflected** in the pod
- Useful for static configuration files that won't change
- Avoid for dynamic configuration that needs runtime updates

**Recommendation**: Always mount ConfigMaps as directories when auto-update is needed.