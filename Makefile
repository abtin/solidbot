.PHONY: check-env build-solidity show-abi



check-env:
	@cd scripts; ./check-env.sh

clean:
	@rm -rf contracts/artifacts

build-solidity:
	@solc --bin --abi --overwrite -o contracts/artifacts contracts/*.sol

show-abi:
	@cat contracts/artifacts/Journal.abi | jq '.'