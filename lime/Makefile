image/lambda/build:
	@docker build -t lime-lambda -f Dockerfile.lambda --platform=linux/arm64 .

image/lambda/build/test:
	@docker build -t lime-lambda:test -f Dockerfile.lambda.test --platform=linux/arm64 .

image/lambda/test:
	@docker run -d -p 9000:8080 --entrypoint /usr/local/bin/aws-lambda-rie lime-lambda:test ./main

lambda/test/req:
	@curl "http://localhost:9000/2015-03-31/functions/function/invocations" -d '{"path":"/"}'

gen/golio/openapi:
	@docker run --rm \
  		-v ${PWD}/../:/local openapitools/openapi-generator-cli:v7.9.0 generate \
  		-i /local/golio/schema/http/folio.openapi.json \
  		-g go \
		--git-repo-id github.com/sunjin110/folio/lime \
		--git-user-id sunjin110 \
  		-o /local/lime/generate/schema/golio/http/openapi

	@rm -rf generate/schema/golio/http/openapi/test