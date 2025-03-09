run:
	go run ./cmd/http/main.go

up:
	docker compose build
	docker compose up --watch

down:
	docker compose down

lint:
	# brew install golangci-lint
	@golangci-lint run --path-prefix=.

build:
	go build ./cmd/http/main.go

image/build:
	@docker build -t golio --platform=linux/arm64 .

image/lambda/build:
	@docker build -t golio-lambda -f Dockerfile.lambda --platform=linux/arm64 .

image/lambda/test/build:
	@docker build -t golio-lambda:test -f Dockerfile.lambda.test --platform=linux/arm64 .

image/lambda/test:
	@docker run -d -p 9000:8080 --entrypoint /usr/local/bin/aws-lambda-rie golio-lambda:test ./main

gen-code:
	@docker run --rm \
  		-v ${PWD}:/local openapitools/openapi-generator-cli:v7.7.0 generate \
  		-i /local/schema/http/folio.openapi.json \
  		-g go-server \
		--additional-properties=onlyInterfaces=true,outputAsLibrary=true,sourceFolder=openapi \
		--git-repo-id kraken/ogp/generate/schema/go \
		--git-user-id asahi-digital \
  		-o /local/generate/schema/http/go

	# import文が不十分なので追加する
	@./scripts/openapi-gen-code-helper.sh
	