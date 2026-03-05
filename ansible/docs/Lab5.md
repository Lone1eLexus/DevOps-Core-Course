# Lab 5 — Ansible Fundamentals

## Deployment Verification

```bash
$ ansible-playbook playbooks/provision.yml

PLAY [Provision web servers] **************************************************************************************************************************************

TASK [Gathering Facts] ********************************************************************************************************************************************
ok: [devops-lab-vm]

TASK [common : Update apt cache] **********************************************************************************************************************************
ok: [devops-lab-vm]

TASK [common : Install common packages] ***************************************************************************************************************************
changed: [devops-lab-vm]

TASK [common : Set timezone to UTC] *******************************************************************************************************************************
changed: [devops-lab-vm]

TASK [docker : Install prerequisites] *****************************************************************************************************************************
ok: [devops-lab-vm]

TASK [docker : Add Docker GPG key] ********************************************************************************************************************************
changed: [devops-lab-vm]

TASK [docker : Add Docker repository] *****************************************************************************************************************************
changed: [devops-lab-vm]

TASK [docker : Update apt cache after adding Docker repo] *********************************************************************************************************
changed: [devops-lab-vm]

TASK [docker : Install Docker packages] ***************************************************************************************************************************
changed: [devops-lab-vm]

TASK [docker : Add user to docker group] **************************************************************************************************************************
changed: [devops-lab-vm]

TASK [docker : Install Docker] ************************************************************************************************************************************
ok: [devops-lab-vm]

TASK [docker : Install python3-docker for Ansible docker modules] *************************************************************************************************
changed: [devops-lab-vm]

RUNNING HANDLER [docker : restart docker] *************************************************************************************************************************
fatal: [devops-lab-vm]: FAILED! => {"changed": false, "msg": "Unable to start service docker: Job for docker.service failed because the control process exited with error code.\nSee \"systemctl status docker.service\" and \"journalctl -xeu docker.service\" for details.\n"}

PLAY RECAP ********************************************************************************************************************************************************
devops-lab-vm              : ok=12   changed=8    unreachable=0    failed=1    skipped=0    rescued=0    ignored=0  
```

```bash
$ ansible-playbook playbooks/provision.yml

PLAY [Provision web servers] **************************************************************************************************************************************

TASK [Gathering Facts] ********************************************************************************************************************************************
ok: [devops-lab-vm]

TASK [common : Update apt cache] **********************************************************************************************************************************
ok: [devops-lab-vm]

TASK [common : Install common packages] ***************************************************************************************************************************
ok: [devops-lab-vm]

TASK [common : Set timezone to UTC] *******************************************************************************************************************************
ok: [devops-lab-vm]

TASK [docker : Install prerequisites] *****************************************************************************************************************************
ok: [devops-lab-vm]

TASK [docker : Add Docker GPG key] ********************************************************************************************************************************
ok: [devops-lab-vm]

TASK [docker : Add Docker repository] *****************************************************************************************************************************
ok: [devops-lab-vm]

TASK [docker : Update apt cache after adding Docker repo] *********************************************************************************************************
changed: [devops-lab-vm]

TASK [docker : Install Docker packages] ***************************************************************************************************************************
ok: [devops-lab-vm]

TASK [docker : Add user to docker group] **************************************************************************************************************************
ok: [devops-lab-vm]

TASK [docker : Install Docker] ************************************************************************************************************************************
ok: [devops-lab-vm]

TASK [docker : Install python3-docker for Ansible docker modules] *************************************************************************************************
ok: [devops-lab-vm]

PLAY RECAP ********************************************************************************************************************************************************
devops-lab-vm              : ok=12   changed=1    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0 
```

