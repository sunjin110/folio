
# S3イベント通知の設定
resource "aws_s3_bucket_notification" "bucket_notification" {
  bucket = var.s3.id

  lambda_function {
    lambda_function_arn = var.lambda.arn
    events              = ["s3:ObjectCreated:*"]
    filter_prefix       = "default/"
  }

  lambda_function {
    lambda_function_arn = var.lambda.arn
    events              = ["s3:ObjectCreated:*"]
    filter_prefix       = "lime/"
  }

  depends_on = [
    aws_lambda_permission.allow_s3_invoke
  ]
}

# LambdaにS3からのトリガーを許可する
resource "aws_lambda_permission" "allow_s3_invoke" {
  statement_id  = "GombAllowS3Invoke"
  action        = "lambda:InvokeFunction"
  function_name = var.lambda.function_name
  principal     = "s3.amazonaws.com"
  source_arn    = var.s3.arn
}
