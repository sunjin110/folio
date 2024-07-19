variable "aws" {
  type = object({
    account_id : string,
    region : string,
    profile : string,
  })
}

variable "name" {
  type        = string
  description = "repository name"
}