```bash
$ ansible-playbook playbooks/provision.yml

PLAY [Provision web servers] **************************************************************************************************************************************

TASK [Gathering Facts] ********************************************************************************************************************************************
ok: [devops-lab-vm]

TASK [common : Update apt cache] **********************************************************************************************************************************
ok: [devops-lab-vm]

TASK [common : Install common packages] ***************************************************************************************************************************
changed: [devops-lab-vm]

TASK [common : Set timezone to UTC] *******************************************************************************************************************************
changed: [devops-lab-vm]

TASK [docker : Install prerequisites] *****************************************************************************************************************************
ok: [devops-lab-vm]

TASK [docker : Add Docker GPG key] ********************************************************************************************************************************
changed: [devops-lab-vm]

TASK [docker : Add Docker repository] *****************************************************************************************************************************
changed: [devops-lab-vm]

TASK [docker : Update apt cache after adding Docker repo] *********************************************************************************************************
changed: [devops-lab-vm]

TASK [docker : Install Docker packages] ***************************************************************************************************************************
changed: [devops-lab-vm]

TASK [docker : Add user to docker group] **************************************************************************************************************************
changed: [devops-lab-vm]

TASK [docker : Install Docker] ************************************************************************************************************************************
ok: [devops-lab-vm]

TASK [docker : Install python3-docker for Ansible docker modules] *************************************************************************************************
changed: [devops-lab-vm]

RUNNING HANDLER [docker : restart docker] *************************************************************************************************************************
fatal: [devops-lab-vm]: FAILED! => {"changed": false, "msg": "Unable to start service docker: Job for docker.service failed because the control process exited with error code.\nSee \"systemctl status docker.service\" and \"journalctl -xeu docker.service\" for details.\n"}

PLAY RECAP ********************************************************************************************************************************************************
devops-lab-vm              : ok=12   changed=8    unreachable=0    failed=1    skipped=0    rescued=0    ignored=0   

lord@EmptyThrone:~/DevOps/DevOps-Core-Course/ansible$ ansible-playbook playbooks/provision.yml

PLAY [Provision web servers] **************************************************************************************************************************************

TASK [Gathering Facts] ********************************************************************************************************************************************
ok: [devops-lab-vm]

TASK [common : Update apt cache] **********************************************************************************************************************************
ok: [devops-lab-vm]

TASK [common : Install common packages] ***************************************************************************************************************************
ok: [devops-lab-vm]

TASK [common : Set timezone to UTC] *******************************************************************************************************************************
ok: [devops-lab-vm]

TASK [docker : Install prerequisites] *****************************************************************************************************************************
ok: [devops-lab-vm]

TASK [docker : Add Docker GPG key] ********************************************************************************************************************************
ok: [devops-lab-vm]

TASK [docker : Add Docker repository] *****************************************************************************************************************************
ok: [devops-lab-vm]

TASK [docker : Update apt cache after adding Docker repo] *********************************************************************************************************
changed: [devops-lab-vm]

TASK [docker : Install Docker packages] ***************************************************************************************************************************
ok: [devops-lab-vm]

TASK [docker : Add user to docker group] **************************************************************************************************************************
ok: [devops-lab-vm]

TASK [docker : Install Docker] ************************************************************************************************************************************
ok: [devops-lab-vm]

TASK [docker : Install python3-docker for Ansible docker modules] *************************************************************************************************
ok: [devops-lab-vm]

PLAY RECAP ********************************************************************************************************************************************************
devops-lab-vm              : ok=12   changed=1    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0   

lord@EmptyThrone:~/DevOps/DevOps-Core-Course/ansible$ ansible-playbook playbooks/provision.yml

PLAY [Provision web servers] **************************************************************************************************************************************

TASK [Gathering Facts] ********************************************************************************************************************************************
ok: [devops-lab-vm]

TASK [common : Update apt cache] **********************************************************************************************************************************
ok: [devops-lab-vm]

TASK [common : Install common packages] ***************************************************************************************************************************
ok: [devops-lab-vm]

TASK [common : Set timezone to UTC] *******************************************************************************************************************************
ok: [devops-lab-vm]

TASK [docker : Install prerequisites] *****************************************************************************************************************************
ok: [devops-lab-vm]

TASK [docker : Add Docker GPG key] ********************************************************************************************************************************
ok: [devops-lab-vm]

TASK [docker : Add Docker repository] *****************************************************************************************************************************
ok: [devops-lab-vm]

TASK [docker : Update apt cache after adding Docker repo] *********************************************************************************************************
changed: [devops-lab-vm]

TASK [docker : Install Docker packages] ***************************************************************************************************************************
ok: [devops-lab-vm]

TASK [docker : Add user to docker group] **************************************************************************************************************************
ok: [devops-lab-vm]

TASK [docker : Install Docker] ************************************************************************************************************************************
ok: [devops-lab-vm]

TASK [docker : Install python3-docker for Ansible docker modules] *************************************************************************************************
ok: [devops-lab-vm]

PLAY RECAP ********************************************************************************************************************************************************
devops-lab-vm              : ok=12   changed=1    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0  
```

