output "permissions" {
  value = data.cloudflare_api_token_permission_groups.all
}

output "api_token" {
  value = cloudflare_api_token.this.value
}

