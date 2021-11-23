KOYEB_API ?= $(shell echo $${KOYEB_API:-https://developer.koyeb.com})
TEST_OPTS=-v -test.timeout 300s


.PHONY: gen
gen: api/v1/koyeb/openapi.json
	docker run --rm -v `pwd`:/builder koyeb/swagger-build:2.0.0 generate -i /builder/api/v1/koyeb/openapi.json -g go -o /builder/api/v1/koyeb --package-name koyeb --additional-properties enumClassPrefix=true --type-mappings=object=interface{} --additional-properties isGoSubmodule=true --additional-properties generateInterfaces=true


clean:
	rm -rf api/v1/koyeb

.PHONY: api/v1/koyeb/openapi.json
api/v1/koyeb/openapi.json:
	mkdir -p api/v1/koyeb
	curl -s $(KOYEB_API)/public.swagger.json > api/v1/koyeb/openapi.json

test:
	go test $(TEST_OPTS) ./...
