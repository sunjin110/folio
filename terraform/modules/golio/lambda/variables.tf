variable "name" {
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
    GOOGLE_OAUTH_CLIENT_ID : string,
    GOOGLE_OAUTH_CLIENT_SECRET : string,
    GOOGLE_OAUTH_REDIRECT_URI : string,
    GOOGLE_OAUTH_CALLBACK_REDIRECT_URI : string,

    SESSION_KV_STORE_ACCOUNT_ID : string,
    SESSION_KV_STORE_NAMESPACE_ID : string,
    SESSION_KV_STORE_API_TOKEN : string,

    D1_DATABASE_ACCOUNT_ID : string,
    D1_DATABASE_DATABASE_ID : string,
    D1_DATABASE_API_TOKEN : string,

    CORS_ALLOWED_ORIGINS : string,
  })
}