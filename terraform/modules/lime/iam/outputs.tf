output "role" {
  value = {
    lambda = {
      arn = aws_iam_role.lambda.arn
    }
    api_gateway_integration = {
      arn = aws_iam_role.api_gateway_integration.arn
    }
  }
}
