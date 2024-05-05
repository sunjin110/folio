
variable "cloudflare" {
  type = object({
    zone_id    = string
    account_id = string
  })
}

variable "cloudflare_pages" {
  type = object({
    project_name   = string
    subdomain_name = string
    domain         = string
  })
}

variable "value" {
  description = "pages"
  type        = string
}
