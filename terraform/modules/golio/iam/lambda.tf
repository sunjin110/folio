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

resource "aws_iam_role_policy_attachment" "lambda_vpc_basi" {
  role       = aws_iam_role.lambda.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaVPCAccessExecutionRole"
}

resource "aws_iam_role_policy_attachment" "lambda_s3_access" {
  role       = aws_iam_role.lambda.name
  policy_arn = aws_iam_policy.lambda_s3_access.arn
}

resource "aws_iam_role_policy_attachment" "lambda_rds_access" {
  role       = aws_iam_role.lambda.name
  policy_arn = aws_iam_policy.lambda_rds_access.arn
}

resource "aws_iam_role_policy_attachment" "lambda_dynamodb_access" {
  role       = aws_iam_role.lambda.name
  policy_arn = aws_iam_policy.lambda_dynamodb_policy.arn
}

resource "aws_iam_policy" "lambda_s3_access" {
  name = "${var.prefix}-lambda-s3-access"

  # TODO 権限をもうちょい絞る
  policy = jsonencode({
    Version : "2012-10-17",
    Statement = [
      {
        Action = [
          "s3:*"
        ],
        Resource = "*",
        Effect   = "Allow"
      }
    ]
  })
}

resource "aws_iam_policy" "lambda_rds_access" {
  name = "${var.prefix}-lambda-rds-access"

  # TODO 権限をもうちょい絞る
  policy = jsonencode({
    Version : "2012-10-17",
    Statement = [
      {
        Action = [
          "rds:*"
        ],
        Resource = "*",
        Effect   = "Allow"
      }
    ]
  })
}

resource "aws_iam_policy" "lambda_dynamodb_policy" {
  name = "${var.prefix}-lambda-dynamodb-policy"

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Action = [
          "dynamodb:GetItem",
          "dynamodb:PutItem",
          "dynamodb:UpdateItem",
          "dynamodb:DeleteItem",
          "dynamodb:Scan",
          "dynamodb:Query",
          "dynamodb:BatchGetItem",
          "dynamodb:BatchWriteItem"
        ]
        Resource = "arn:aws:dynamodb:${var.aws.region}:${var.aws.account_id}:table/*"
      }
    ]
  })
}