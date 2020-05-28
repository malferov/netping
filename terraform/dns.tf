variable "web" {
  default = "127.0.0.1"
}

variable "api" {
  default = "127.0.0.1"
}

data "digitalocean_domain" "domain" {
  name = "${var.app}.${var.domain}"
}

resource "digitalocean_record" "web" {
  domain = data.digitalocean_domain.domain.name
  type   = "A"
  name   = "@"
  value  = var.web
  ttl    = 60
}

resource "digitalocean_record" "api" {
  domain = data.digitalocean_domain.domain.name
  type   = "A"
  name   = "api"
  value  = var.api
  ttl    = 60
}
