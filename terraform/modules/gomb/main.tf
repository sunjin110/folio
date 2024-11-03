
module "ecr" {
  source = "../aws/ecr"
  aws    = var.aws
  name   = "${var.prefix}-gomb"
  init_docker_build = {
    working_dir = "${path.module}/../../../gomb"
    command     = "DOCKER_BUILDKIT=1 docker build -t ${var.prefix}-gomb:latest -f Dockerfile.lambda --platform=linux/arm64 ."
  }
}

module "iam" {
  source = "./iam"
  aws    = var.aws
  prefix = var.prefix
}

module "lambda" {
  source       = "../aws/lambda"
  name         = "${var.prefix}-gomb"
  prefix       = var.prefix
  ecr          = module.ecr.repository
  iam          = module.iam
  architecture = "arm64"
  memory_size  = 512 # MB
  environment = {
    "MEDIA_S3_BUCKET_NAME" : var.media_s3.name
  }
}

module "s3_notification" {
  source = "./s3_notification"
  s3     = var.media_s3
  lambda = module.lambda
}
