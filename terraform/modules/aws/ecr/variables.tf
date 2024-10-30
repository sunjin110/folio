variable "aws" {
  type = object({
    account_id : string,
    region : string,
    profile : string,
  })
}

variable "name" {
  type        = string
  description = "ecrの名前"
}

variable "init_docker_build" {
  type = object({
    working_dir : string
    # DOCKER_BUILDKIT=1 docker build -t ${var.name}:latest -f Dockerfile.lambda --platform=linux/arm64 .
    command : string
  })
  description = "初回build設定、imageはnameと同じであること"
}
