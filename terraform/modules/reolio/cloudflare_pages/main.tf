terraform {
  required_providers {
    cloudflare = {
      source  = "cloudflare/cloudflare"
      version = "~> 4"
    }
  }
}

resource "cloudflare_pages_project" "this" {
  account_id        = var.cloudflare.account_id
  name              = var.name
  production_branch = var.production_branch

  deployment_configs {
    production {
      environment_variables = {
        REACT_APP_GOLIO_BASE_URL : var.golio.base_url,
      }
    }
    preview {
      environment_variables = {
        REACT_APP_GOLIO_BASE_URL : var.golio.base_url,
      }
    }
  }
}

