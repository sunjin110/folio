terraform {
  required_providers {
    cloudflare = {
      source  = "cloudflare/cloudflare"
      version = "~> 4"
    }
  }
}

resource "cloudflare_record" "this" {
  zone_id = var.cloudflare.zone_id
  name    = var.cloudflare_pages.subdomain_name
  type    = "CNAME"
  value   = var.value
}

resource "cloudflare_pages_domain" "this" {
  account_id   = var.cloudflare.account_id
  project_name = var.cloudflare_pages.project_name
  domain       = var.cloudflare_pages.domain
}
