

# resource "aws_vpc_endpoint" "aws_translate" {
#   vpc_id             = var.network.vpc_id
#   service_name       = "com.amazonaws.${var.aws.region}.translate"
#   vpc_endpoint_type  = "Interface"
#   security_group_ids = var.security_group_ids

#   tags = {
#     Name = "${var.prefix}-aws-translate-vpc-endpoint"
#   }
# }
