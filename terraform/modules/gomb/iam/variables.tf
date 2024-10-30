variable "prefix" {
  type = string
}

variable "aws" {
  type = object({
    account_id : string,
    region : string,
    profile : string,
  })
}
