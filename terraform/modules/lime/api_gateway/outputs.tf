output "rest_api_id" {
  value = aws_api_gateway_rest_api.this.id
}

output "domain_name" {
  value = "${aws_api_gateway_rest_api.this.id}.execute-api.${var.aws.region}.amazonaws.com"
}

output "arn" {
  value = aws_api_gateway_rest_api.this.arn
}

output "stage_arn" {
  value = aws_api_gateway_stage.this.arn
}

output "invoke_url" {
  value = aws_api_gateway_stage.this.invoke_url
}

output "stage_name" {
  value = local.stage_name
}