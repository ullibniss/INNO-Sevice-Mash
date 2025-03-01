resource "random_password" "master-lb-password" {
  length           = 40
  special          = false
}

resource "random_password" "coffee-password" {
  length           = 40
  special          = false
}

resource "random_password" "drink-password" {
  length           = 40
  special          = false
}

resource "random_password" "milk-password" {
  length           = 40
  special          = false
}

resource "random_password" "syrup-password" {
  length           = 40
  special          = false
}

locals {
    master-lb-pass = random_password.master-lb-password.result
    coffee-pass   = random_password.coffee-password.result
    drink-pass    = random_password.drink-password.result
    milk-pass     = random_password.milk-password.result
    syrup-pass    = random_password.syrup-password.result
}