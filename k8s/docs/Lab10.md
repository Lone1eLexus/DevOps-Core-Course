# Lab 10 — Helm Package Manager

## Task 1 — Helm Fundamentals

```bash
$ curl -fsSL -o get_helm.sh https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-4
chmod 700 get_helm.sh
./get_helm.sh
Downloading https://get.helm.sh/helm-v4.1.3-linux-amd64.tar.gz
Verifying checksum... Done.
Preparing to install helm into /usr/local/bin
helm installed into /usr/local/bin/helm
```

```bash
$ helm version
version.BuildInfo{Version:"v4.1.3", GitCommit:"c94d381b03be117e7e57908edbf642104e00eb8f", GitTreeState:"clean", GoVersion:"go1.25.8", KubeClientVersion:"v1.35"}
```

```bash
$ helm show chart prometheus-community/prometheus
annotations:
  artifacthub.io/license: Apache-2.0
  artifacthub.io/links: |
    - name: Chart Source
      url: https://github.com/prometheus-community/helm-charts
    - name: Upstream Project
      url: https://github.com/prometheus/prometheus
apiVersion: v2
appVersion: v3.11.0
dependencies:
- condition: alertmanager.enabled
  name: alertmanager
  repository: https://prometheus-community.github.io/helm-charts
  version: 1.34.*
- condition: kube-state-metrics.enabled
  name: kube-state-metrics
  repository: https://prometheus-community.github.io/helm-charts
  version: 7.2.*
- condition: prometheus-node-exporter.enabled
  name: prometheus-node-exporter
  repository: https://prometheus-community.github.io/helm-charts
  version: 4.52.*
- condition: prometheus-pushgateway.enabled
  name: prometheus-pushgateway
  repository: https://prometheus-community.github.io/helm-charts
  version: 3.6.*
description: Prometheus is a monitoring system and time series database.
home: https://prometheus.io/
icon: https://raw.githubusercontent.com/prometheus/prometheus.github.io/master/assets/prometheus_logo-cb55bb5c346.png
keywords:
- monitoring
- prometheus
kubeVersion: '>=1.19.0-0'
maintainers:
- email: gianrubio@gmail.com
  name: gianrubio
  url: https://github.com/gianrubio
- email: zanhsieh@gmail.com
  name: zanhsieh
  url: https://github.com/zanhsieh
- email: miroslav.hadzhiev@gmail.com
  name: Xtigyro
  url: https://github.com/Xtigyro
- email: naseem@transit.app
  name: naseemkullah
  url: https://github.com/naseemkullah
- email: rootsandtrees@posteo.de
  name: zeritti
  url: https://github.com/zeritti
name: prometheus
sources:
- https://github.com/prometheus/alertmanager
- https://github.com/prometheus/prometheus
- https://github.com/prometheus/pushgateway
- https://github.com/prometheus/node_exporter
- https://github.com/kubernetes/kube-state-metrics
type: application
version: 28.15.0
```

### Helm Fundamentals

- Helm version: v4.x.x
- Chart inspected: bitnami/nginx
  - Chart version: 15.x.x
  - App version: 1.27.x
- A **chart** is a collection of templates that describe Kubernetes resources.
- A **release** is a specific instance of a chart running in a cluster.
- **Values** are configurable parameters that allow the same chart to be deployed differently across environments.
- Helm solves YAML duplication, provides versioning, and enables repeatable deployments.

## Task 2 — Create Your Helm Chart

```bash
$ helm lint .
==> Linting .
[INFO] Chart.yaml: icon is recommended

1 chart(s) linted, 0 chart(s) failed
```

