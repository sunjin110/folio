resource "aws_iam_role" "api_gateway_integration" {
  name               = "${var.prefix}-golio-${var.aws.region}-api-gateway-integration"
  path               = "/service-role/"
  assume_role_policy = data.aws_iam_policy_document.api_gateway_integration.json
}

data "aws_iam_policy_document" "api_gateway_integration" {
  version = "2012-10-17"
  statement {
    sid     = "GolioApiGateway"
    actions = ["sts:AssumeRole"]
    effect  = "Allow"
    principals {
      type        = "Service"
      identifiers = ["apigateway.amazonaws.com"]
    }
  }
}

resource "aws_iam_role_policy_attachment" "api_gateway_intergration_lambda_invoke" {
  role       = aws_iam_role.api_gateway_integration.name
  policy_arn = aws_iam_policy.api_gateway_integration_lambda_invoke.arn
}

resource "aws_iam_policy" "api_gateway_integration_lambda_invoke" {
  name   = "${var.prefix}-golio-${var.aws.region}-api-gateway-lambda-invoke"
  policy = data.aws_iam_policy_document.api_gateway_integration_lambda_invoke.json
}

data "aws_iam_policy_document" "api_gateway_integration_lambda_invoke" {
  version = "2012-10-17"
  statement {
    sid     = ""
    actions = ["lambda:InvokeFunction"]
    effect  = "Allow"
    resources = [
      "arn:aws:lambda:${var.aws.region}:${var.aws.account_id}:function:${var.lambda_name}"
    ]
  }
}
