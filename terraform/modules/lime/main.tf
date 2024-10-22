module "ecr" {
  source = "./ecr"
  aws    = var.aws
  name   = "${var.prefix}-lime"
}

module "iam" {
  source      = "./iam"
  aws         = var.aws
  prefix      = var.prefix
  lambda_name = "${var.prefix}-lime"
}

module "lambda" {
  source      = "./lambda"
  name        = "${var.prefix}-lime"
  prefix      = var.prefix
  ecr         = module.ecr.repository
  iam         = module.iam
  environment = {}
}
