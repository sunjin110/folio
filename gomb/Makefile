image/lambda/build:
	@docker build -t gomb-lambda -f Dockerfile.lambda --platform=linux/arm64 .

image/lambda/build/test:
	@docker build -t gomb-lambda:test -f Dockerfile.lambda.test --platform=linux/arm64 .

image/lambda/test:
	@docker run -d -p 9000:8080 --entrypoint /usr/local/bin/aws-lambda-rie gomb-lambda:test ./main

lambda/test/req:
	@curl "http://localhost:9000/2015-03-31/functions/function/invocations" -d '{"path":"/"}'

deploy/lambda/prd:
	@make image/lambda/build
	@make login/aws/ecr/prd
	@make push/lambda/prd
	@aws --region ap-northeast-1 --profile folio-terraform lambda update-function-code --function-name production-gomb --image-uri 891377020091.dkr.ecr.ap-northeast-1.amazonaws.com/production-gomb:latest

login/aws/ecr/prd:
	@aws --region ap-northeast-1 --profile folio-terraform ecr get-login-password | docker login --username AWS --password-stdin https://891377020091.dkr.ecr.ap-northeast-1.amazonaws.com

push/lambda/prd:
	@docker push 891377020091.dkr.ecr.ap-northeast-1.amazonaws.com/production-gomb:latest
