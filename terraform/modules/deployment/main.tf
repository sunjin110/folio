terraform {
  required_providers {
    cloudflare = {
      source  = "cloudflare/cloudflare"
      version = "~> 4"
    }
  }
}

# cloudflare
resource "github_actions_secret" "cloudflare_api_token" {
  repository      = "folio"
  secret_name     = "CLOUDFLARE_API_TOKEN"
  plaintext_value = cloudflare_api_token.cloudflare_api_token.value
}

resource "github_actions_secret" "cloudflare_account_id" {
  repository      = "folio"
  secret_name     = "CLOUDFLARE_ACCOUNT_ID"
  plaintext_value = var.cloudflare.account_id
}

resource "cloudflare_api_token" "cloudflare_api_token" {
  name = "github_actions_access_token"
  policy {
    permission_groups = [
      data.cloudflare_api_token_permission_groups.all.account["Pages Write"],
      data.cloudflare_api_token_permission_groups.all.account["Pages Read"],
      data.cloudflare_api_token_permission_groups.all.account["Account Settings Read"]
    ]
    resources = {
      "com.cloudflare.api.account.${var.cloudflare.account_id}" = "*"
    }
  }

  policy {
    permission_groups = [
      data.cloudflare_api_token_permission_groups.all.user["User Details Read"],
      data.cloudflare_api_token_permission_groups.all.user["Memberships Read"]
    ]
    resources = {
      # ここの文字列が何かわからないので、これを究明する必要がある
      "com.cloudflare.api.user.c6c05087bac8efd53e5dff1e626574e8" = "*"
    }
  }
}

# Account permissions
data "cloudflare_api_token_permission_groups" "all" {
  # https://developers.cloudflare.com/fundamentals/api/reference/permissions/
}

# aws
