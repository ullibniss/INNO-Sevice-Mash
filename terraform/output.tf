output "master-lb-password" {
    value     = local.master-lb-pass
    sensitive = true
}

output "coffee-password" {
    value = local.coffee-pass
    sensitive = true
}

output "drink-password" {
    value = local.drink-pass
    sensitive = true
}

output "milk-password" {
    value     = local.milk-pass
    sensitive = true
}

output "syrup-password" {
    value = local.syrup-pass
    sensitive = true
}