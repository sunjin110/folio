variable "aws" {
  type = object({
    account_id : string,
    region : string,
    profile : string,
  })
}

variable "prefix" {
  type = string
}
