output "network" {
  value = {
    vpc_id = aws_vpc.this.id
    private_cidr_blocks = [
      aws_subnet.private_subnet_1.cidr_block,
      aws_subnet.private_subnet_2.cidr_block,
    ]
    private_subnet_ids = [
      aws_subnet.private_subnet_1.id,
      aws_subnet.private_subnet_2.id,
    ]
  }
}
