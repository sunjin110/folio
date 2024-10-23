variable "aws" {
  type = object({
    account_id : string,
    region : string,
    profile : string,
  })
}

variable "cloudflare" {
  type = object({
    account_id = string
    zone_id    = string
  })
}

variable "prefix" {
  type = string
}

variable "domain" {
  type = object({
    domain_name : string
    name : string
  })
}

variable "line" {
  type = object({
    channel_secret : string
  })
}