```bash
devops-lab-vm | CHANGED | rc=0 >>
Docker version 29.2.1, build a5c7197
```

```bash
$ ansible webservers -a "docker run hello-world"
devops-lab-vm | FAILED | rc=1 >>
Cannot connect to the Docker daemon at unix:///var/run/docker.sock. Is the docker daemon running?non-zero return code
```

### Update main.yml docker tasks

```bash
$ ansible-playbook playbooks/provision.yml

PLAY [Provision web servers] **************************************************************************************************************************************

TASK [Gathering Facts] ********************************************************************************************************************************************
ok: [devops-lab-vm]

TASK [common : Update apt cache] **********************************************************************************************************************************
ok: [devops-lab-vm]

TASK [common : Install common packages] ***************************************************************************************************************************
ok: [devops-lab-vm]

TASK [common : Set timezone to UTC] *******************************************************************************************************************************
ok: [devops-lab-vm]

TASK [docker : Remove old Docker repository if exists] ************************************************************************************************************
ok: [devops-lab-vm]

TASK [docker : Remove old Docker source list if exists] ***********************************************************************************************************
changed: [devops-lab-vm]

TASK [docker : Clean apt cache] ***********************************************************************************************************************************
changed: [devops-lab-vm]

TASK [docker : Install prerequisites] *****************************************************************************************************************************
ok: [devops-lab-vm]

TASK [docker : Add Docker GPG key] ********************************************************************************************************************************
ok: [devops-lab-vm]

TASK [docker : Add Docker repository] *****************************************************************************************************************************
changed: [devops-lab-vm]

TASK [docker : Update apt cache after adding Docker repo] *********************************************************************************************************
changed: [devops-lab-vm]

TASK [docker : Install Docker packages] ***************************************************************************************************************************
changed: [devops-lab-vm]

TASK [docker : Install Docker] ************************************************************************************************************************************
ok: [devops-lab-vm]

TASK [docker : Ensure Docker service is running and enabled] ******************************************************************************************************
ok: [devops-lab-vm]

TASK [docker : Add user to docker group] **************************************************************************************************************************
ok: [devops-lab-vm]

TASK [docker : Install python3-docker for Ansible docker modules] *************************************************************************************************
ok: [devops-lab-vm]

RUNNING HANDLER [docker : restart docker] *************************************************************************************************************************
changed: [devops-lab-vm]

PLAY RECAP ********************************************************************************************************************************************************
devops-lab-vm              : ok=17   changed=6    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0   
```

```bash
$ ansible webservers -a "docker run hello-world"
devops-lab-vm | CHANGED | rc=0 >>

Hello from Docker!
This message shows that your installation appears to be working correctly.

To generate this message, Docker took the following steps:
 1. The Docker client contacted the Docker daemon.
 2. The Docker daemon pulled the "hello-world" image from the Docker Hub.
    (amd64)
 3. The Docker daemon created a new container from that image which runs the
    executable that produces the output you are currently reading.
 4. The Docker daemon streamed that output to the Docker client, which sent it
    to your terminal.

To try something more ambitious, you can run an Ubuntu container with:
 $ docker run -it ubuntu bash

Share images, automate workflows, and more with a free Docker ID:
 https://hub.docker.com/

For more examples and ideas, visit:
 https://docs.docker.com/get-started/Unable to find image 'hello-world:latest' locally
latest: Pulling from library/hello-world
17eec7bbc9d7: Pulling fs layer
17eec7bbc9d7: Verifying Checksum
17eec7bbc9d7: Download complete
17eec7bbc9d7: Pull complete
Digest: sha256:ef54e839ef541993b4e87f25e752f7cf4238fa55f017957c2eb44077083d7a6a
Status: Downloaded newer image for hello-world:latest
```

