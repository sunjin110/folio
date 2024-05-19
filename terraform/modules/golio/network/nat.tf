resource "aws_eip" "nat_1" {
  domain = "vpc"
  tags = {
    Name = "${var.prefix}-golio-nat-1"
  }
}

resource "aws_eip" "nat_2" {
  domain = "vpc"
  tags = {
    Name = "${var.prefix}-golio-nat-2"
  }
}

resource "aws_nat_gateway" "nat_gateway_1" {
  allocation_id = aws_eip.nat_1.id
  subnet_id     = aws_subnet.public_subnet_1.id
  tags = {
    Name = "${var.prefix}-golio-nat-gateway-1"
  }
}

resource "aws_nat_gateway" "nat_gateway_2" {
  allocation_id = aws_eip.nat_2.id
  subnet_id     = aws_subnet.public_subnet_2.id
  tags = {
    Name = "${var.prefix}-golio-nat-gateway-2"
  }
}

resource "aws_internet_gateway" "internet_gateway" {
  vpc_id = aws_vpc.this.id
  tags = {
    Name = "${var.prefix}-golio-internet-gateway"
  }
}
