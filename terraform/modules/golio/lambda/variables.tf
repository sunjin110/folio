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

    POSTGRES_DATASOURCE : string,
    MEDIA_S3_REGION : string,
    MEDIA_S3_BUCKET_NAME : string,

    SESSION_DYNAMODB_TABLE_NAME : string,

    CHAT_GPT_API_KEY : string,
    GOOGLE_CUSTOM_SEARCH_KEY : string,

    WORDS_API_RAPID_API_KEY : string,
    WORDS_API_RAPID_API_HOST : string,
  })
}

variable "network" {
  type = object({
    vpc_id : string,
    private_cidr_blocks : list(string)
    private_subnet_ids : list(string)
  })
}
