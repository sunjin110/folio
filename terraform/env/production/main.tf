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

locals {
  aws = {
    account_id = data.aws_caller_identity.this.account_id
    region     = data.aws_region.this.name
    profile    = "folio-terraform"
  }
}

module "golio" {
  source = "../../modules/golio"
  #   providers = {
  #     aws = aws
  #   }

  aws    = local.aws
  prefix = "production"
}

data "aws_caller_identity" "this" {}

data "aws_region" "this" {}
