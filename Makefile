.PHONY: env.check solidity.clean solidity.build solidity.show.abi

SHELL := /bin/bash # Use bash syntax

env.check:
	@cd ./scripts; ./check-env.sh

solidity.clean:
	@rm -rf ./contracts/artifacts

solidity.build:
	@solc --bin --abi --overwrite -o ./contracts/artifacts contracts/*.sol

solidity.show.abi:
	@cat ./contracts/artifacts/Journal.abi | jq '.'


backend.api.gen:
	docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli generate \
        -i ./openapi/backend-api.yaml\
        -g go \
        -o ./internal/backend/


test.integration:
	@docker-compose -f ./docker-compose-test.yaml up -d
	@$(shell timeout 7 sh -c 'until nc -z localhost 8545; do sleep 1; done')
	@go test -v ./integrationtest/integration_test.go -count=1
	@docker-compose -f ./docker-compose-test.yaml down