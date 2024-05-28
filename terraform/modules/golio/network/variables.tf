variable "cidr_block" {
  type = string
}

variable "prefix" {
  type = string
}

variable "aws" {
  type = object({
    region     = string
    account_id = string
  })
}
