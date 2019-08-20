package main

import (
	"github.com/threefoldtech/rivine/extensions/minting"
  mintingcli "github.com/threefoldtech/rivine/extensions/minting/client"
  

	"github.com/threefoldtech/rivine/pkg/client"
	"github.com/threefoldtech/rivine/types"

  tfbchaintypes "github.com/threefoldtech/tfbchain/pkg/types"
)

// RegisterStandardTransactions registers the tfbchain-specific transactions as required for the standard network.
func RegisterStandardTransactions(cli *client.CommandLineClient) {
	registerTransactions(cli)
}

// RegisterTestnetTransactions registers the tfbchain-specific transactions as required for the test network.
func RegisterTestnetTransactions(cli *client.CommandLineClient) {
	registerTransactions(cli)
}

// RegisterDevnetTransactions registers the tfbchain-specific transactions as required for the dev network.
func RegisterDevnetTransactions(cli *client.CommandLineClient) {
	registerTransactions(cli)
}

func registerTransactions(cli *client.CommandLineClient) {
  // create minting plugin client...
  mintingCLI := mintingcli.NewPluginConsensusClient(cli)
  // ...and register minting types
  types.RegisterTransactionVersion(tfbchaintypes.MinterDefinitionTxVersion, minting.MinterDefinitionTransactionController{
    MintConditionGetter: mintingCLI,
    TransactionVersion:  tfbchaintypes.MinterDefinitionTxVersion,
  })
  types.RegisterTransactionVersion(tfbchaintypes.CoinCreationTxVersion, minting.CoinCreationTransactionController{
    MintConditionGetter: mintingCLI,
    TransactionVersion:  tfbchaintypes.CoinCreationTxVersion,
  })
	types.RegisterTransactionVersion(tfbchaintypes.CoinDestructionTxVersion, minting.CoinDestructionTransactionController{
		TransactionVersion: tfbchaintypes.CoinDestructionTxVersion,
	})

  
}