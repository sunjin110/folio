terraform {
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
}

provider "aws" {
  region  = "ap-northeast-1"
  profile = "folio-terraform"
}

provider "cloudflare" {
  api_key = var.cloudflare_api_key
  email   = var.cloudflare_email
}

locals {
  aws = {
    account_id = data.aws_caller_identity.this.account_id
    region     = data.aws_region.this.name
    profile    = "folio-terraform"
  }
}

module "golio" {
  source = "../../modules/golio"
  aws    = local.aws
  cloudflare = {
    account_id = var.cloudflare_account_id
  }

  prefix = "production"
  google_oauth = {
    client_id             = var.google_oauth_client_id
    client_secret         = var.google_oauth_secret_id
    redirect_uri          = var.google_oauth_redirect_uri
    callback_redirect_uri = var.google_oauth_callback_redirect_uri
  }

  reolio = {
    base_url = "https://folio.sunjin.info"
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

  golio = {
    base_url = module.golio.golio.base_url
  }
}

data "aws_caller_identity" "this" {}

data "aws_region" "this" {}
