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

variable "environment" {
  type = object({
    LINE_CHANNEL_SECRET : string
    LINE_CHANNEL_TOKEN : string
    MEDIA_S3_BUCKET_NAME : string,
  })
}
