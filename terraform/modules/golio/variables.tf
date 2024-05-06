variable "aws" {
  type = object({
    account_id = string,
    region     = string,
    profile    = string,
  })
}

variable "cloudflare" {
  type = object({
    account_id = string
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