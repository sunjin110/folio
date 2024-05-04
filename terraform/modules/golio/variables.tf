variable "aws" {
  type = object({
    account_id = string,
    region     = string,
    profile    = string,
  })
}

variable "cloudflare" {
  type = object({
    account_id = string
  })
}

variable "prefix" {
  type = string
}
