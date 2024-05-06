locals {
  lambda_name = "${var.prefix}-golio-lambda"
}

module "ecr" {
  source = "./ecr"
  aws    = var.aws
  name   = "${var.prefix}-golio"
}

module "iam" {
  source      = "./iam"
  aws         = var.aws
  prefix      = var.prefix
  lambda_name = local.lambda_name
}

module "lambda" {
  source = "./lambda"
  name   = local.lambda_name
  ecr    = module.ecr.repository
  iam    = module.iam
  environment = {
    GOOGLE_OAUTH_CLIENT_ID : var.google_oauth.client_id,
    GOOGLE_OAUTH_CLIENT_SECRET : var.google_oauth.client_secret,
    GOOGLE_OAUTH_REDIRECT_URI : var.google_oauth.redirect_uri,
    GOOGLE_OAUTH_CALLBACK_REDIRECT_URI : var.google_oauth.callback_redirect_uri,

    SESSION_KV_STORE_ACCOUNT_ID : var.cloudflare.account_id,
    SESSION_KV_STORE_NAMESPACE_ID : module.session_repository.namespace_id,
    SESSION_KV_STORE_API_TOKEN : module.session_repository.api_token,

    D1_DATABASE_ACCOUNT_ID : var.cloudflare.account_id,
    D1_DATABASE_DATABASE_ID : module.d1.db_id,
    D1_DATABASE_API_TOKEN : module.d1.api_token

    CORS_ALLOWED_ORIGINS : var.reolio.base_url
  }
}

module "api_gateway" {
  source = "./api_gateway"
  aws    = var.aws
  lambda = module.lambda.lambda
  iam    = module.iam
  prefix = var.prefix
}

module "session_repository" {
  source                = "../kv_store"
  name                  = "${var.prefix}-folio-session"
  cloudflare_account_id = var.cloudflare.account_id
}

module "d1" {
  source                = "../d1"
  name                  = "${var.prefix}-folio-db"
  cloudflare_account_id = var.cloudflare.account_id
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
  api_gateway = module.api_gateway.api_gateway
  cloudflare  = var.cloudflare
  acm         = module.acm.acm
}
