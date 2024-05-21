variable "cloudflare_account_id" {
  type = string
}

variable "cloudflare_api_key" {
  type      = string
  sensitive = true
}

variable "cloudflare_email" {
  type = string
}

variable "cloudflare_zone_id" {
  type = string
}

variable "google_oauth_client_id" {
  type = string
}

variable "google_oauth_secret_id" {
  type      = string
  sensitive = true
}

# TODO これはapiのcustom domainを決めたタイミングでどうにかする
# variable "google_oauth_redirect_uri" {
#   type = string
# }

# variable "google_oauth_callback_redirect_uri" {
#   type = string
# }

