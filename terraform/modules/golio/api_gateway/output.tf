output "api_gateway" {
  value = {
    rest_api_id : aws_api_gateway_rest_api.this.id
    domain_name : "${aws_api_gateway_rest_api.this.id}.execute-api.${var.aws.region}.amazonaws.com"
    arn : aws_api_gateway_rest_api.this.arn
    stage_arn : aws_api_gateway_stage.this.arn
    invoke_url : aws_api_gateway_stage.this.invoke_url
  }
}
