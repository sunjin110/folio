variable "name" {
  type = string
}

variable "environment" {
  type = object({

  })
}

variable "ecr" {
  type = object({
    repository_url : string
  })
}

variable "iam" {
  type = object({
    role : object({
      lambda : object({
        arn : string
      })
    })
  })
}

