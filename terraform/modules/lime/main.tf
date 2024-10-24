locals {
  media_s3_name = "${var.prefix}-golio-media"
}

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
  source = "./lambda"
  name   = "${var.prefix}-lime"
  prefix = var.prefix
  ecr    = module.ecr.repository
  iam    = module.iam
  environment = {
    LINE_CHANNEL_SECRET  = var.line.channel_secret
    LINE_CHANNEL_TOKEN   = var.line.channel_token
    MEDIA_S3_BUCKET_NAME = local.media_s3_name
  }
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
