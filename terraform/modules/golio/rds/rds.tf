locals {
  user    = "golion"
  db_name = "golio"
}

resource "aws_rds_cluster" "this" {
  cluster_identifier = "${var.prefix}-golio-db"
  engine             = "aurora-postgresql"
  engine_version     = "16.2"
  engine_mode        = "provisioned" # aurora serverless v2はprovisionedエンジンモードで動作する https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/rds_cluster
  database_name      = local.db_name
  master_username    = local.user
  master_password    = random_password.user_password.result

  skip_final_snapshot = true

  serverlessv2_scaling_configuration {
    max_capacity = 1.0
    min_capacity = 0.5
  }

  vpc_security_group_ids = [aws_security_group.this.id]
  db_subnet_group_name   = aws_db_subnet_group.this.name
}

resource "aws_db_subnet_group" "this" {
  name       = "${var.prefix}-golio-db-subnet-group"
  subnet_ids = var.network.private_subnet_ids
  tags = {
    Name = "${var.prefix}-golio-db-subnet-group"
  }
}

resource "random_password" "user_password" {
  length  = 32
  special = false
}
