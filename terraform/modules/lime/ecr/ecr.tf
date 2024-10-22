resource "aws_ecr_repository" "this" {
    name = var.name
    image_tag_mutability = "MUTABLE"

    image_scanning_configuration {
    scan_on_push = true
  }

  force_delete = true

  # 初回はイメージをデプロイする
  # see also: https://developer.hashicorp.com/terraform/language/resources/provisioners/syntax
  provisioner "local-exec" {
    command    = "aws --profile ${var.aws.profile} ecr get-login-password --region ${var.aws.region} | docker login --username AWS --password-stdin ${var.aws.account_id}.dkr.ecr.${var.aws.region}.amazonaws.com"
    on_failure = continue ## 他とタイミング被るとログインにコケるのでその時は無視
  }

  provisioner "local-exec" {
    working_dir = "${path.module}/../../../../lime"
    command     = "DOCKER_BUILDKIT=1 docker build -t ${var.name}:latest -f Dockerfile.lambda --platform=linux/arm64 ."
  }

  provisioner "local-exec" {
    command = "docker tag ${var.name}:latest ${var.aws.account_id}.dkr.ecr.${var.aws.region}.amazonaws.com/${var.name}:latest"
  }

  provisioner "local-exec" {
    command = "docker push ${var.aws.account_id}.dkr.ecr.${var.aws.region}.amazonaws.com/${var.name}:latest"
  }

}


resource "aws_ecr_lifecycle_policy" "this" {
  repository = aws_ecr_repository.this.name
  policy = jsonencode({
    "rules" : [
      {
        "rulePriority" : 1,
        "description" : "keep latest 3 images",
        "selection" : {
          "tagStatus" : "any",
          "countType" : "imageCountMoreThan",
          "countNumber" : 3
        },
        "action" : {
          "type" : "expire"
        }
      }
    ]
  })
}
