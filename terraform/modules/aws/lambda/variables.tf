variable "environment" {
  type    = map(string)
  default = {}
}

variable "name" {
  type = string
}

variable "prefix" {
  type = string
}

variable "ecr" {
  type = object({
    repository_url = string
    arn            = string
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

variable "architecture" {
  type        = string
  description = "arm64 or amd64"
}

variable "memory_size" {
  type    = number
  default = 128 # MB
}

variable "timeout" {
  type    = number
  default = 30 # seconds
}
