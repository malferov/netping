terraform {
  required_version = ">= 0.12"
  backend "remote" {}
}

variable "letsencrypt_url" {}

provider "acme" {
  server_url = var.letsencrypt_url
}
