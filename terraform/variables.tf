variable "cloud_id" {
  description = "Yandex Cloud ID"
  type        = string
}

variable "folder_id" {
  description = "Yandex Cloud folder ID"
  type        = string
}

variable "zone" {
  description = "Availability zone"
  type        = string
  default     = "ru-central1-a"
}

variable "subnet_id" {
  description = "Subnet ID for VM placement"
  type        = string
}

variable "ssh_public_key" {
  description = "SSH public key for VM access"
  type        = string
}

variable "instance_name" {
  description = "VM instance name"
  type        = string
  default     = "devops-lab-vm"
}

variable "instance_cores" {
  description = "Number of CPU cores"
  type        = number
  default     = 2
}

variable "instance_memory" {
  description = "RAM in GB"
  type        = number
  default     = 1
}

variable "instance_platform" {
  description = "Instance platform"
  type        = string
  default     = "standard-v2"
}
