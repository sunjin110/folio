output "repository" {
  value = {
    repository_url = aws_ecr_repository.this.repository_url
    arn            = aws_ecr_repository.this.arn
  }
}
