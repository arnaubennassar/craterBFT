package etherman

import (
	"math/big"

	"github.com/arnaubennassar/craterBFT/etherman/cratercontract"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
)

type Comet struct {
	*cratercontract.CraterComet
}

func NewComet(msg []byte, nonce, blockDeadline *big.Int) *Comet {
	return &Comet{
		&cratercontract.CraterComet{
			Message:       msg,
			Nonce:         nonce,
			BlockDeadline: blockDeadline,
		},
	}
}

func (c *Comet) HashToSign() common.Hash {
	types := []string{
		"bytes", // oldAccInputHash
		"uint",  // currentTransactionsHash
		"uint",  // globalExitRoot
	}
	values := []interface{}{
		c.Message,
		c.Nonce,
		c.BlockDeadline,
	}
	return common.Hash(solsha3.SoliditySHA3(types, values))
}

func (e *Etherman) SendComet(comet Comet, signaturesAndAddrs []byte) (common.Hash, error) {
	tx, err := e.contract.SendComet(e.auth, *comet.CraterComet, signaturesAndAddrs)
	if err != nil {
		return common.Hash{}, err
	}
	return tx.Hash(), nil
}

func (e *Etherman) GetCrater(i *big.Int) ([]byte, error) {
	return e.contract.Craters(&bind.CallOpts{Pending: false}, i)
}
