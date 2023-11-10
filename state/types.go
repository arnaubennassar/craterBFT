package state

import (
	"github.com/ethereum/go-ethereum/common"
)

type Comet struct {
	Message            []byte
	CometNum           uint64
	BlockDeadline      uint64
	SignaturesAndAddrs []byte
}

type Crater struct {
	TxHash    common.Hash
	CraterNum uint64
}
