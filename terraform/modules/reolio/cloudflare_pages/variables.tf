variable "cloudflare" {
  type = object({
    account_id : string
  })
}

variable "name" {
  type = string
}

variable "production_branch" {
  type = string
}
