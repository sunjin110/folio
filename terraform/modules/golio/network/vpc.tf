resource "aws_vpc" "this" {
  tags = {
    Name = "${var.prefix}-golio-vpc"
  }
  cidr_block = var.cidr_block
}

resource "aws_subnet" "private_subnet_1" {
  vpc_id                              = aws_vpc.this.id
  cidr_block                          = cidrsubnet(aws_vpc.this.cidr_block, 4, 8)
  private_dns_hostname_type_on_launch = "ip-name"
  availability_zone_id                = data.aws_availability_zones.available.zone_ids[0]
}

resource "aws_subnet" "private_subnet_2" {
  vpc_id                              = aws_vpc.this.id
  cidr_block                          = cidrsubnet(aws_vpc.this.cidr_block, 4, 9)
  private_dns_hostname_type_on_launch = "ip-name"
  availability_zone_id                = data.aws_availability_zones.available.zone_ids[1]
}

data "aws_availability_zones" "available" {}
