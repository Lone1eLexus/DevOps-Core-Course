# Lab 11 — Kubernetes Secrets & HashiCorp Vault

## Task 1 — Kubernetes Secrets Fundamentals

```bash
$ kubectl create secret generic app-credentials   --from-literal=username=Stranger   --from-literal=password=Funny
secret/app-credentials created
```

```bash
$ kubectl get secret app-credentials -o yaml
apiVersion: v1
data:
  password: RnVubnk=
  username: U3RyYW5nZXI=
kind: Secret
metadata:
  creationTimestamp: "2026-04-09T17:02:52Z"
  name: app-credentials
  namespace: default
  resourceVersion: "30894"
  uid: 8ab9bd68-6096-4edb-9c5e-01b2da0a85db
type: Opaque
```

```bash
$ echo "U3RyYW5nZXI=" | base64 -d
Stranger
```

```bash
$ echo "RnVubnk=" | base64 -d
Funny
```

## Task 2 — Helm-Managed Secrets

```bash
$ helm upgrade devops-app-helm .   --set secrets.username=Stranger   --set secrets.password=Funny   --set image.tag=latest
Release "devops-app-helm" has been upgraded. Happy Helming!
NAME: devops-app-helm
LAST DEPLOYED: Thu Apr  9 20:28:00 2026
NAMESPACE: default
STATUS: deployed
REVISION: 2
DESCRIPTION: Upgrade complete
TEST SUITE: None
```

```bash
$ kubectl get secret devops-app-helm-secret -o yaml
apiVersion: v1
data:
  password: RnVubnk=
  username: U3RyYW5nZXI=
kind: Secret
metadata:
  annotations:
    meta.helm.sh/release-name: devops-app-helm
    meta.helm.sh/release-namespace: default
  creationTimestamp: "2026-04-09T17:28:00Z"
  labels:
    app.kubernetes.io/instance: devops-app-helm
    app.kubernetes.io/managed-by: Helm
  name: devops-app-helm-secret
  namespace: default
  resourceVersion: "32560"
  uid: 84a31224-e080-4102-be9c-dd55ebad4018
type: Opaque
```

```bash
$ kubectl exec -it devops-app-helm-devops-info-service-785f6b4684-2z4sl -- /bin/sh -c 'echo $username; echo $password'
Stranger
Funny
```

## Task 3 — HashiCorp Vault Integration

```bash
$ git clone https://github.com/hashicorp/vault-helm.git
Cloning into 'vault-helm'...
remote: Enumerating objects: 4961, done.
remote: Counting objects: 100% (168/168), done.
remote: Compressing objects: 100% (115/115), done.
remote: Total 4961 (delta 125), reused 57 (delta 53), pack-reused 4793 (from 3)
Receiving objects: 100% (4961/4961), 1.32 MiB | 2.09 MiB/s, done.
Resolving deltas: 100% (3684/3684), done.
```

```bash
$ cd vault-helm
```

```bash
$ helm install vault . --set "server.dev.enabled=true" --set "injector.enabled=true"
NAME: vault
LAST DEPLOYED: Thu Apr  9 20:37:37 2026
NAMESPACE: default
STATUS: deployed
REVISION: 1
DESCRIPTION: Install complete
NOTES:
Thank you for installing HashiCorp Vault!

Now that you have deployed Vault, you should look over the docs on using
Vault with Kubernetes available here:

https://developer.hashicorp.com/vault/docs


Your release is named vault. To learn more about the release, try:

  $ helm status vault
  $ helm get manifest vault
```

```bash
$ kubectl exec -it vault-0 -- /bin/sh
/ $ vault secrets enable -path=secret kv-v2
Error enabling: Error making API request.

URL: POST http://127.0.0.1:8200/v1/sys/mounts/secret
Code: 400. Errors:

* path is already in use at secret/
/ $ vault kv put secret/devops-app username="Stranger" password="Funny"
===== Secret Path =====
secret/data/devops-app

======= Metadata =======
Key                Value
---                -----
created_time       2026-04-09T17:48:56.446946587Z
custom_metadata    <nil>
deletion_time      n/a
destroyed          false
version            1
/ $ vault kv get secret/devops-app
===== Secret Path =====
secret/data/devops-app

======= Metadata =======
Key                Value
---                -----
created_time       2026-04-09T17:48:56.446946587Z
custom_metadata    <nil>
deletion_time      n/a
destroyed          false
version            1

====== Data ======
Key         Value
---         -----
password    Funny
username    Stranger
/ $ vault auth enable kubernetes
Success! Enabled kubernetes auth method at: kubernetes/
cat /var/run/secrets/kubernetes.io/serviceaccount/token
...
/ $ vault write auth/kubernetes/config \
  kubernetes_host="https://10.96.0.1:443" \
  token_reviewer_jwt="...
/ $ vault policy write devops-app-policy - <<EOF
> path "secret/data/devops-app" {
>   capabilities = ["read"]
> }
> EOF
Success! Uploaded policy: devops-app-policy
/ $ vault write auth/kubernetes/role/devops-app-role \
>   bound_service_account_names=default \
>   bound_service_account_namespaces=default \
>   policies=devops-app-policy \
>   ttl=24h
WARNING! The following warnings were returned from Vault:

  * Role devops-app-role does not have an audience configured. While audiences
  are not required, consider specifying one if your use case would benefit
  from additional JWT claim verification.
```

