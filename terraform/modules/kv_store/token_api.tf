
# Account permissions
data "cloudflare_api_token_permission_groups" "all" {
  # https://developers.cloudflare.com/fundamentals/api/reference/permissions/
}

resource "cloudflare_api_token" "this" {
  name = "${var.name}_access_token"
  policy {
    permission_groups = [
      data.cloudflare_api_token_permission_groups.all.account["Workers KV Storage Write"],
      data.cloudflare_api_token_permission_groups.all.account["Workers KV Storage Read"]
    ]
    resources = {
      "com.cloudflare.api.account.${var.cloudflare_account_id}" = "*"
    }
  }
}
