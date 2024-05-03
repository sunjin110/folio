terraform {
  required_providers {
    cloudflare = {
      source  = "cloudflare/cloudflare"
      version = "~> 4"
    }
  }
}


resource "cloudflare_d1_database" "this" {
  account_id = var.cloudflare_account_id
  name       = var.name
}

# token api

# Account permissions
data "cloudflare_api_token_permission_groups" "all" {
  # https://developers.cloudflare.com/fundamentals/api/reference/permissions/
}

resource "cloudflare_api_token" "this" {
  name = "${var.name}_access_token"
  policy {
    permission_groups = [
      data.cloudflare_api_token_permission_groups.all.account["D1 Read"],
      data.cloudflare_api_token_permission_groups.all.account["D1 Write"]
    ]
    resources = {
      "com.cloudflare.api.account.${var.cloudflare_account_id}" = "*"
    }
  }
}
