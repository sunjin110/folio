resource "aws_lambda_function" "this" {
  function_name = var.name
  timeout       = 60 * 15 # 15分 最大
  image_uri     = "${var.ecr.repository_url}:latest"
  package_type  = "Image"
  role          = var.iam.role.lambda.arn
  architectures = ["arm64"]

  memory_size = 1024 * 2 # 2GB 画像とか扱うから
  environment {
    variables = var.environment
  }
}
