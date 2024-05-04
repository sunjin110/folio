variable "aws" {
  type = object({
    account_id : string
    region : string
    profile : string
  })
}

variable "lambda" {
  type = object({
    invoke_arn : string
  })
}

variable "iam" {
  type = object({
    role : object({
      api_gateway_integration : object({
        arn = string
      })
    })
  })
}

variable "prefix" {
  type = string
}
