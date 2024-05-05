variable "cloudflare" {
  type = object({
    account_id : string
    zone_id : string
  })
}


variable "cloudflare_pages" {
  type = object({
    production_branch = string
    project_name      = string
    subdomain_name    = string
  })
}
