output "lambda" {
  value = {
    function_name = aws_lambda_function.this.function_name
    invoke_arn    = aws_lambda_function.this.invoke_arn
    arn           = aws_lambda_function.this.arn
    # security_group_id = aws_security_group.this.id
  }
}

