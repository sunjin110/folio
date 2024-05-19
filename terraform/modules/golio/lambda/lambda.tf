resource "aws_lambda_function" "this" {
  function_name = var.name
  timeout       = 30 # seconds
  image_uri     = "${var.ecr.repository_url}:latest"
  package_type  = "Image"
  role          = var.iam.role.lambda.arn
  architectures = ["arm64"]
  # https://docs.aws.amazon.com/lambda/latest/dg/gettingstarted-limits.html
  memory_size = 128 # MB
  environment {
    variables = var.environment
  }

  vpc_config {
    subnet_ids         = var.network.private_subnet_ids
    security_group_ids = [aws_security_group.this.id]
  }
}
