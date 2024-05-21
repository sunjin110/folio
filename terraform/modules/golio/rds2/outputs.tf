output "datasource" {
  value = "postgres://${local.user}:${local.password}@${aws_rds_cluster.this.endpoint}:${local.port}/${local.database}"
}
