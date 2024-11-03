variable "s3" {
  type = object({
    id: string
    arn : string
  })
}

variable "lambda" {
  type = object({
    function_name : string
    arn : string
  })
}
