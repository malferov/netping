variable "ip" {
  default = "127.0.0.1"
}

data "digitalocean_domain" "domain" {
  name = "${var.app}.${var.domain}"
}

resource "digitalocean_record" "app" {
  domain = data.digitalocean_domain.domain.name
  type   = "A"
  name   = "@"
  value  = var.ip
  ttl    = 60
}
