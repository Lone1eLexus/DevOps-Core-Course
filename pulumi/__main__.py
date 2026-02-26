import pulumi
import pulumi_yandex as yandex
import os


config = pulumi.Config()

folder_id = os.getenv("YC_FOLDER_ID") or config.require("folder_id")
zone = os.getenv("YC_ZONE") or config.get("zone") or "ru-central1-a"
cloud_id = os.getenv("YC_CLOUD_ID") or config.require("cloud_id")

# SSH public key
ssh_public_key_path = os.path.expanduser("~/.ssh/id_rsa.pub")
with open(ssh_public_key_path) as f:
    ssh_public_key = f.read().strip()

# --- VPC Network ---
network = yandex.VpcNetwork("lab-network")

# --- Subnet ---
subnet = yandex.VpcSubnet("lab-subnet",
    zone=zone,
    network_id=network.id,
    v4_cidr_blocks=[subnet_cidr])

# --- Security Group ---
security_group = yandex.VpcSecurityGroup("lab-sg",
    network_id=network.id,
    ingress=[
        yandex.VpcSecurityGroupIngressArgs(
            protocol="TCP",
            port=22,
            v4_cidr_blocks=["0.0.0.0/0"],
            description="SSH",
        ),
        yandex.VpcSecurityGroupIngressArgs(
            protocol="TCP",
            port=80,
            v4_cidr_blocks=["0.0.0.0/0"],
            description="HTTP",
        ),
        yandex.VpcSecurityGroupIngressArgs(
            protocol="TCP",
            port=5000,
            v4_cidr_blocks=["0.0.0.0/0"],
            description="Custom app port",
        ),
    ],
    egress=[yandex.VpcSecurityGroupEgressArgs(
        protocol="ANY",
        from_port=0,
        to_port=65535,
        v4_cidr_blocks=["0.0.0.0/0"],
        description="All outbound",
    )])

# --- VM Image ---
# Get the latest Ubuntu 22.04 LTS image
ubuntu_image = yandex.get_compute_image(family="ubuntu-2204-lts")

# --- VM Instance ---
vm = yandex.ComputeInstance("devops-vm",
    name="devops-lab-vm",
    platform_id="standard-v2",
    zone=zone,
    resources=yandex.ComputeInstanceResourcesArgs(
        cores=2,
        memory=1,
        core_fraction=20,
    ),
    boot_disk=yandex.ComputeInstanceBootDiskArgs(
        initialize_params=yandex.ComputeInstanceBootDiskInitializeParamsArgs(
            image_id=ubuntu_image.id,
            size=10,
            type="network-hdd",
        )
    ),
    network_interfaces=[yandex.ComputeInstanceNetworkInterfaceArgs(
        subnet_id=subnet.id,
        nat=True,
        security_group_ids=[security_group.id],
    )],
    metadata={
        "ssh-keys": f"ubuntu:{ssh_public_key}",
        "user-data": """#cloud-config
package_update: true
package_upgrade: true
packages:
  - curl
  - vim
  - git
  - docker.io
runcmd:
  - [ systemctl, enable, --now, ssh ]
  - [ systemctl, enable, --now, docker ]
"""
    })

# --- Outputs ---
pulumi.export("vm_public_ip", vm.network_interfaces[0].nat_ip_address)
pulumi.export("ssh_command", pulumi.Output.concat("ssh ubuntu@", vm.network_interfaces[0].nat_ip_address))
