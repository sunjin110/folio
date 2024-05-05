terraform {
  required_providers {
    cloudflare = {
      source  = "cloudflare/cloudflare"
      version = "~> 4"
    }
  }
}

resource "github_actions_secret" "cloudflare_api_token" {
  repository      = "folio"
  secret_name     = "CLOUDFLARE_PAGES_DEPLOY_API_KEY"
  plaintext_value = cloudflare_api_token.cloudflare_api_token.value
}

resource "github_actions_secret" "cloudflare_account_id" {
  repository = "folio"
  secret_name = "CLOUDFLARE_ACCOUNT_ID"
  plaintext_value = var.cloudflare.account_id
}

resource "cloudflare_api_token" "cloudflare_api_token" {
  name = "github_actions_access_token"
  policy {
    permission_groups = [
      data.cloudflare_api_token_permission_groups.all.account["Pages Write"],
      data.cloudflare_api_token_permission_groups.all.account["Pages Read"]
    ]
    resources = {
      "com.cloudflare.api.account.${var.cloudflare.account_id}" = "*"
    }
  }

}

# Account permissions
data "cloudflare_api_token_permission_groups" "all" {
  # https://developers.cloudflare.com/fundamentals/api/reference/permissions/
}
