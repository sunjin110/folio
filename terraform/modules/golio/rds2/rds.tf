locals {
  user     = "golio"
  password = "goliopass"
  port     = 5432
  database = "golio"
}

resource "aws_rds_cluster" "this" {
  cluster_identifier     = "${var.prefix}-golio"
  engine                 = "aurora-postgresql"
  engine_version         = "15.4"
  engine_mode            = "provisioned"
  master_username        = local.user
  master_password        = local.password
  port                   = local.port
  database_name          = local.database
  vpc_security_group_ids = [aws_security_group.this.id]
  db_subnet_group_name   = aws_db_subnet_group.this.name
  # iam_database_authentication_enabled = true

  serverlessv2_scaling_configuration {
    min_capacity = 0.5
    max_capacity = 1.0
  }

  skip_final_snapshot = true
  apply_immediately   = true

  tags = {
    Name = "${var.prefix}-golio"
  }
}

resource "aws_rds_cluster_instance" "this" {
  cluster_identifier = aws_rds_cluster.this.id
  identifier         = "${var.prefix}-golio-serverless-instance"

  engine               = aws_rds_cluster.this.engine
  engine_version       = aws_rds_cluster.this.engine_version
  instance_class       = "db.serverless"
  db_subnet_group_name = aws_db_subnet_group.this.name

  publicly_accessible = false
}

resource "aws_db_subnet_group" "this" {
  name       = "${var.prefix}-golio-subnet-group-name"
  subnet_ids = var.network.private_subnet_ids

  tags = {
    Name = "${var.prefix}-golio"
  }
}

