terraform {
  required_version = "> 1.3.0"
  required_providers {
    github = {
      source  = "integrations/github"
      version = "~> 6.0"
    }
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
    bucket         = "deployment-folio-terraform-state"
    key            = "terraform.tfstate"
    region         = "ap-northeast-1"
    dynamodb_table = "deployment-folio-terraform-state"
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
}

provider "github" {
  token = var.github_token
}

provider "cloudflare" {
  api_key = var.cloudflare_api_key
  email   = var.cloudflare_email
}

provider "aws" {
  region  = "ap-northeast-1"
  profile = "folio-terraform"
}

module "deployment" {
  source = "../../modules/deployment"
  cloudflare = {
    account_id = var.cloudflare_account_id
  }
  aws = local.aws
  aws_deploy = {
    role_name = "githubactions"
  }
}

module "tfstate" {
  source = "../../modules/tfstate"
  name   = "deployment-folio-terraform-state"
}

data "aws_caller_identity" "this" {}

data "aws_region" "this" {}
