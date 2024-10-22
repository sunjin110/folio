variable "domain_name" {
  type = string
}

variable "cloudflare" {
  type = object({
    zone_id : string
  })
}