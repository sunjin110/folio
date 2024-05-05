module "cloudflare_pages" {
  source            = "./cloudflare_pages"
  production_branch = var.cloudflare_pages.production_branch
  name              = var.cloudflare_pages.name
  cloudflare        = var.cloudflare
}

