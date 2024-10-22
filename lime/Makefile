image/lambda/build:
	@docker build -t lime-lambda -f Dockerfile.lambda --platform=linux/arm64 .

image/lambda/build/test:
	@docker build -t lime-lambda:test -f Dockerfile.lambda.test --platform=linux/arm64 .

image/lambda/test:
	@docker run -d -p 9000:8080 --entrypoint /usr/local/bin/aws-lambda-rie lime-lambda:test ./main

lambda/test/req:
	@curl "http://localhost:9000/2015-03-31/functions/function/invocations" -d '{"path":"/"}'
