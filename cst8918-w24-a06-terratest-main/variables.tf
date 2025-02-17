variable "labelPrefix" {
  type        = string
  default     = "sidh0148"
  description = "Prefix used to name the resources, derived from your college username."
}

variable "region" {
  type        = string
  default     = "EastUS"
  description = "The Azure region where all resources will be deployed."
}

variable "admin_username" {
  type        = string
  default     = "adminuser"
  description = "The administrative username for accessing the virtual machine."
}

variable "ssh_public_key_path" {
  type        = string
  default     = "~/.ssh/id_rsa.pub"
  description = "Path to the public key used for SSH access to the VM."
}

