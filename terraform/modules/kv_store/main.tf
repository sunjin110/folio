terraform {
  required_providers {
    cloudflare = {
        source = "cloudflare/cloudflare"
        version = "~> 4"
    }
  }
}

resource "cloudflare_workers_kv_namespace" "this" {
    account_id = var.cloudflare_account_id
    title = var.name
}
