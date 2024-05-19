resource "aws_security_group" "this" {
  name                   = "${var.prefix}-golio"
  description            = "${var.prefix}-golio lambda security group"
  revoke_rules_on_delete = true
  vpc_id                 = var.network.vpc_id
}

resource "aws_security_group_rule" "allow_egress_rdb_access" {
  security_group_id = aws_security_group.this.id
  type              = "egress"
  description       = "allow access to db"
  protocol          = "tcp"
  cidr_blocks       = var.network.private_cidr_blocks
  from_port         = 5432
  to_port           = 5432
}

