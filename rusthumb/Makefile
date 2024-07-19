run:
	cargo run --bin lambda

fmt:
	cargo fmt

test:
	docker build -f Dockerfile.test -t rusthumb:test .
	docker run --rm rusthumb:test

build:
	cargo build --release

build/lambda:
	cargo lambda build --locked --release --arm64

image/lambda/build:
	@docker build --progress=plain -t rusthumb-lambda -f Dockerfile.lambda --platform=linux/arm64 .

image/lambda/test/build:
	@docker build -t rusthumb-lambda:test -f Dockerfile.lambda.test --platform=linux/arm64 .

image/lambda/test:
	# @docker run -p 9000:8080 --entrypoint /usr/local/bin/aws-lambda-rie rusthumb-lambda:test ./lambda
	docker run -p 9000:8080 rusthumb-lambda:test

image/lambda/deploy/production:
	aws --profile folio-terraform ecr get-login-password --region ap-northeast-1 | docker login --username AWS --password-stdin 891377020091.dkr.ecr.ap-northeast-1.amazonaws.com
	@make image/lambda/build
	docker tag rusthumb-lambda:latest 891377020091.dkr.ecr.ap-northeast-1.amazonaws.com/production-rusthumb:latest
	docker push 891377020091.dkr.ecr.ap-northeast-1.amazonaws.com/production-rusthumb:latest

	# lambda
	aws --profile folio-terraform --region ap-northeast-1 lambda update-function-code --function-name production-rusthumb-lambda --image-uri 891377020091.dkr.ecr.ap-northeast-1.amazonaws.com/production-rusthumb:latest