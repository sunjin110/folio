resource "aws_iam_role" "lambda" {
  name               = "${var.prefix}-golio-${var.aws.region}-lambda"
  path               = "/service-role/"
  assume_role_policy = data.aws_iam_policy_document.lambda.json
}


data "aws_iam_policy_document" "lambda" {
  version = "2012-10-17"
  statement {
    sid = "GolioLambda"
    actions = [
      "sts:AssumeRole"
    ]
    effect = "Allow"
    principals {
      type        = "Service"
      identifiers = ["lambda.amazonaws.com"]
    }
  }
}

resource "aws_iam_role_policy_attachment" "lambda_basic" {
  role       = aws_iam_role.lambda.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}
