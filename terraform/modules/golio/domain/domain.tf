terraform {
  required_providers {
    cloudflare = {
      source  = "cloudflare/cloudflare"
      version = "~> 4"
    }
  }
}

resource "aws_api_gateway_domain_name" "this" {
  domain_name     = var.domain_name
  certificate_arn = var.acm.certificate_arn
}

resource "aws_api_gateway_base_path_mapping" "this" {
  domain_name = aws_api_gateway_domain_name.this.domain_name
  api_id      = var.api_gateway.rest_api_id
  stage_name  = var.api_gateway.stage_name
}

resource "cloudflare_record" "this" {
  zone_id = var.cloudflare.zone_id
  name    = var.name
  value   = aws_api_gateway_domain_name.this.cloudfront_domain_name
  type    = "CNAME"
  proxied = false
}

