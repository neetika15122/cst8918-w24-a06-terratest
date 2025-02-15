# Define config variables
variable "labelPrefix" {
  type        = string
  description = "My college username"
  default = "pras0044"
}

variable "region" {
  default = "westus3"
}

variable "admin_username" {
  type        = string
  default     = "azureadmin"
  description = "The username for the local user account on the VM."
}