```bash
$ helm template test-release .
---
# Source: devops-info-service/templates/service.yml
apiVersion: v1
kind: Service
metadata:
  name: test-release-devops-info-service
  labels:
    helm.sh/chart: devops-info-service-0.1.0
    app.kubernetes.io/name: devops-info-service
    app.kubernetes.io/instance: test-release
    app.kubernetes.io/version: "1.0.0"
    app.kubernetes.io/managed-by: Helm
spec:
  type: NodePort
  selector:
    app.kubernetes.io/name: devops-info-service
    app.kubernetes.io/instance: test-release
  ports:
    - port: 80
      targetPort: 8000
      protocol: TCP
      name: http
---
# Source: devops-info-service/templates/deployment.yml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: test-release-devops-info-service
  labels:
    helm.sh/chart: devops-info-service-0.1.0
    app.kubernetes.io/name: devops-info-service
    app.kubernetes.io/instance: test-release
    app.kubernetes.io/version: "1.0.0"
    app.kubernetes.io/managed-by: Helm
spec:
  replicas: 3
  selector:
    matchLabels:
      app.kubernetes.io/name: devops-info-service
      app.kubernetes.io/instance: test-release
  template:
    metadata:
      labels:
        app.kubernetes.io/name: devops-info-service
        app.kubernetes.io/instance: test-release
    spec:
      containers:
      - name: devops-info-service
        image: "lehus1/devops-info-service:latest"
        imagePullPolicy: IfNotPresent
        ports:
        - containerPort: 8000
          name: http
        env:
        - name: PORT
          value: "8000"
        resources:
          limits:
            cpu: 500m
            memory: 512Mi
          requests:
            cpu: 200m
            memory: 256Mi
        livenessProbe:
          httpGet:
            path: /health
            port: 8000
          initialDelaySeconds: 10
          periodSeconds: 5
        readinessProbe:
          httpGet:
            path: /health
            port: 8000
          initialDelaySeconds: 5
          periodSeconds: 3
```

```bash
$ helm install --dry-run --debug test-release .
level=WARN msg="--dry-run is deprecated and should be replaced with '--dry-run=client'"
level=DEBUG msg="Original chart version" version=""
level=DEBUG msg="Chart path" path=/home/lord/DevOps/DevOps-Core-Course/k8s/devops-chart
level=DEBUG msg="number of dependencies in the chart" chart=devops-info-service dependencies=0
NAME: test-release
LAST DEPLOYED: Thu Apr  2 20:12:49 2026
NAMESPACE: default
STATUS: pending-install
REVISION: 1
DESCRIPTION: Dry run complete
TEST SUITE: None
USER-SUPPLIED VALUES:
{}

COMPUTED VALUES:
env:
- name: PORT
  value: "8000"
image:
  pullPolicy: IfNotPresent
  repository: lehus1/devops-info-service
  tag: latest
livenessProbe:
  enabled: true
  httpGet:
    path: /health
    port: 8000
  initialDelaySeconds: 10
  periodSeconds: 5
readinessProbe:
  enabled: true
  httpGet:
    path: /health
    port: 8000
  initialDelaySeconds: 5
  periodSeconds: 3
replicaCount: 3
resources:
  limits:
    cpu: 500m
    memory: 512Mi
  requests:
    cpu: 200m
    memory: 256Mi
service:
  nodePort: 30081
  port: 80
  targetPort: 8000
  type: NodePort

HOOKS:
MANIFEST:
---
# Source: devops-info-service/templates/service.yml
apiVersion: v1
kind: Service
metadata:
  name: test-release-devops-info-service
  labels:
    helm.sh/chart: devops-info-service-0.1.0
    app.kubernetes.io/name: devops-info-service
    app.kubernetes.io/instance: test-release
    app.kubernetes.io/version: "1.0.0"
    app.kubernetes.io/managed-by: Helm
spec:
  type: NodePort
  selector:
    app.kubernetes.io/name: devops-info-service
    app.kubernetes.io/instance: test-release
  ports:
    - port: 80
      targetPort: 8000
      protocol: TCP
      name: http
---
# Source: devops-info-service/templates/deployment.yml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: test-release-devops-info-service
  labels:
    helm.sh/chart: devops-info-service-0.1.0
    app.kubernetes.io/name: devops-info-service
    app.kubernetes.io/instance: test-release
    app.kubernetes.io/version: "1.0.0"
    app.kubernetes.io/managed-by: Helm
spec:
  replicas: 3
  selector:
    matchLabels:
      app.kubernetes.io/name: devops-info-service
      app.kubernetes.io/instance: test-release
  template:
    metadata:
      labels:
        app.kubernetes.io/name: devops-info-service
        app.kubernetes.io/instance: test-release
    spec:
      containers:
      - name: devops-info-service
        image: "lehus1/devops-info-service:latest"
        imagePullPolicy: IfNotPresent
        ports:
        - containerPort: 8000
          name: http
        env:
        - name: PORT
          value: "8000"
        resources:
          limits:
            cpu: 500m
            memory: 512Mi
          requests:
            cpu: 200m
            memory: 256Mi
        livenessProbe:
          httpGet:
            path: /health
            port: 8000
          initialDelaySeconds: 10
          periodSeconds: 5
        readinessProbe:
          httpGet:
            path: /health
            port: 8000
          initialDelaySeconds: 5
          periodSeconds: 3
```

