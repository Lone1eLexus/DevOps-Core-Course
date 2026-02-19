terraform {
  required_providers {
    yandex = {
      source  = "yandex-cloud/yandex"
      version = "~> 0.98"
    }
  }
}

provider "yandex" {
  cloud_id  = var.cloud_id
  folder_id = var.folder_id
  zone      = var.zone
}

data "yandex_vpc_subnet" "existing" {
  subnet_id = var.subnet_id
}

resource "yandex_vpc_security_group" "devops_sg" {
  name       = "devops-lab-sg"
  network_id = data.yandex_vpc_subnet.existing.network_id

  ingress {
    protocol       = "TCP"
    port           = 22
    v4_cidr_blocks = ["0.0.0.0/0"]
    description    = "SSH access"
  }

  ingress {
    protocol       = "TCP"
    port           = 80
    v4_cidr_blocks = ["0.0.0.0/0"]
    description    = "HTTP access"
  }

  ingress {
    protocol       = "TCP"
    port           = 5000
    v4_cidr_blocks = ["0.0.0.0/0"]
    description    = "Custom app port"
  }

  egress {
    protocol       = "ANY"
    from_port      = 0
    to_port        = 65535
    v4_cidr_blocks = ["0.0.0.0/0"]
    description    = "All outbound"
  }

  labels = {
    environment = "lab"
    lab         = "lab04"
    managed-by  = "terraform"
  }
}

data "yandex_compute_image" "ubuntu" {
  family = "ubuntu-2204-lts"
}

resource "yandex_compute_instance" "devops_vm" {
  name        = var.instance_name
  platform_id = var.instance_platform
  zone        = var.zone

  resources {
    cores         = var.instance_cores
    core_fraction = 20
    memory        = var.instance_memory
  }

  boot_disk {
    initialize_params {
      image_id = data.yandex_compute_image.ubuntu.id
      type     = "network-hdd"
      size     = 10
    }
  }

  network_interface {
    subnet_id          = var.subnet_id
    nat                = true
    security_group_ids = [yandex_vpc_security_group.devops_sg.id]
  }

  metadata = {
    ssh-keys  = "ubuntu:${var.ssh_public_key}"
    user-data = <<-EOF
      #cloud-config
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
      EOF
  }

  labels = {
    environment = "lab"
    lab         = "lab04"
    managed-by  = "terraform"
  }
}