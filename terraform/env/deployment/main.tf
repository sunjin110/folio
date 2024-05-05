terraform {
  required_providers {
    github = {
      source  = "integrations/github"
      version = "~> 6.0"
    }
    cloudflare = {
      source  = "cloudflare/cloudflare"
      version = "~> 4"
    }
  }
}

provider "github" {
  token = var.github_token
}

provider "cloudflare" {
  api_key = var.cloudflare_api_key
  email   = var.cloudflare_email
}

module "deployment" {
  source = "../../modules/deployment"
  cloudflare = {
    account_id = var.cloudflare_account_id
  }
}