```bash
$ helm install devops-app-helm .
NAME: devops-app-helm
LAST DEPLOYED: Thu Apr  2 20:13:48 2026
NAMESPACE: default
STATUS: deployed
REVISION: 1
DESCRIPTION: Install complete
TEST SUITE: None
```

```bash
$ kubectl get endpoints
Warning: v1 Endpoints is deprecated in v1.33+; use discovery.k8s.io/v1 EndpointSlice
NAME                                  ENDPOINTS                                                        AGE
devops-app-helm-devops-info-service   10.244.0.51:8000,10.244.0.52:8000,10.244.0.53:8000               36s
devops-info-service                   10.244.0.36:8000,10.244.0.39:8000,10.244.0.42:8000 + 2 more...   7d2h
devops-info-service-go                10.244.0.38:8000,10.244.0.40:8000,10.244.0.41:8000 + 2 more...   6d22h
kubernetes                            192.168.49.2:8443 
```


## Task 3 — Multi-Environment Support

```bash
$ helm install devops-app-helm-dev . -f values-dev.yaml
NAME: devops-app-helm-dev
LAST DEPLOYED: Thu Apr  2 20:20:53 2026
NAMESPACE: default
STATUS: deployed
REVISION: 1
DESCRIPTION: Install complete
TEST SUITE: None
```

```bash
$ helm install devops-app-helm-prod . -f values-prod.yaml
NAME: devops-app-helm-prod
LAST DEPLOYED: Thu Apr  2 20:21:16 2026
NAMESPACE: default
STATUS: deployed
REVISION: 1
DESCRIPTION: Install complete
TEST SUITE: None
```

```bash
$ kubectl get endpoints
Warning: v1 Endpoints is deprecated in v1.33+; use discovery.k8s.io/v1 EndpointSlice
NAME                                       ENDPOINTS                                                        AGE
devops-app-helm-dev-devops-info-service    10.244.0.54:8000                                                 2m7s
devops-app-helm-devops-info-service        10.244.0.51:8000,10.244.0.52:8000,10.244.0.53:8000               9m13s
devops-app-helm-prod-devops-info-service   10.244.0.55:8000,10.244.0.56:8000,10.244.0.57:8000               105s
devops-info-service                        10.244.0.36:8000,10.244.0.39:8000,10.244.0.42:8000 + 2 more...   7d2h
devops-info-service-go                     10.244.0.38:8000,10.244.0.40:8000,10.244.0.41:8000 + 2 more...   6d22h
kubernetes                                 192.168.49.2:8443 
```

```bash
$ kubectl get pods -l "app.kubernetes.io/instance=devops-app-helm"
NAME                                                   READY   STATUS    RESTARTS   AGE
devops-app-helm-devops-info-service-5446cd687d-48jps   1/1     Running   0          10m
devops-app-helm-devops-info-service-5446cd687d-ls9fz   1/1     Running   0          10m
devops-app-helm-devops-info-service-5446cd687d-t7qtx   1/1     Running   0          10m
```

```bash
$ kubectl get deployments
NAME                                       READY   UP-TO-DATE   AVAILABLE   AGE
devops-app-helm-dev-devops-info-service    1/1     1            1           9m13s
devops-app-helm-devops-info-service        3/3     3            3           16m
devops-app-helm-prod-devops-info-service   3/3     3            3           8m51s
devops-info-service                        5/5     5            5           7d3h
devops-info-service-go                     5/5     5            5           6d22h
```

