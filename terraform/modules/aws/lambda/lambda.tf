resource "aws_lambda_function" "this" {
  function_name = var.name
  timeout       = var.timeout
  image_uri     = "${var.ecr.repository_url}:latest"
  package_type  = "Image"
  role          = var.iam.role.lambda.arn
  architectures = [var.architecture]
  # https://docs.aws.amazon.com/lambda/latest/dg/gettingstarted-limits.html
  memory_size = var.memory_size
  environment {
    variables = var.environment
  }
}
