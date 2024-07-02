resource "aws_dynamodb_table" "users" {
  name         = "${var.prefix}_users"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "email"

  attribute {
    name = "email"
    type = "S"
  }

  tags = {
    Name = "${var.prefix}_users"
  }
}
