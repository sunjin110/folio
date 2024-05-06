variable "cloudflare" {
  type = object({
    account_id : string
  })
}

variable "aws" {
  type = object({
    account_id : string
    region : string
    profile : string
  })
}

variable "aws_deploy" {
  type = object({
    role_name : string
  })
}