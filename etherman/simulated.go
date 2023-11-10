package etherman

import (
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
)

func NewSimulated(auth *bind.TransactOpts) (*Etherman, error) {
	if auth == nil {
		return nil, errors.New("auth should not be nil")
	}

	// Simulated client setup
	balance, _ := new(big.Int).SetString("10000000000000000000000000", 10) //nolint:gomnd
	address := auth.From
	genesisAlloc := map[common.Address]core.GenesisAccount{
		address: {
			Balance: balance,
		},
	}
	blockGasLimit := uint64(999999999999999999) //nolint:gomnd
	client := backends.NewSimulatedBackend(genesisAlloc, blockGasLimit)
	_, _, contract, err := DeployCraterContract(client, auth)
	if err != nil {
		return nil, err
	}

	// Deploy contract
	return &Etherman{
		ethClient: client,
		auth:      auth,
		contract:  contract,
	}, nil
}
