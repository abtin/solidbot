package integrationtest

import (
	"github.com/cellcycle/go-web3"
	"github.com/cellcycle/go-web3/providers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAccounts(t *testing.T) {
	var w3 = web3.NewWeb3(providers.NewHTTPProvider("127.0.0.1:8545", 10, false))
	accounts, err := w3.Eth.ListAccounts()
	if err != nil {
		assert.Failf(t, "Failed to connect to local node", "no local node is running")
	}
	assert.NotEmpty(t, accounts)
	for i := range accounts {
		t.Log(accounts[i])
	}
}
