package bft

import (
	"github.com/arnaubennassar/craterBFT/state"
	"github.com/cometbft/cometbft/libs/log"
	"github.com/ethereum/go-ethereum/common"
)

type CraterNode struct {
	ID     string
	logger log.Logger
	addr   common.Address
	state  *state.State
	eth    *Ethermaner
}

func New(
	eth *Ethermaner, logger log.Logger, identity string, addr common.Address, state *state.State,
) *CraterNode {
	return &CraterNode{
		ID:     identity,
		logger: logger,
		addr:   addr,
		state:  state,
		eth:    eth,
	}
}
