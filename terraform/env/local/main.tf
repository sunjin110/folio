terraform {
  required_version = "> 1.3.0"
  required_providers {
    cloudflare = {
      source  = "cloudflare/cloudflare"
      version = "~> 4"
    }
    #  neon = {
    #   source = "kislerdm/neon"
    # }
  }
}

# provider "neon" {
#   api_key = var.neon_api_key
# }

provider "cloudflare" {
  api_key = var.cloudflare_api_key
  email   = var.cloudflare_email
}

module "session_repository" {
  source                = "../../modules/kv_store"
  name                  = "local-folio-session"
  cloudflare_account_id = var.cloudflare_account_id
}

module "database" {
  source                = "../../modules/d1"
  name                  = "local-folio-db"
  cloudflare_account_id = var.cloudflare_account_id
}

output "session_repository_token" {
  sensitive = true
  value     = module.session_repository.api_token
}

output "session_repository_namespace_id" {
  value = module.session_repository.namespace_id
}

output "database_token" {
  sensitive = true
  value     = module.database.api_token
}

output "database_id" {
  sensitive = true
  value     = module.database.db_id
}