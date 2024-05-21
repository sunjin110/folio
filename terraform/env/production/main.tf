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

locals {
  aws = {
    account_id = data.aws_caller_identity.this.account_id
    region     = data.aws_region.this.name
    profile    = "folio-terraform"
  }
  reolio_base_url = "https://folio.sunjin.info"
  golio_domain    = "folio-api.sunjin.info"
  tfstate_name    = "production-folio-terraform-state"
}

module "golio" {
  source = "../../modules/golio"

  providers = {
    aws          = aws
    aws.virginia = aws.virginia
  }

  aws = local.aws
  cloudflare = {
    account_id = var.cloudflare_account_id
    zone_id    = var.cloudflare_zone_id
  }

  prefix = "production"
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
    domain_name = local.golio_domain
    name        = "folio-api"
  }

  cidr_block = "10.0.0.0/16"
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

module "tfstate" {
  source = "../../modules/tfstate"
  name   = local.tfstate_name
}

data "aws_caller_identity" "this" {}

data "aws_region" "this" {}
