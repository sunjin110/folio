locals {
  stage_name = "v1"
}

resource "aws_api_gateway_deployment" "this" {
  depends_on  = [
    aws_api_gateway_integration.this,
    aws_api_gateway_integration.root,
  ]
  rest_api_id = aws_api_gateway_rest_api.this.id
  stage_name  = local.stage_name
  lifecycle {
    create_before_destroy = true
  }
}

resource "aws_api_gateway_stage" "this" {
  deployment_id        = aws_api_gateway_deployment.this.id
  rest_api_id          = aws_api_gateway_rest_api.this.id
  stage_name           = local.stage_name
  xray_tracing_enabled = true

  lifecycle {
    ignore_changes = [deployment_id]
  }
}
