resource "twc_firewall" "coffee-firewall" {
  name = "Coffee firewall"
  description = "Coffee firewall"

  link {
    id = twc_server.coffee-server.id
    type = "server"
  }
}

resource "twc_firewall_rule" "https-firewall-rule" {
  firewall_id = twc_firewall.coffee-firewall.id
  description = "Coffee rule"

  direction = "ingress"
  port = 443
  protocol = "tcp"
  cidr = "0.0.0.0/0"
}

resource "twc_firewall_rule" "ssh-1901-firewall-rule" {
  firewall_id = twc_firewall.coffee-firewall.id
  description = "SSH 1901 rule"

  direction = "ingress"
  port = 1901
  protocol = "tcp"
  cidr = "0.0.0.0/0"
}

resource "twc_firewall_rule" "ssh-1902-firewall-rule" {
  firewall_id = twc_firewall.coffee-firewall.id
  description = "SSH 1902 rule"

  direction = "ingress"
  port = 1902
  protocol = "tcp"
  cidr = "0.0.0.0/0"
}

resource "twc_firewall_rule" "ssh-1903-firewall-rule" {
  firewall_id = twc_firewall.coffee-firewall.id
  description = "SSH 1903 rule"

  direction = "ingress"
  port = 1903
  protocol = "tcp"
  cidr = "0.0.0.0/0"
}

resource "twc_firewall_rule" "ssh-1904-firewall-rule" {
  firewall_id = twc_firewall.coffee-firewall.id
  description = "SSH 1904 rule"

  direction = "ingress"
  port = 1904
  protocol = "tcp"
  cidr = "0.0.0.0/0"
}

resource "twc_firewall_rule" "ssh-1905-firewall-rule" {
  firewall_id = twc_firewall.coffee-firewall.id
  description = "SSH 1905 rule"

  direction = "ingress"
  port = 1905
  protocol = "tcp"
  cidr = "0.0.0.0/0"
}