locals {
  lambda_name = "${var.prefix}-rusthumb-lambda"
}

module "ecr" {
  source = "./ecr"
  aws    = var.aws
  name   = "${var.prefix}-rusthumb"
}

module "iam" {
  source      = "./iam"
  aws         = var.aws
  prefix      = var.prefix
  lambda_name = local.lambda_name
}

module "lambda" {
  source      = "./lambda"
  name        = local.lambda_name
  ecr         = module.ecr.repository
  iam         = module.iam
  environment = {}
}
