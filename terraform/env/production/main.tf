terraform {
  required_version = "> 1.3.0"
  required_providers {
    cloudflare = {
      source  = "cloudflare/cloudflare"
      version = "~> 4"
    }
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0.1"
    }

    neon = {
      source  = "kislerdm/neon"
      version = ">= 0.5.0"
    }
  }
  backend "s3" {
    bucket         = "production-folio-terraform-state"
    key            = "terraform.tfstate"
    region         = "ap-northeast-1"
    dynamodb_table = "production-folio-terraform-state"
    encrypt        = true
    profile        = "folio-terraform"
  }
}

provider "neon" {
  api_key = var.neon_api_key
}

locals {
  aws = {
    account_id = data.aws_caller_identity.this.account_id
    region     = data.aws_region.this.name
    profile    = "folio-terraform"
  }
  gcp = {
    project_id = "folio-sunjin"
  }
  reolio_base_url = "https://folio.sunjin.info"
  golio_domain    = "folio-api.sunjin.info"
  lime_domain     = "lime.sunjin.info"
  tfstate_name    = "production-folio-terraform-state"
  env             = "production"
}

module "golio" {
  source = "../../modules/golio"

  providers = {
    aws          = aws
    aws.virginia = aws.virginia
  }

  aws = local.aws
  gcp = local.gcp
  cloudflare = {
    account_id = var.cloudflare_account_id
    zone_id    = var.cloudflare_zone_id
  }

  prefix = local.env
  google_oauth = {
    client_id             = var.google_oauth_client_id
    client_secret         = var.google_oauth_secret_id
    redirect_uri          = "https://${local.golio_domain}/auth/google-oauth/callback"
    callback_redirect_uri = "${local.reolio_base_url}/login"
  }

  reolio = {
    base_url = local.reolio_base_url
  }

  domain = {
    domain_name     = local.golio_domain
    name            = "folio-api"
    reolio_base_url = local.reolio_base_url
  }

  cidr_block = "10.0.0.0/16"

  chat_gpt = {
    api_key = var.chat_gpt_api_key
  }

  google_custom_search_key = var.google_custom_search_key

  words_api = {
    rapid_api_key  = var.words_api_rapid_api_key
    rapid_api_host = var.words_api_rapid_api_host
  }
}

module "reolio" {
  source = "../../modules/reolio"
  cloudflare = {
    account_id = var.cloudflare_account_id
    zone_id    = var.cloudflare_zone_id
  }

  cloudflare_pages = {
    production_branch = "main"
    project_name      = "reolio"
    subdomain_name    = "folio"
  }
}

module "lime" {
  source = "../../modules/lime"
  providers = {
    aws          = aws
    aws.virginia = aws.virginia
  }
  aws    = local.aws
  prefix = local.env
  domain = {
    domain_name = local.lime_domain
    name        = "lime"
  }
  cloudflare = {
    account_id = var.cloudflare_account_id
    zone_id    = var.cloudflare_zone_id
  }
  line = {
    channel_secret = var.line_channel_secret
    channel_token  = var.line_channel_token
  }
}

module "gomb" {
  source   = "../../modules/gomb"
  aws      = local.aws
  prefix   = local.env
  media_s3 = module.golio.media_s3
}

module "rusthumb" {
  source = "../../modules/rusthumb"
  prefix = local.env
  aws    = local.aws
}

module "tfstate" {
  source = "../../modules/tfstate"
  name   = local.tfstate_name
}

data "aws_caller_identity" "this" {}

data "aws_region" "this" {}

# なんかいつもここの依存が外れる
# import {
#   to = module.lime.module.api_gateway.aws_api_gateway_stage.this
#   id = "xzyga1fse4/v1"
# }
