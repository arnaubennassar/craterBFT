package bft

import (
	"context"
	"math/big"

	"github.com/arnaubennassar/craterBFT/etherman"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type Ethermaner interface {
	GetCrater(i *big.Int) ([]byte, error)
	CheckTxWasMined(ctx context.Context, txHash common.Hash) (bool, *types.Receipt, error)
	SendComet(comet etherman.Comet, signatiresAndAddrs []byte) (common.Hash, error)
}
