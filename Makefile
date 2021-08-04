.PHONY: env.check solidity.clean solidity.build solidity.show.abi



env.check:
	@cd scripts; ./check-env.sh

clean:
	@rm -rf contracts/artifacts

solidity.build:
	@solc --bin --abi --overwrite -o contracts/artifacts contracts/*.sol

solidity.show.abi:
	@cat contracts/artifacts/Journal.abi | jq '.'

build.all: solidity.build go.build


test.integration:
  go test  integration-test/ -count=1