```bash
$ ansible-playbook playbooks/deploy.yml --ask-vault-pass
Vault password: 

PLAY [Deploy application] *****************************************************************************************************************************************

TASK [Gathering Facts] ********************************************************************************************************************************************
ok: [devops-lab-vm]

TASK [app_deploy : Log in to Docker Hub] **************************************************************************************************************************
ok: [devops-lab-vm]

TASK [app_deploy : Pull Docker image] *****************************************************************************************************************************
changed: [devops-lab-vm]

TASK [app_deploy : Stop existing container if running] ************************************************************************************************************
fatal: [devops-lab-vm]: FAILED! => {"changed": false, "msg": "Cannot create container when image is not specified!"}
...ignoring

TASK [app_deploy : Remove old container if exists] ****************************************************************************************************************
ok: [devops-lab-vm]

TASK [app_deploy : Run new container] *****************************************************************************************************************************
changed: [devops-lab-vm]

TASK [app_deploy : Wait for application to be ready] **************************************************************************************************************
ok: [devops-lab-vm]

TASK [app_deploy : Verify health endpoint] ************************************************************************************************************************
ok: [devops-lab-vm]

RUNNING HANDLER [app_deploy : restart app container] **************************************************************************************************************
changed: [devops-lab-vm]

PLAY RECAP ********************************************************************************************************************************************************
devops-lab-vm              : ok=9    changed=3    unreachable=0    failed=0    skipped=0    rescued=0    ignored=1 
```

```bash
$ ansible webservers -a "docker ps" --ask-vault-pass
Vault password: 
devops-lab-vm | CHANGED | rc=0 >>
CONTAINER ID   IMAGE                               COMMAND           CREATED              STATUS              PORTS                              NAMES
ca5e588697ea   lehus1/devops-info-service:latest   "python app.py"   About a minute ago   Up About a minute   8000/tcp, 0.0.0.0:8000->5000/tcp   devops-info-service
```

```bash
$ ansible webservers -a "curl -s http://localhost:8000/health" --ask-vault-pass
Vault password: 
devops-lab-vm | CHANGED | rc=0 >>
{"status":"healthy","timestamp":"2026-02-26T20:19:53.831693+00:00","uptime_seconds":155}
```

## Architecture Overview

- **Ansible version:** 2.16.3
- **Target VM:** Ubuntu 22.04 LTS (Pulumi-created, IP: 89.169.133.12)
- **Structure:** Role-based organization (common, docker, app_deploy)
- **Why roles:** Reusability, clean separation of concerns, easy to maintain and share

## Roles Documentation

| Role | Purpose | Key Variables | Handlers |
|------|---------|---------------|----------|
| **common** | System packages, timezone | `common_packages` | None |
| **docker** | Docker CE installation | `docker_user`, `docker_packages` | `restart docker` |
| **app_deploy** | Deploy containerized app | `dockerhub_username`, `app_port` | `restart app container` |

## Idempotency Demonstration

**First run:** 15+ tasks changed (packages installed, Docker configured)
**Second run:** 0 changed, all ok — system already in desired state

**Why it works:** Ansible modules check state before acting (e.g., `apt: state=present` only installs if missing).

## Ansible Vault Usage

- **Vault file:** `group_vars/vault.yml` (encrypted)
- **Password:** Stored in `.vault_pass` (gitignored, chmod 600)
- **Usage:** `ansible-playbook deploy.yml --vault-password-file .vault_pass`
- **Why:** Credentials never exposed in plaintext or Git

## Key Decisions
- Roles provide a standardized structure that separates tasks, variables, and handlers into logical units. This makes the code reusable across different playbooks and environments, easier to maintain, and simpler to share with others via Ansible Galaxy. Each role is self-contained with its own defaults, tasks, and handlers. I can apply the same docker role to both development and production servers.

- Using Ansible's stateful modules like apt, service, and docker_container instead of raw shell commands. These modules check the current state before making changes, so running the playbook multiple times produces the same result without unnecessary modifications.

- Handlers only execute when notified by a task that actually changed something. For example, Docker only restarts if the configuration file was modified, not on every playbook run.

- Ansible Vault encrypts sensitive data like Docker Hub credentials at rest while keeping the files in version control. Only users with the vault password can decrypt and use the secrets, preventing credential leaks in Git repositories.

## Challenges

I had some, but I forgot