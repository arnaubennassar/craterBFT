package etherman

import (
	"github.com/arnaubennassar/craterBFT/etherman/cratercontract"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func DeployCraterContract(client bind.ContractBackend, auth *bind.TransactOpts) (common.Address, *types.Transaction, *cratercontract.Cratercontract, error) {
	return cratercontract.DeployCratercontract(auth, client)
}
