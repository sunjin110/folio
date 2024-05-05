variable "cloudflare" {
  type = object({
    account_id : string
  })
}


variable "cloudflare_pages" {
  type = object({
    production_branch = string
    name              = string
  })
}
