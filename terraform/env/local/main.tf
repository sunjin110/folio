terraform {
  required_providers {
    cloudflare = {
        source = "cloudflare/cloudflare"
        version = "~> 4"
    }
  }
}

provider "cloudflare" {
    api_key = var.cloudflare_api_key
    email = var.cloudflare_email
}

module "session_repository" {
  source = "../../modules/kv_store"

  name = "local-folio-session"
  cloudflare_account_id = var.cloudflare_account_id
}

output "session_repository_token" {
  sensitive = true
  value = module.session_repository.api_token
}

output "session_repository_namespace_id" {
  value = module.session_repository.namespace_id
}