## Task 4 — Chart Hooks

```bash
$ helm install devops-app-helm-hook .
NAME: devops-app-helm-hook
LAST DEPLOYED: Thu Apr  2 20:58:34 2026
NAMESPACE: default
STATUS: deployed
REVISION: 1
DESCRIPTION: Install complete
TEST SUITE: None
```

```bash
$ kubectl get jobs -w
NAME                                                   STATUS    COMPLETIONS   DURATION   AGE
devops-app-helm-hook-devops-info-service-pre-install   Running   0/1           2s         2s
devops-app-helm-hook-devops-info-service-pre-install   Running   0/1           5s         5s
devops-app-helm-hook-devops-info-service-pre-install   Running   0/1           7s         7s
devops-app-helm-hook-devops-info-service-pre-install   SuccessCriteriaMet   0/1           8s         8s
devops-app-helm-hook-devops-info-service-pre-install   Complete             1/1           8s         8s
devops-app-helm-hook-devops-info-service-pre-install   Complete             1/1           8s         8s
devops-app-helm-hook-devops-info-service-post-install   Running              0/1                      0s
devops-app-helm-hook-devops-info-service-post-install   Running              0/1           0s         0s
devops-app-helm-hook-devops-info-service-post-install   Running              0/1           4s         4s
devops-app-helm-hook-devops-info-service-post-install   Running              0/1           14s        14s
devops-app-helm-hook-devops-info-service-post-install   SuccessCriteriaMet   0/1           15s        15s
devops-app-helm-hook-devops-info-service-post-install   Complete             1/1           15s        15s
devops-app-helm-hook-devops-info-service-post-install   Complete             1/1           15s        15s
```

