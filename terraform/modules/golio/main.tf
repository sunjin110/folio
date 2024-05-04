
module "ecr" {
  source = "./ecr"
  aws    = var.aws
  name   = "${var.prefix}-golio"
}
