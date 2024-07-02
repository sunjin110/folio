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

resource "aws_dynamodb_table" "user_sessions_v2" {
  name         = "${var.prefix}_user_sessions_v2"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "access_token"

  attribute {
    name = "access_token"
    type = "S"
  }

  attribute {
    name = "email"
    type = "S"
  }

  global_secondary_index {
    # ログアウトできるように一括検索よう
    name            = "email_access_token_index"
    hash_key        = "email"
    range_key       = "access_token"
    projection_type = "ALL"
  }

  ttl {
    attribute_name = "expire_time" # unix time
    enabled        = true
  }

  tags = {
    Name = "${var.prefix}_user_sessions_v2"
  }
}