```bash
$ kubectl get pods -w
NAME                                                         READY   STATUS    RESTARTS        AGE
devops-app-helm-dev-devops-info-service-7b87df56b7-9p24q     1/1     Running   0               37m
devops-app-helm-devops-info-service-5446cd687d-48jps         1/1     Running   0               44m
devops-app-helm-devops-info-service-5446cd687d-ls9fz         1/1     Running   0               44m
devops-app-helm-devops-info-service-5446cd687d-t7qtx         1/1     Running   0               44m
devops-app-helm-hook-devops-info-service-pre-install-rnzft   1/1     Running   0               4s
devops-app-helm-prod-devops-info-service-5894b88499-khfn5    1/1     Running   0               37m
devops-app-helm-prod-devops-info-service-5894b88499-prpxp    1/1     Running   0               37m
devops-app-helm-prod-devops-info-service-5894b88499-xtp2t    1/1     Running   0               37m
devops-info-service-78f44cdc7d-9lbts                         1/1     Running   1 (6d22h ago)   7d3h
devops-info-service-78f44cdc7d-gwbd9                         1/1     Running   1 (6d22h ago)   7d3h
devops-info-service-78f44cdc7d-kbfjx                         1/1     Running   1 (6d22h ago)   7d3h
devops-info-service-78f44cdc7d-ksxq4                         1/1     Running   1 (6d22h ago)   7d3h
devops-info-service-78f44cdc7d-rdx7w                         1/1     Running   1 (6d22h ago)   7d3h
devops-info-service-go-686d4f5dcd-9h44m                      1/1     Running   1 (6d22h ago)   6d22h
devops-info-service-go-686d4f5dcd-n2t2f                      1/1     Running   1 (6d22h ago)   6d22h
devops-info-service-go-686d4f5dcd-nxgbh                      1/1     Running   1 (6d22h ago)   6d22h
devops-info-service-go-686d4f5dcd-qd4nx                      1/1     Running   1 (6d22h ago)   6d22h
devops-info-service-go-686d4f5dcd-zf92x                      1/1     Running   1 (6d22h ago)   6d22h
devops-app-helm-hook-devops-info-service-pre-install-rnzft   0/1     Completed   0               6s
devops-app-helm-hook-devops-info-service-pre-install-rnzft   0/1     Completed   0               7s
devops-app-helm-hook-devops-info-service-pre-install-rnzft   0/1     Completed   0               8s
devops-app-helm-hook-devops-info-service-pre-install-rnzft   0/1     Completed   0               8s
devops-app-helm-hook-devops-info-service-pre-install-rnzft   0/1     Completed   0               8s
devops-app-helm-hook-devops-info-service-6dc67997fd-g24pn    0/1     Pending     0               0s
devops-app-helm-hook-devops-info-service-6dc67997fd-2j4cx    0/1     Pending     0               0s
devops-app-helm-hook-devops-info-service-6dc67997fd-blbq9    0/1     Pending     0               0s
devops-app-helm-hook-devops-info-service-6dc67997fd-g24pn    0/1     Pending     0               0s
devops-app-helm-hook-devops-info-service-post-install-thjtx   0/1     Pending     0               0s
devops-app-helm-hook-devops-info-service-6dc67997fd-blbq9     0/1     Pending     0               0s
devops-app-helm-hook-devops-info-service-6dc67997fd-2j4cx     0/1     Pending     0               0s
devops-app-helm-hook-devops-info-service-post-install-thjtx   0/1     Pending     0               0s
devops-app-helm-hook-devops-info-service-6dc67997fd-g24pn     0/1     ContainerCreating   0               0s
devops-app-helm-hook-devops-info-service-6dc67997fd-blbq9     0/1     ContainerCreating   0               0s
devops-app-helm-hook-devops-info-service-6dc67997fd-2j4cx     0/1     ContainerCreating   0               0s
devops-app-helm-hook-devops-info-service-post-install-thjtx   0/1     ContainerCreating   0               0s
devops-app-helm-hook-devops-info-service-6dc67997fd-blbq9     0/1     Running             0               1s
devops-app-helm-hook-devops-info-service-6dc67997fd-g24pn     0/1     Running             0               1s
devops-app-helm-hook-devops-info-service-6dc67997fd-2j4cx     0/1     Running             0               1s
devops-app-helm-hook-devops-info-service-post-install-thjtx   1/1     Running             0               3s
devops-app-helm-hook-devops-info-service-6dc67997fd-g24pn     1/1     Running             0               8s
devops-app-helm-hook-devops-info-service-6dc67997fd-blbq9     1/1     Running             0               8s
devops-app-helm-hook-devops-info-service-6dc67997fd-2j4cx     1/1     Running             0               8s
devops-app-helm-hook-devops-info-service-post-install-thjtx   0/1     Completed           0               13s
devops-app-helm-hook-devops-info-service-post-install-thjtx   0/1     Completed           0               14s
devops-app-helm-hook-devops-info-service-post-install-thjtx   0/1     Completed           0               15s
devops-app-helm-hook-devops-info-service-post-install-thjtx   0/1     Completed           0               15s
devops-app-helm-hook-devops-info-service-post-install-thjtx   0/1     Completed           0               15s
```

```bash
$ kubectl get jobs
No resources found in default namespace.
```

## Task 5 — Documentation

### Chart Overview

This Helm chart deploys the **DevOps Info Service** (Python FastAPI application) to Kubernetes. It includes:

- **Deployment** – Configurable replicas, image, resources, probes.
- **Service** – NodePort (default) or LoadBalancer.
- **Hooks** – Pre‑install validation and post‑install smoke test.

#### Key Template Files

| File | Purpose |
|------|---------|
| `deployment.yaml` | Defines the Pod template; uses `{{ .Values.replicaCount }}`, `{{ .Values.image }}`, `{{ .Values.resources }}`, and probes. |
| `service.yaml` | Exposes the app; uses `{{ .Values.service.type }}`, `{{ .Values.service.port }}`. |
| `_helpers.tpl` | Provides `fullname`, `labels`, `selectorLabels` for consistent naming and labeling. |

---

### Configuration Guide

#### Important Values

