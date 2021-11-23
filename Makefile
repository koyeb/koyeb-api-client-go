KOYEB_API ?= $(shell echo $${KOYEB_API:-https://developer.koyeb.com})
TEST_OPTS=-v -test.timeout 300s
GIT_USER_ID?=koyeb
GIT_REPO_ID?=koyeb-api-client-go


.PHONY: gen
gen: api/v1/koyeb/openapi.json
	docker run --rm -v `pwd`:/builder koyeb/swagger-build:2.0.0 generate --git-user-id ${GIT_USER_ID} --git-repo-id ${GIT_REPO_ID} -i /builder/api/v1/koyeb/openapi.json -g go -o /builder/api/v1/koyeb --package-name koyeb --additional-properties enumClassPrefix=true --type-mappings=object=interface{} --additional-properties isGoSubmodule=true --additional-properties generateInterfaces=true
	cd api/v1/koyeb && go mod tidy


clean:
	rm -rf api/v1/koyeb

.PHONY: api/v1/koyeb/openapi.json
api/v1/koyeb/openapi.json:
	mkdir -p api/v1/koyeb
	curl -s $(KOYEB_API)/public.swagger.json > api/v1/koyeb/openapi.json

test:
	cd api/v1/koyeb && go test $(TEST_OPTS) ./...
