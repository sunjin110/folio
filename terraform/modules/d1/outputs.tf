output "api_token" {
  value = cloudflare_api_token.this.value
}

output "db_id" {
  value = cloudflare_d1_database.this.id
}
