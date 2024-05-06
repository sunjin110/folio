
provider "aws" {
  region  = "ap-northeast-1"
  profile = "folio-terraform"
}

# ACM SSL証明書用provider(us-east-1で作成することが必要)
provider "aws" {
  region = "us-east-1"
  alias  = "virginia"

  profile = "folio-terraform"
}


provider "cloudflare" {
  api_key = var.cloudflare_api_key
  email   = var.cloudflare_email
}