| Value | Description | Default |
|-------|-------------|---------|
| `replicaCount` | Number of Pod replicas | `2` |
| `image.repository` | Docker image repository | `yourusername/devops-info-service` |
| `image.tag` | Image tag | `latest` |
| `service.type` | Kubernetes service type (`NodePort` / `LoadBalancer`) | `NodePort` |
| `resources.limits.cpu` | CPU limit | `500m` |
| `resources.limits.memory` | Memory limit | `512Mi` |
| `livenessProbe.enabled` | Enable liveness probe | `true` |
| `readinessProbe.enabled` | Enable readiness probe | `true` |

#### Environment Overrides

**Development (`values-dev.yaml`):**
- 1 replica, `latest` image tag
- Relaxed resource limits (`100m` CPU, `128Mi` memory)
- `NodePort` service

**Production (`values-prod.yaml`):**
- 3 replicas, fixed image tag `1.0.0`
- Production‑grade resources (`500m` CPU, `512Mi` memory)
- `LoadBalancer` service (or `ClusterIP` with Ingress)

### Pre-install hook

- Purpose: Runs before any resources are created. In this chart, it validates that required values are set (e.g., someRequiredValue) and prints a confirmation. It uses a busybox container.

- Execution: Because of weight -5, it runs before any other pre‑install hooks. After successful completion, the job is deleted (policy hook-succeeded).

### Post-install hook

- Purpose: Runs after all resources are installed and ready. It performs a smoke test by repeatedly curling the application’s /health endpoint. If the endpoint becomes reachable within 5 attempts, the hook succeeds; otherwise it fails (causing the Helm installation to be marked as failed).

- Deletion: The job is removed after success, keeping the cluster clean.

### Ev.

```bash
$ helm list
NAME                    NAMESPACE       REVISION        UPDATED                                 STATUS          CHART                           APP VERSION
devops-app-helm         default         1               2026-04-02 20:13:48.313526916 +0300 MSK deployed        devops-info-service-0.1.0       1.0.0      
devops-app-helm-dev     default         1               2026-04-02 20:20:53.980521074 +0300 MSK deployed        devops-info-service-0.1.0       1.0.0      
devops-app-helm-hook    default         1               2026-04-02 20:58:34.4455097 +0300 MSK   deployed        devops-info-service-0.1.0       1.0.0      
devops-app-helm-prod    default         1               2026-04-02 20:21:16.583175547 +0300 MSK deployed        devops-info-service-0.1.0       1.0.0      
```

### Bonus: Library Chart

**Purpose:** Share common template logic across multiple applications.

**Chart:** `common-lib` (type: `library`)

**Provided helpers:**
- `common.name` – sanitized chart name
- `common.fullname` – release‑qualified name
- `common.labels` – standard Kubernetes labels
- `common.selectorLabels` – selector labels

**Used by:**
- `devops-info-service`
- `second-app` (example nginx app)

**Benefits:**
- Single source of truth for labeling and naming.
- Reduced duplication across charts.
- Easier updates: change once, `helm dependency update` everywhere.

```bash
$ helm dependency update
Hang tight while we grab the latest from your chart repositories...
...Unable to get an update from the "prometheus-community" chart repository (https://prometheus-community.github.io/helm-charts):
        Get "https://prometheus-community.github.io/helm-charts/index.yaml": context deadline exceeded (Client.Timeout exceeded while awaiting headers)
Update Complete. ⎈Happy Helming!⎈
Saving 1 charts
Deleting outdated charts
```

```bash
$ kubectl get deployments
NAME                                         READY   UP-TO-DATE   AVAILABLE   AGE
devops-app-go-helm-lib-devops-info-service   3/3     3            3           37s
devops-app-helm-dev-devops-info-service      1/1     1            1           99m
devops-app-helm-devops-info-service          3/3     3            3           106m
devops-app-helm-hook-devops-info-service     3/3     3            3           61m
devops-app-helm-lib-devops-info-service      3/3     3            3           9m42s
devops-app-helm-prod-devops-info-service     3/3     3            3           99m
devops-info-service                          5/5     5            5           7d5h
devops-info-service-go                       5/5     5            5           7d
```