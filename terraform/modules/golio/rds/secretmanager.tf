resource "aws_secretsmanager_secret" "this" {
  name = "${var.prefix}-golio-db-credentails"
  #   TODO kms
}

resource "aws_secretsmanager_secret_version" "this" {
  secret_id = aws_secretsmanager_secret.this.id
  secret_string = jsonencode({
    username : "golio",
    password = random_password.user_password.result
  })
}
