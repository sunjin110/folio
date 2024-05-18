resource "aws_db_proxy" "this" {
  name                   = "${var.prefix}-golio-db-proxy"
  debug_logging          = false
  engine_family          = "POSTGRESQL"
  idle_client_timeout    = 300
  require_tls            = true
  role_arn               = aws_iam_role.rds_proxy_role.arn
  vpc_security_group_ids = [aws_security_group.this.id]
  vpc_subnet_ids         = var.network.private_subnet_ids

  auth {
    auth_scheme = "SECRETS"
    iam_auth    = "DISABLED"
    secret_arn  = aws_secretsmanager_secret.this.arn
  }
  tags = {
    Name = "${var.prefix}-golio-db-proxy"
  }
}

resource "aws_db_proxy_default_target_group" "this" {
  db_proxy_name = aws_db_proxy.this.name
  connection_pool_config {
    connection_borrow_timeout = 120
    max_connections_percent   = 100
  }
}

resource "aws_db_proxy_target" "this" {
  db_cluster_identifier = aws_rds_cluster.this.id
  db_proxy_name         = aws_db_proxy.this.name
  target_group_name     = aws_db_proxy_default_target_group.this.name
}

resource "aws_db_proxy_endpoint" "this" {
  db_proxy_name          = aws_db_proxy.this.name
  db_proxy_endpoint_name = "${var.prefix}-golio-db-endpoint"
  vpc_security_group_ids = [aws_security_group.this.id]
  vpc_subnet_ids         = var.network.private_subnet_ids
  target_role            = "READ_WRITE"
}

resource "aws_iam_role" "rds_proxy_role" {
  name = "${var.prefix}-golio-rds-proxy-role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [{
      Effect = "Allow"
      Principal = {
        Service = "rds.amazonaws.com"
      }
      Action = "sts:AssumeRole"
    }]
  })
}

resource "aws_iam_role_policy_attachment" "rds_proxy_role_secret_manager" {
  role       = aws_iam_role.rds_proxy_role.name
  policy_arn = "arn:aws:iam::aws:policy/SecretsManagerReadWrite"
}
