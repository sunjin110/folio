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

module "api_gateway" {
  source = "./api_gateway"
  aws    = var.aws
  lambda = module.lambda
  iam    = module.iam
  prefix = var.prefix
}

module "acm" {
  source = "./acm"
  providers = {
    aws = aws.virginia
  }

  cloudflare  = var.cloudflare
  domain_name = var.domain.domain_name
}

module "domain" {
  source      = "./domain"
  domain_name = var.domain.domain_name
  name        = var.domain.name
  api_gateway = module.api_gateway
  cloudflare  = var.cloudflare
  acm         = module.acm
}