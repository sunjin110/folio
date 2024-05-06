module "cloudflare_pages" {
  source            = "./cloudflare_pages"
  production_branch = var.cloudflare_pages.production_branch
  name              = var.cloudflare_pages.project_name
  cloudflare        = var.cloudflare
  golio             = var.golio
}

module "domain" {
  source     = "./domain"
  value      = "${var.cloudflare_pages.project_name}.pages.dev"
  cloudflare = var.cloudflare
  cloudflare_pages = {
    project_name   = var.cloudflare_pages.project_name
    subdomain_name = var.cloudflare_pages.subdomain_name
    domain         = "${var.cloudflare_pages.subdomain_name}.sunjin.info"
  }
  depends_on = [module.cloudflare_pages]
}
