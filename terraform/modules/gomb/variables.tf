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

variable "media_s3" {
  type = object({
    id: string
    name : string
    arn : string
  })
}
