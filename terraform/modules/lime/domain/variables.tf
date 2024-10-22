variable "api_gateway" {
  type = object({
    rest_api_id : string
    stage_name : string
  })
}

variable "cloudflare" {
  type = object({
    zone_id : string
  })
}

variable "domain_name" {
  type = string # folio-api.sunjin.info
}

variable "name" {
  type = string # folio-api
}

variable "acm" {
  type = object({
    certificate_arn : string
  })
}