```bash
$ helm upgrade devops-app-helm . --set image.tag=latest
Release "devops-app-helm" has been upgraded. Happy Helming!
NAME: devops-app-helm
LAST DEPLOYED: Thu Apr  9 21:12:52 2026
NAMESPACE: default
STATUS: deployed
REVISION: 8
DESCRIPTION: Upgrade complete
TEST SUITE: None
```

```bash
$ kubectl logs devops-app-helm-devops-info-service-5d66db4cbb-pkzpd -c vault-agent
==> Vault Agent started! Log data will stream in below:

==> Vault Agent configuration:

2026-04-09T18:12:59.819Z [INFO]  agent.sink.file: creating file sink
2026-04-09T18:12:59.819Z [INFO]  agent.sink.file: file sink configured: path=/home/vault/.vault-token mode=-rw-r----- owner=100 group=1000
2026-04-09T18:12:59.819Z [INFO]  agent.exec.server: starting exec server
2026-04-09T18:12:59.819Z [INFO]  agent.exec.server: no env templates or exec config, exiting
           Api Address 1: http://bufconn
                     Cgo: disabled
               Log Level: info
                 Version: Vault v1.21.2, built 2026-01-06T08:33:05Z
             Version Sha: 781ba452d731fe2d59ccbc1b37ca7c5a18edb998

2026-04-09T18:12:59.819Z [INFO]  agent.sink.server: starting sink server
2026-04-09T18:12:59.819Z [INFO]  agent.auth.handler: starting auth handler
2026-04-09T18:12:59.819Z [INFO]  agent.template.server: starting template server
2026-04-09T18:12:59.819Z [INFO]  agent.auth.handler: authenticating
2026-04-09T18:12:59.819Z [INFO]  agent: (runner) creating new runner (dry: false, once: false)
2026-04-09T18:12:59.820Z [INFO]  agent: (runner) creating watcher
2026-04-09T18:12:59.821Z [INFO]  agent.auth.handler: authentication successful, sending token to sinks
2026-04-09T18:12:59.821Z [INFO]  agent.auth.handler: starting renewal process
2026-04-09T18:12:59.821Z [INFO]  agent.sink.file: token written: path=/home/vault/.vault-token
2026-04-09T18:12:59.821Z [INFO]  agent.template.server: template server received new token
2026-04-09T18:12:59.821Z [INFO]  agent: (runner) stopping
2026-04-09T18:12:59.821Z [INFO]  agent: (runner) creating new runner (dry: false, once: false)
2026-04-09T18:12:59.821Z [INFO]  agent: (runner) creating watcher
2026-04-09T18:12:59.821Z [INFO]  agent: (runner) starting
2026-04-09T18:12:59.822Z [INFO]  agent.auth.handler: renewed auth token
```

```bash
$ kubectl exec -it devops-app-helm-devops-info-service-5d66db4cbb-pkzpd -- /bin/sh
Defaulted container "devops-info-service" out of: devops-info-service, vault-agent, vault-agent-init (init)
$ cat /vault/secrets/config
data: map[password:Funny username:Stranger]
metadata: map[created_time:2026-04-09T17:48:56.446946587Z custom_metadata:<nil> deletion_time: destroyed:false version:1]
```

## Task 4 — Documentation

Kubernetes Secrets are only base64‑encoded, not encrypted by default.

### Security Analysis

| Aspect | Kubernetes Secrets | HashiCorp Vault |
|--------|-------------------|-----------------|
| **Encryption at rest** | Optional (etcd encryption) | Always encrypted |
| **Access control** | RBAC | Fine‑grained policies + auth methods |
| **Audit logging** | Limited | Detailed audit logs |
| **Dynamic secrets** | No | Yes (e.g., database credentials) |
| **Secret rotation** | Manual | Automated with leases |

### Production recommendations

- Use **Kubernetes Secrets** for non‑sensitive configs or when simplicity is needed.
- Use **Vault** for production secrets, especially when audit, rotation, or multi‑cloud is required.
- Always enable **etcd encryption** if using K8s Secrets for sensitive data.
- **Never commit real secrets to Git** – use `--set` or external secret managers.

## Bonus task

```bash
$ kubectl exec -it devops-app-helm-devops-info-service-767947b8-22r8n  -- cat /vault/secrets/config
Defaulted container "devops-info-service" out of: devops-info-service, vault-agent, vault-agent-init (init)
USERNAME=Stranger
PASSWORD=Funny
```

### Dynamic Secret Rotation with Vault Agent

Vault Agent's templating engine automatically renews secrets when their leases expire. The sidecar container runs continuously, authenticates with Vault, and monitors the lease lifecycle.

- When a secret's TTL expires, the Agent requests a new lease and re-renders the template.
- The updated secret is written to the shared volume, and the application can detect and reload it (e.g., using inotify or periodic checks).
- This eliminates the need for pod restarts or manual secret rotation.

```bash
$ kubectl exec -it devops-app-helm-devops-info-service-76b89c8c-lfkqr   -- printenv | grep -E "APP_ENV|LOG_LEVEL"
Defaulted container "devops-info-service" out of: devops-info-service, vault-agent, vault-agent-init (init)
APP_ENV=production
LOG_LEVEL=warn
```


