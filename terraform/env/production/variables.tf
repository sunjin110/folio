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

variable "chat_gpt_api_key" {
  type      = string
  sensitive = true
}

variable "google_custom_search_key" {
  type      = string
  sensitive = true
}

variable "words_api_rapid_api_key" {
  type      = string
  sensitive = true
}

variable "words_api_rapid_api_host" {
  type      = string
  sensitive = true
}

variable "neon_api_key" {
  type = string
}

variable "line_channel_secret" {
  type      = string
  sensitive = true
}

variable "line_channel_token" {
  type = string
}


# TODO これはapiのcustom domainを決めたタイミングでどうにかする
# variable "google_oauth_redirect_uri" {
#   type = string
# }

# variable "google_oauth_callback_redirect_uri" {
#   type = string
# }

