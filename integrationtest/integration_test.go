package integrationtest

import (
	"fmt"
	"github.com/cellcycle/go-web3"
	"github.com/cellcycle/go-web3/dto"
	"github.com/cellcycle/go-web3/providers"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"math/big"
	"os"
	"testing"
)

const (
	debug    = false
	testNode = "127.0.0.1:8545"
)

var (
	contractDeployer string
	user             string
)

// TestAccounts verifies if we can get all 10 test accounts from the test node
func TestAccounts(t *testing.T) {
	var w3 = web3.NewWeb3(providers.NewHTTPProvider(testNode, 10, false))
	accounts, err := w3.Eth.ListAccounts()
	if err != nil {
		assert.Failf(t, "Failed to connect to local node", "no local node is running")
	}
	assert.NotEmpty(t, accounts)
	assert.Equal(t, len(accounts), 10)
	if debug {
		for i := range accounts {
			t.Log(accounts[i])
		}
	}
}

func TestDeployContract(t *testing.T) {
	var w3 = web3.NewWeb3(providers.NewHTTPProvider(testNode, 10, false))
	accounts, err := w3.Eth.ListAccounts()
	if err != nil {
		assert.Failf(t, "Failed to connect to local node", "no local node is running")
	}
	assert.NotEmpty(t, accounts)
	contractDeployer = accounts[0]
	bf, err := os.Open("../contracts/artifacts/Bounties.bin")
	if err != nil {
		assert.Failf(t, "Failed to open contract bytecode", err.Error())
	}
	defer bf.Close()
	bytecode, err := ioutil.ReadAll(bf)
	if err != nil {
		assert.Failf(t, "Failed to read contract bytecode", err.Error())
	}
	af, err := os.Open("../contracts/artifacts/Bounties.abi")
	if err != nil {
		assert.Failf(t, "Failed to open contract abi", err.Error())
	}
	defer af.Close()
	abi, err := ioutil.ReadAll(af)
	if err != nil {
		assert.Failf(t, "Failed to read contract abi", err.Error())
	}

	contract, err := w3.Eth.NewContract(string(abi))
	if err != nil {
		assert.Failf(t, "Failed to create the contract", err.Error())
	}

	// coinbase, err := w3.Eth.GetCoinbase()
	if err != nil {
		assert.Failf(t, "Failed to get coinbase", err.Error())
	}

	var transaction dto.TransactionParameters
	transaction.From = contractDeployer
	transaction.Gas = big.NewInt(90820)

	hash, err := contract.Deploy(&transaction, string(bytecode), nil)
	if err != nil {
		assert.Failf(t, "Failed to deploy the contract", err.Error())
	}

	fmt.Println(hash)
}
