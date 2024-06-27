locals {
  lambda_name                 = "${var.prefix}-golio-lambda"
  media_s3_name               = "${var.prefix}-golio-media"
  dynamodb_user_sessions_name = "${var.prefix}_user_sessions"
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

    POSTGRES_DATASOURCE : module.rds.datasource,
    MEDIA_S3_REGION : var.aws.region,
    MEDIA_S3_BUCKET_NAME : local.media_s3_name,

    SESSION_DYNAMODB_TABLE_NAME : local.dynamodb_user_sessions_name,

    CHAT_GPT_API_KEY : var.chat_gpt.api_key

    GOOGLE_CUSTOM_SEARCH_KEY : var.google_custom_search_key,

    WORDS_API_RAPID_API_KEY : var.words_api.rapid_api_key,
    WORDS_API_RAPID_API_HOST : var.words_api.rapid_api_host,
  }

  network = module.network.network
  prefix  = var.prefix
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


module "network" {
  source     = "./network"
  cidr_block = var.cidr_block
  prefix     = var.prefix
  aws        = var.aws
}


module "rds" {
  source  = "./rds2"
  network = module.network.network
  prefix  = var.prefix
}

module "media_s3" {
  source = "./s3"
  name   = local.media_s3_name
  cors = {
    allowed_origins = [var.domain.reolio_base_url]
  }
}

module "dynamodb" {
  source = "./dynamodb"
  prefix = var.prefix
}

module "aws_translate" {
  source             = "./aws_translate"
  prefix             = var.prefix
  aws                = var.aws
  network            = module.network.network
  security_group_ids = [module.lambda.lambda.security_group_id]
}

module "google_project_service" {
  source = "./google_project_service"
  gcp    = var.gcp
}

module "google_service_account" {
  source = "./google_service_account"
  prefix = var.prefix
  depends_on = [
    module.google_project_service
  ]
}

module "neon" {
  source = "./neon"
  prefix = var.prefix
}
