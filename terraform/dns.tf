variable "ip" {}

data "digitalocean_domain" "domain" {
  name = "${var.app}.${var.domain}"
}

resource "digitalocean_record" "web" {
  domain = data.digitalocean_domain.domain.name
  type   = "A"
  name   = "@"
  value  = var.ip
  ttl    = 60
}

resource "digitalocean_record" "api" {
  domain = data.digitalocean_domain.domain.name
  type   = "A"
  name   = "api"
  value  = var.ip
  ttl    = 60
}
