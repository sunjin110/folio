variable "aws" {
  type = object({
    account_id = string,
    region     = string,
    profile    = string,
  })
}

variable "gcp" {
  type = object({
    project_id = string
  })
}

variable "cloudflare" {
  type = object({
    account_id = string
    zone_id    = string
  })
}

variable "prefix" {
  type = string
}

variable "google_oauth" {
  type = object({
    client_id             = string
    client_secret         = string
    redirect_uri          = string
    callback_redirect_uri = string
  })
}

variable "reolio" {
  type = object({
    base_url : string
  })
}

variable "domain" {
  type = object({
    domain_name : string
    name : string
    reolio_base_url : string
  })
}
variable "cidr_block" {
  type = string
}

variable "chat_gpt" {
  type = object({
    api_key : string
  })
}

variable "google_custom_search_key" {
  type = string
}

variable "words_api" {
  type = object({
    rapid_api_key : string,
    rapid_api_host : string,
  })
}