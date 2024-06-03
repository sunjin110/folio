resource "aws_dynamodb_table" "user_sessions" {
  name         = "${var.prefix}_user_sessions"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "email"

  attribute {
    name = "email"
    type = "S"
  }

  attribute {
    name = "access_token"
    type = "S"
  }

  global_secondary_index {
    name            = "access_token_index"
    hash_key        = "access_token"
    projection_type = "ALL"
  }

  tags = {
    Name = "${var.prefix}_user_sessions"
  }
}
