
run:
	
	go run ./cmd/http/main.go

gen-code:
	@docker run --rm \
  		-v ${PWD}:/local openapitools/openapi-generator-cli generate \
  		-i /local/schema/http/swagger.json \
  		-g go-server \
		--additional-properties=onlyInterfaces=true,outputAsLibrary=true,sourceFolder=openapi \
		--git-repo-id kraken/ogp/generate/schema/go \
		--git-user-id asahi-digital \
  		-o /local/generate/schema/http/go

	
	