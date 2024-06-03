resource "aws_vpc_endpoint" "dynamodb" {
  vpc_id            = aws_vpc.this.id
  service_name      = "com.amazonaws.${var.aws.region}.dynamodb"
  vpc_endpoint_type = "Gateway"
  route_table_ids = [
    aws_route_table.private_route_table_1.id,
    aws_route_table.private_route_table_2.id,
  ]

  tags = {
    Name = "${var.prefix}-dynamodb-vpc-endpoint"
  }
}
