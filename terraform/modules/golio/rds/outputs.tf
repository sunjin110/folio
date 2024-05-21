output "uri" {
  value = "postgres://${local.user}:${random_password.user_password.result}@${aws_db_proxy_endpoint.this.endpoint}:5432/${local.db_name}"
}
