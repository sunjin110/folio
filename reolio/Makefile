# gen-code:
# 	openapi-generator-cli generate -i input.yaml -g typescript-fetch -o output


gen-code:
	@docker run --rm \
  		-v ${PWD}/../:/local openapitools/openapi-generator-cli generate \
  		-i /local/golio/schema/http/folio.openapi.json \
  		-g typescript-fetch \
  		-o /local/reolio/src/generate/schema/http
